package api

import (
	"fmt"
	"net/http"
	"receipt-processor/service"
	"receipt-processor/storage"

	"github.com/google/uuid"
)

type GetReceiptPointsHandler struct {
	getReceiptPointsUseCase service.GetReceiptPointsUseCase
	responseWriter          ResponseWriter
}

func NewGetReceiptPointsHandler(getReceiptPointsUseCase service.GetReceiptPointsUseCase, responseWriter ResponseWriter) *GetReceiptPointsHandler {
	return &GetReceiptPointsHandler{
		getReceiptPointsUseCase: getReceiptPointsUseCase,
		responseWriter:          responseWriter,
	}
}

func (h *GetReceiptPointsHandler) Handle(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("Request Received - /receipts/{id}/points")
	if r.Method != http.MethodGet {
		return storage.ErrNotFound
	}
	rid := r.PathValue("id")
	id, err := uuid.Parse(rid)
	if err != nil {
		return err
	}
	points, err := h.getReceiptPointsUseCase.GetReceiptPoints(id)
	if err != nil {
		return err
	}

	return h.responseWriter.WriteJSONResponse(w, http.StatusOK, PointsResponse{points})
}

func (h *GetReceiptPointsHandler) RegisterRoutes() {
	fmt.Println("Registering route - GET /receipts/{id}/points")
	http.HandleFunc("/receipts/{id}/points", func(w http.ResponseWriter, r *http.Request) {
		err := h.Handle(w, r)
		if err != nil {
			fmt.Println(err)
			http.Error(w, storage.ErrNotFound.Error(), http.StatusNotFound)
		}
	})
}

type PointsResponse struct {
	Points int `json:"points"`
}
