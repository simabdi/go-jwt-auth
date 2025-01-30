package helper

import (
	"encoding/base64"
	"encoding/json"
	"github.com/simabdi/go-jwt-auth/model"
	"log"
	"os"
	"time"
)

func JsonResponse(code int, message string, success bool, error string, data interface{}) model.Response {
	meta := model.Meta{
		Code:    code,
		Status:  success,
		Message: message,
		Error:   error,
	}

	response := model.Response{
		Meta: meta,
		Data: data,
	}

	return response
}

func Logger(requestType string, request interface{}, bodyBytes []byte) {
	fo, err := os.OpenFile("storage/logs/log-"+time.Now().Format("2006-01-02")+".txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	req, err := json.Marshal(&request)
	if err != nil {
		log.Println("[Error Log Marshal] : ", err.Error())
	}

	var resJson map[string]interface{}
	err = json.Unmarshal(bodyBytes, &resJson)
	if err != nil {
		log.Println("[Error Log Unmarshal] : ", err.Error())
	}

	resDecode, err := json.Marshal(resJson)
	if err != nil {
		log.Println("[Error Log res Marshal] : ", err.Error())
	}

	text := []byte(
		"===========================================================================================\n" +
			requestType + " " + time.Now().Format("2006-01-02 15:04:05") +
			"\n===========================================================================================\n" +
			"=================================REQUEST===================================\n" +
			string(req) + "\n" +
			"=================================RESPONSE==================================\n" +
			string(resDecode) + "\n\n\n")

	_, err = fo.WriteString(string(text))
	if err != nil {
		log.Println("[Error Log WriteString] : ", err.Error())
	}

	defer fo.Close()
}

func Std64Encode(plainText string) string {
	return base64.StdEncoding.EncodeToString([]byte(plainText))
}

func Std64Decode(encoded string) string {
	decodedByte, _ := base64.StdEncoding.DecodeString(encoded)
	return string(decodedByte)
}
