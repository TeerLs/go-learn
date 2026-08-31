package verify

type SendResponse struct {
	Email string `json:"email"`
}

type VerificationResponse struct {
	Message string `json:"message"`
}

type VerifyResponse struct {
	Code string `json:"code"`
}

type EmailData struct {
	Code string
}

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type LinkHashJSON struct {
	Links map[string]string `json:"links"`
}