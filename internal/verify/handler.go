package verify

import (
	"encoding/json"
	"go/advanced-demo/configs"
	"go/advanced-demo/pkg/math"
	"go/advanced-demo/pkg/res"
	"net/http"
	"net/smtp"
)

type VerifyHandler struct {
	Config *configs.SMTPConfig
	Service *VerifyService
}

type VerifyHandlerDeps struct {
	Config *configs.SMTPConfig
}

func NewVerifyHandler(router *http.ServeMux, deps VerifyHandlerDeps) {
	service := NewVerifyService()
	handler := VerifyHandler{
		Config: deps.Config,
		Service: service,
	}
	router.HandleFunc("POST /verify/send", handler.Send())
	router.HandleFunc("GET /verify/{hash}", handler.Verify())
}

func (handler *VerifyHandler) Send() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		resBody := SendResponse{}
		err := json.NewDecoder(r.Body).Decode(&resBody)
		if err != nil {
			res.Json(w, ErrorResponse{
            Message: "invalid JSON",
        }, http.StatusBadRequest)
			return;
		}

		hash, err := math.GenerateHashLink(resBody.Email)
		if err != nil {
			res.Json(w, ErrorResponse{
            Message: "Error Creating link",
        }, http.StatusInternalServerError)
			return;
		}

		e := handler.Service.CreateEmailBody(resBody.Email, hash)

		jsonSaveError := handler.Service.AddHashToJSONFile(resBody.Email, hash) 
		if jsonSaveError != nil {
			res.Json(w, ErrorResponse{
				Message: "Error during saving link into JSON: " + jsonSaveError.Error(),
			}, http.StatusInternalServerError)
			return;
		}

		smtpErr := e.Send(handler.Config.Host + ":" + handler.Config.Port, smtp.PlainAuth("", handler.Config.Email, handler.Config.Password, handler.Config.Host))
		if smtpErr != nil {
			res.Json(w, ErrorResponse{
				Message: "Error sending email: " + smtpErr.Error(),
			}, http.StatusInternalServerError)
			return
		}
		w.Write([]byte("Verification link sent on your email"))
	}
}

func (handler *VerifyHandler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hash := r.PathValue("hash")

		err := handler.Service.VerifyEmail(hash)

		if err != nil {
			res.Json(w, ErrorResponse{
				Message: "Error during verification: " + err.Error(),
			}, http.StatusInternalServerError)
		} else {
			res.Json(w, VerificationResponse{
				Message: "Your email verified successfully",
			}, http.StatusOK)
		}
	}
}