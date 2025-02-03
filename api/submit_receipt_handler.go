package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"receipt-processor/domain"
	"receipt-processor/service"
)

var ErrBadRequest = errors.New("The receipt is invalid.")

type SubmitReceiptHandler struct {
	submitReceiptUseCase service.SubmitReceiptUseCase
	responseWriter       ResponseWriter
}

func NewSubmitReceiptHandler(submitReceiptUseCase service.SubmitReceiptUseCase, responseWriter ResponseWriter) *SubmitReceiptHandler {
	return &SubmitReceiptHandler{
		submitReceiptUseCase: submitReceiptUseCase,
		responseWriter:       responseWriter,
	}
}

func (h *SubmitReceiptHandler) Handle(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodPost {
		return ErrBadRequest
	}
	receipt := &domain.ReceiptDto{}
	err := json.NewDecoder(r.Body).Decode(receipt)
	if err != nil {
		return err
	}
	id, err := h.submitReceiptUseCase.SubmitReceipt(receipt)
	if err != nil {
		return err
	}

	return h.responseWriter.WriteJSONResponse(w, http.StatusOK, SubmitResponse{id.String()})
}

func (h *SubmitReceiptHandler) RegisterRoutes() {
	fmt.Println("Registering route - POST /receipts/process")
	http.HandleFunc("/receipts/process", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request Received - /receipts/process")
		err := h.Handle(w, r)
		if err != nil {
			fmt.Println(err)
			http.Error(w, ErrBadRequest.Error(), http.StatusBadRequest)
		}
	})
}

type SubmitResponse struct {
	ID string `json:"id"`
}
