package verify

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/textproto"
	"os"

	"github.com/jordan-wright/email"
)

type VerifyService struct{}

func NewVerifyService() (*VerifyService) {
	verifyService := VerifyService{}
	return &verifyService
}

func (verifyService *VerifyService) CreateEmailBody(resEmail string, hash string) (*email.Email) {
	hashLink := "http://localhost:8081/verify/" + hash

	htmlTemplate := "<h1>Your link for verifying is" + hashLink + "!</h1>"

	e := &email.Email{
		To:      []string{resEmail},
		From:    "Jordan Wright <test@gmail.com>",
		Subject: "Awesome Subject",
		Text:    []byte("Text Body is, of course, supported!"),
		HTML:    []byte(htmlTemplate),
		Headers: textproto.MIMEHeader{},
	}

	return e
}

func (verifyService *VerifyService) AddHashToJSONFile(resEmail string, hash string) (error) {
	filePath := "link-hash.json"

	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	var linksHash LinkHashJSON
	if err := json.Unmarshal(data, &linksHash); err != nil {
		return err
	}

	if _, has := linksHash.Links[resEmail]; has {
		return errors.New("Verification link already exists for this email") 
	} else {
		linksHash.Links[resEmail] = hash
	}

	updatedData, err := json.MarshalIndent(linksHash, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(filePath, updatedData, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return err
	}

	return nil
}

func (verifyService *VerifyService) VerifyEmail(hash string) (error) {
	filePath := "link-hash.json"

	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	var linksHash LinkHashJSON
	if err := json.Unmarshal(data, &linksHash); err != nil {
		return err
	}

	found := false

	for key, value := range linksHash.Links {
		if value == hash {
			found = true
			delete(linksHash.Links, key)
		}
	}

	if !found {
		return errors.New("Couldn't find your email on verification list")
	}

	updatedData, err := json.MarshalIndent(linksHash, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(filePath, updatedData, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return err
	}

	return nil
}