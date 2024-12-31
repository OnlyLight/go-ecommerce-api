package sendto

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/onlylight29/go-ecommerce-backend-api/global"
	"go.uber.org/zap"
)

type MailRequest struct {
	ToEmail     string `json:"toEmail"`
	MessageBody string `json:"messageBody"`
	Subject     string `json:"subject"`
	Attachment  string `json:"attachment"`
}

func sendEmailToJavaByApi(otp string, email string, purpose string) error {
	postUrl := ""

	// Data json
	mailRequest := MailRequest{
		ToEmail:     email,
		MessageBody: "OTP is::" + otp,
		Subject:     "Verify OTP::" + purpose,
		Attachment:  "",
	}

	// convert struct to java
	requestBody, err := json.Marshal(mailRequest)
	if err != nil {
		global.Logger.Error("Error while converting struct to java", zap.Error(err))
		return err
	}

	// Create request
	req, err := http.NewRequest("POST", postUrl, bytes.NewBuffer(requestBody))
	if err != nil {
		global.Logger.Error("Error while creating request", zap.Error(err))
		return err
	}

	// Set Header
	req.Header.Set("Content-Type", "application/json")

	// execute request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		global.Logger.Error("Error while executing request", zap.Error(err))
		return err
	}
	defer res.Body.Close()

	return nil
}
