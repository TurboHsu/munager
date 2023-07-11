package netease

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"math/rand"
	"net/url"
	"strings"
	"time"

	"github.com/TurboHsu/munager/cryptography"
)

// Crypto part of Netease API is from Muget, see README.md for more details

type APIType int

const (
	Default APIType = iota
	WEAPI
	LinuxAPI
)

const (
	Base62                      = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	PresetKey                   = "0CoJUm6Qyw8W8jud"
	IV                          = "0102030405060708"
	LinuxAPIKey                 = "rFgB&h#%2?^eDg:Q"
	EAPIKey                     = "e82ckenh8dichen8"
	DefaultRSAPublicKeyModulus  = "e0b509f6259df8642dbc35662901477df22677ec152b5ff68ace615bb7b725152b3ab17a876aea8a5aa76d2e417629ec4ee341f56135fccf695280104e0312ecbda92557c93870114af6c9d05c4f7f0c3685b7a46bee255932575cce10b424d813cfe4875d3e82047b97ddef52741d546b8e289dc6935b3ece0462db0a22b8e7"
	DefaultRSAPublicKeyExponent = 0x10001
)

func convertParams(origData interface{}, apiType APIType) url.Values {
	switch apiType {
	case WEAPI:
		plainText, _ := json.Marshal(origData)
		params := base64.StdEncoding.EncodeToString(cryptography.AESCBCEncrypt(plainText, []byte(PresetKey), []byte(IV)))
		secKey := createSecretKey(16, Base62)
		params = base64.StdEncoding.EncodeToString(cryptography.AESCBCEncrypt([]byte(params), secKey, []byte(IV)))
		return url.Values{
			"params":    {params},
			"encSecKey": {cryptography.RSAEncrypt(bytesReverse(secKey), DefaultRSAPublicKeyModulus, DefaultRSAPublicKeyExponent)},
		}
	case LinuxAPI:
		plainText, _ := json.Marshal(origData)
		return url.Values{
			"eparams": {strings.ToUpper(hex.EncodeToString(cryptography.AESECBEncrypt(plainText, []byte(LinuxAPIKey))))},
		}
	case Default:
		fallthrough
	default:
		// Append everything from origData
		params := url.Values{}
		for k, v := range origData.(map[string]interface{}) {
			params.Add(k, v.(string))
		}
		return params
	}
}

func createSecretKey(size int, charset string) []byte {
	secKey, n := make([]byte, size), len(charset)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range secKey {
		secKey[i] = charset[r.Intn(n)]
	}
	return secKey
}

func bytesReverse(b []byte) []byte {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return b
}
