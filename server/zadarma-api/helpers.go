package zadarma_api

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"net/url"
	"sort"
	"strings"
	"unicode/utf8"
)

func BuildAPIUrl(methodName string, params url.Values) string {
	var resultURL string

	if strings.HasPrefix(methodName, "/") && strings.HasSuffix(API_URL, "/") {
		resultURL = API_URL + strings.TrimLeft(methodName, "/") + "?"
	} else {
		resultURL = API_URL + methodName + "?"
	}

	for name, value := range params {
		resultURL += name + "=" + value[0] + "&"
	}

	resultURL = strings.TrimRight(resultURL, "&")
	return resultURL
}

func Sign(api apiClient, methodName string, params url.Values) string {
	var paramParts []string
	for name, value := range params {
		paramParts = append(paramParts, name+"="+value[0])
	}

	sort.Slice(paramParts, func(i, j int) bool {
		firstRune, _ := utf8.DecodeRuneInString(paramParts[i])
		secondRune, _ := utf8.DecodeRuneInString(paramParts[j])

		return firstRune < secondRune
	})

	paramsUrlStr := strings.Join(paramParts, "&")

	md5Params := fmt.Sprintf("%x", md5.Sum([]byte(paramsUrlStr)))

	sign := methodName + paramsUrlStr + md5Params
	hmacer := hmac.New(sha1.New, []byte(api.Secret))
	hmacer.Write([]byte(sign))
	sha1Result := hmacer.Sum(nil)
	sha1Hash := fmt.Sprintf("%x", sha1Result)

	base64Encoded := base64.StdEncoding.EncodeToString([]byte(sha1Hash))

	return base64Encoded
}
