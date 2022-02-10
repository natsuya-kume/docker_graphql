package utils

import (
	"encoding/base64"
	"log"
)

func Decode(inputID string) []byte {
	dec, err := base64.StdEncoding.DecodeString(inputID)
	if err != nil {
		log.Fatal(err)
	}
	return dec
}
