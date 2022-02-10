package utils

import (
	"encoding/base64"
)

func Encode(modelName string, modelID string) string {
	src := []byte(modelName + modelID)
	encID := base64.StdEncoding.EncodeToString(src)
	return encID
}
