package helpers

import (
	"bytes"
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
	"second/pkg/models"
)

func BadRequest(w http.ResponseWriter, logger *zap.Logger, err error) {
	logger.Info("", zap.Error(err))
	http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
}

func Forbidden(w http.ResponseWriter, logger *zap.Logger, err error) {
	logger.Info("", zap.Error(err))
	http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
}

func InternalServerError(w http.ResponseWriter, logger *zap.Logger, err error) {
	logger.Error("", zap.Error(err))
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func NotFound(w http.ResponseWriter, logger zap.Logger, err error) {
	logger.Info("", zap.Error(err))
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}

func SendAnswer(w http.ResponseWriter, msg string) error {
	answer := models.Answer{Answer: msg}
	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(answer)
	if err != nil {
		return err
	}
	_, err = w.Write(buf.Bytes())
	if err != nil {
		return err
	}
	return nil
}
