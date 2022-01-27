package main

import (
	b64 "encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"
)

func main() {
	testToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyLCJuZXN0ZWQiOnsidmFsdWUiOnRydWUsImFyciI6WzEsMiwzXX19.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
	err, _, payload := decodeJwt(testToken)

	if err != nil {
		fmt.Println("Found error")
		fmt.Printf("%s", err)
	}

	printJwt(payload)
}

func printJwt(token map[string]interface{}) {
	printObject(token, 0, false)
}

func printObject(obj map[string]interface{}, indent int, skipOpeningBrace bool) {
	prevIndent := strings.Repeat("\t", indent)
	if !skipOpeningBrace {
		fmt.Printf("%s{\n", prevIndent)
	}
	newIndent := indent + 1
	curIndent := strings.Repeat("\t", newIndent)

	for key, val := range obj {
		valType := fmt.Sprintf("%s", reflect.TypeOf(val).Kind())
		if valType == "map" {
			fmt.Printf("%s%s: {\n", curIndent, key)
			nestedMap := val.(map[string]interface{})
			printObject(nestedMap, newIndent, true)
		} else if valType == "string" {
			fmt.Printf("%s%s: \"%s\",\n", curIndent, key, val)
		} else if strings.HasPrefix(valType, "int") {
			fmt.Println(valType)
			fmt.Printf("%s%s: %d,\n", curIndent, key, val)
		} else if strings.HasPrefix(valType, "float") {
			fmt.Printf("%s%s: %.f,\n", curIndent, key, val)
		} else if valType == "bool" {
			fmt.Printf("%s%s: %t,\n", curIndent, key, val)
		} else if valType == "slice" {
			semiformat := fmt.Sprint(val)
			fmt.Printf("%s%s: %s,\n", curIndent, key, semiformat)
		} else {
			fmt.Printf("unsupported type, %s\n", valType)
		}
	}
	fmt.Printf("%s}\n", prevIndent)
}

// given some jwt of the form header.payload.body,
// return an interface containint the header and payload
func decodeJwt(rawToken string) (error, map[string]interface{}, map[string]interface{}) {
	tokenSplit := strings.Split(rawToken, ".")
	if len(tokenSplit) != 3 {
		return errors.New("Invalid JWT passed"), nil, nil
	}

	// decode header and payload
	_, header := decodeAndUnmarshall(tokenSplit[0])
	_, payload := decodeAndUnmarshall(tokenSplit[1])
	return nil, header, payload
}

func decodeAndUnmarshall(tokenChunk string) (error, map[string]interface{}) {
	decodedRaw, err := b64.StdEncoding.DecodeString(tokenChunk)
	if err != nil {
		decodedRaw, err = b64.RawStdEncoding.DecodeString(tokenChunk)
		if err != nil {
			return err, nil
		}
	}

	decoded := string(decodedRaw)

	var token map[string]interface{}

	err = json.Unmarshal([]byte(decoded), &token)

	if err != nil {
		return err, nil
	}

	return nil, token

}
