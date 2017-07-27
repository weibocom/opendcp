package ssh

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"runtime"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	gossh "golang.org/x/crypto/ssh"
	"weibo.com/opendcp/jupiter/models"
	"weibo.com/opendcp/jupiter/dao"
	"strings"
)

var (
	ErrKeyGeneration     = errors.New("Unable to generate key")
	ErrValidation        = errors.New("Unable to validate key")
	ErrPublicKey         = errors.New("Unable to convert public key")
	ErrUnableToWriteFile = errors.New("Unable to write file")
)

const SSH_AWS = "/go/src/weibo.com/opendcp/jupiter/conf/zhaowei9.pem"

type KeyPair struct {
	PrivateKey []byte
	PublicKey  []byte
}


// Generate a new SSH keypair
// This will return a private & public key encoded as DER.
func NewKeyPair() (keyPair *KeyPair, err error) {
	priv, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, ErrKeyGeneration
	}

	if err := priv.Validate(); err != nil {
		return nil, ErrValidation
	}

	privDer := x509.MarshalPKCS1PrivateKey(priv)

	pubSSH, err := gossh.NewPublicKey(&priv.PublicKey)
	if err != nil {
		return nil, ErrPublicKey
	}

	return &KeyPair{
		PrivateKey: privDer,
		PublicKey:  gossh.MarshalAuthorizedKey(pubSSH),
	}, nil
}

// Write keypair to files
func (kp *KeyPair) WriteToFile(privateKeyPath string, publicKeyPath string) error {
	files := []struct {
		File  string
		Type  string
		Value []byte
	}{
		{
			File: privateKeyPath,
			Value: pem.EncodeToMemory(&pem.Block{
				Type:    "RSA PRIVATE KEY",
				Headers: nil,
				Bytes:   kp.PrivateKey,
			}),
		},
		{
			File:  publicKeyPath,
			Value: kp.PublicKey,
		},
	}

	for _, v := range files {
		f, err := os.Create(v.File)
		if err != nil {
			return ErrUnableToWriteFile
		}

		if _, err := f.Write(v.Value); err != nil {
			return ErrUnableToWriteFile
		}

		// windows does not support chmod
		switch runtime.GOOS {
		case "darwin", "linux":
			if err := f.Chmod(0600); err != nil {
				return err
			}
		}
	}

	return nil
}

// Calculate the fingerprint of the public key
func (kp *KeyPair) Fingerprint() string {
	b, _ := base64.StdEncoding.DecodeString(string(kp.PublicKey))
	h := md5.New()

	io.WriteString(h, string(b))

	return fmt.Sprintf("%x", h.Sum(nil))
}

func GenSSHKey(path string) error {
	if _, err := os.Stat(path); err != nil {
		if !os.IsNotExist(err) {
			return err
		}
		kp, err := NewKeyPair()
		if err != nil {
			return err
		}
		if err := kp.WriteToFile(path, fmt.Sprintf("%s.pub", path)); err != nil {
			return err
		}
	}
	return nil
}

// Disable ssh passowrd login
func (sshCli *Client) DisPassLogin() error {
	cmd := "sed -r -i 's/^(PasswordAuthentication )yes$/\\1no/' /etc/ssh/sshd_config | kill -HUP `cat /var/run/sshd.pid`"
	_, err := sshCli.Run(cmd)
	if err != nil {
		return err
	}
	return nil
}

func (kp *KeyPair) UploadKey(sshCli *Client) error {
	cmd := fmt.Sprintf("mkdir -p ~/.ssh; echo '%s' > ~/.ssh/authorized_keys", string(kp.PublicKey))
	_, err := sshCli.Run(cmd)
	if err != nil {
		return err
	}
	return nil
}

// Generate SSH keypair based on path of the private key
// The public key would be generated to the same path with ".pub" added
func (sshCli *Client) GenerateSSHKey(instanceId string, path string) error {
	if _, err := os.Stat(path); err != nil {
		if !os.IsNotExist(err) {
			return err
		}
		kp, err := NewKeyPair()
		if err != nil {
			return err
		}
		if err := kp.StoreDb(instanceId); err != nil {
			return err
		}
		if err := kp.WriteToFile(path, fmt.Sprintf("%s.pub", path)); err != nil {
			return err
		}
		if err := kp.UploadKey(sshCli); err != nil {
			return err
		}

	}
	return nil
}

func (sshCli *Client) StoreSSHKey(instanceId string) error {
	kp, err := NewKeyPair()

	ins ,err := dao.GetInstance(instanceId)
	if err != nil {
		return err
	}

	if strings.EqualFold(ins.Provider, "aws") {
		outputBytes, _ := ioutil.ReadFile(SSH_AWS)
		privateBlock, _ := pem.Decode(outputBytes)
		kp.PrivateKey = privateBlock.Bytes
	}

	if err != nil {
		return err
	}
	if err := kp.StoreDb(instanceId); err != nil {
		return err
	}
	if err := kp.UploadKey(sshCli); err != nil {
		return err
	}
	return nil
}

func (kp *KeyPair) StoreDb(instanceId string) error {
	o := orm.NewOrm()
	ma := models.Instance{InstanceId: instanceId}
	if o.Read(&ma, "InstanceId") == nil {
		ma.PrivateKey = string(pem.EncodeToMemory(&pem.Block{
			Type:    "RSA PRIVATE KEY",
			Headers: nil,
			Bytes:   kp.PrivateKey,
		}),
		)
		ma.PublicKey = string(kp.PublicKey)
		_, err := o.Update(&ma)
		if err != nil {
			return err
		}
	} else {
		return o.Read(&ma)
	}
	return nil
}

func GetSSHKeyFromDb(mac *models.Instance, path string, isPriv bool) error {
	if _, err := os.Stat(path); err != nil {
		if !os.IsNotExist(err) {
			return err
		}
	}
	if isPriv {
		err := ioutil.WriteFile(path, []byte(mac.PrivateKey), 0400)
		if err != nil {
			return err
		}
	} else {
		err := ioutil.WriteFile(path, []byte(mac.PublicKey), 0400)
		if err != nil {
			return err
		}
	}
	return nil
}
