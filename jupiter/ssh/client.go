package ssh

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"time"

	"github.com/astaxie/beego"
	"golang.org/x/crypto/ssh"
)

type Client struct {
	Config   *ssh.ClientConfig
	Hostname string
	Port     int
}

const (
	maxDialAttempts = 10
)

func NewClient(user string, host string, port int, auth *Auth) (*Client, error) {
	config, err := NewConfig(user, auth)
	if err != nil {
		return nil, err
	}

	return &Client{
		Config:   config,
		Hostname: host,
		Port:     port,
	}, nil
}

func NewConfig(user string, auth *Auth) (*ssh.ClientConfig, error) {
	var authMethods []ssh.AuthMethod

	for _, k := range auth.Keys {
		key, err := ioutil.ReadFile(k)
		if err != nil {
			return nil, err
		}

		privateKey, err := ssh.ParsePrivateKey(key)
		if err != nil {
			return nil, err
		}

		authMethods = append(authMethods, ssh.PublicKeys(privateKey))
	}

	for _, p := range auth.Passwords {
		authMethods = append(authMethods, ssh.Password(p))
	}

	return &ssh.ClientConfig{
		User:   user,
		Auth:   authMethods,
		Timeout:10 * time.Second,
	}, nil
}

func waitFor(f func() bool) error {
	return waitForSpecific(f, 60, 3*time.Second)
}

func waitForSpecific(f func() bool, maxAttempts int, waitInterval time.Duration) error {
	for i := 0; i < maxAttempts; i++ {
		if f() {
			return nil
		}
		time.Sleep(waitInterval)
	}
	return fmt.Errorf("Maximum number of retries (%d) exceeded", maxAttempts)
}

func dialSuccess(client *Client) func() bool {
	return func() bool {
		if _, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", client.Hostname, client.Port), client.Config); err != nil {
			beego.Debug("Error dialing TCP: %s", err)
			return false
		}
		return true
	}
}

func (client *Client) RunWithAttempt(command string, maxAttempt int) (Output, error) {
	var output Output
	if err := waitForSpecific(dialSuccess(client), maxAttempt, 3*time.Second); err != nil {
		beego.Warn("error to connect: ", client.Hostname)
		return output, fmt.Errorf("Error attempting SSH client dial: %s", err)
	}
	return client.Run(command)
}

func (client *Client) Run(command string) (Output, error) {
	var (
		output         Output
		stdout, stderr bytes.Buffer
	)

	if err := waitFor(dialSuccess(client)); err != nil {
		return output, fmt.Errorf("Error attempting SSH client dial: %s", err)
	}

	conn, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", client.Hostname, client.Port), client.Config)
	if err != nil {
		return output, fmt.Errorf("Mysterious error dialing TCP for SSH (we already succeeded at least once) : %s", err)
	}

	session, err := conn.NewSession()
	if err != nil {
		return output, fmt.Errorf("Error getting new session: %s", err)
	}

	defer session.Close()

	session.Stdout = &stdout
	session.Stderr = &stderr

	output = Output{
		Stdout: &stdout,
		Stderr: &stderr,
	}

	return output, session.Run(command)
}

type Auth struct {
	Passwords []string
	Keys      []string
}

type Output struct {
	Stdout io.Reader
	Stderr io.Reader
}
