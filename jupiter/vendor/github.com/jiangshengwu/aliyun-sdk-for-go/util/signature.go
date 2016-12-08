package util

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"io"
	"net/url"
	"sort"
	"strconv"
	"strings"
)

const SEPARATOR = "&"

// Get signature from params in map
func MapToSign(params map[string]interface{}, keySecret string, httpMethod string) string {
	key := []byte(keySecret + SEPARATOR)
	h := hmac.New(sha1.New, key)
	query := canonicalizedFromMap(params)
	toSign := httpMethod + SEPARATOR + PercentEncode("/") + SEPARATOR + query
	h.Write([]byte(toSign))
	sign := base64.StdEncoding.EncodeToString(h.Sum(nil))
	sign = PercentEncode(sign)
	return sign
}

// Get canonicalized query string from params in map
func canonicalizedFromMap(params map[string]interface{}) string {
	keys := make([]string, len(params))
	i := 0
	for k := range params {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	sign := ""
	for _, v := range keys {
		sign += SEPARATOR + PercentEncode(v) + "="
		switch params[v].(type) {
		case string:
			sign += PercentEncode(params[v].(string))
		case int:
			sign += PercentEncode(strconv.Itoa(params[v].(int)))
		default:

		}

	}
	if len(sign) > 0 {
		sign = PercentEncode(sign[1:])
	}
	return sign
}

// URL encode
func PercentEncode(str string) string {
	s := url.QueryEscape(str)
	s = strings.Replace(s, "+", "%20", -1)
	s = strings.Replace(s, "*", "%2A", -1)
	s = strings.Replace(s, "%7E", "~", -1)
	return s
}

// MD5 hash
func getMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// Generate random guid
func GetGuid() string {
	b := make([]byte, 48)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return getMd5String(base64.URLEncoding.EncodeToString(b))
}
