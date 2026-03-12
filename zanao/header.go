package zanao

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var _m = ""

func init() {
	getM(20)
}

func getM(length int) string {
	if _m == "" {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		result := make([]byte, length)
		for i := range result {
			result[i] = byte(r.Intn(10)) + '0'
		}
		_m = string(result)
	}
	return _m
}

func md5Hash(input string) string {
	hasher := md5.New()
	hasher.Write([]byte(input))
	return hex.EncodeToString(hasher.Sum(nil))
}
func getHeaders(userToken, schoolalias string) map[string]string {
	m := getM(20)
	td := time.Now().Unix()
	signString := fmt.Sprintf("%s_%s_%d_1b6d2514354bc407afdd935f45521a8c", schoolalias, m, td)
	return map[string]string{
		"X-Sc-Version":  "3.4.4",
		"X-Sc-Nwt":      "wifi",
		"X-Sc-Wf":       "",
		"X-Sc-Nd":       m,
		"X-Sc-Cloud":    "0",
		"X-Sc-Platform": "windows",
		"X-Sc-Appid":    "wx3921ddb0258ff14f",
		"X-Sc-Alias":    schoolalias,
		"X-Sc-Od":       userToken,
		"Content-Type":  "application/x-www-form-urlencoded",
		"X-Sc-Ah":       md5Hash(signString),
		"xweb_xhr":      "1",
		"User-Agent":    "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36 MicroMessenger/7.0.20.1781(0x6700143B) NetType/WIFI MiniProgramEnv/Windows WindowsWechat/WMPF WindowsWechat(0x63090c33)XWEB/14185",
		"X-Sc-Td":       strconv.FormatInt(td, 10),
		"Accept":        "*/*",
	}
}
