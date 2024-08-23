package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"GoDataOpsAPI/internal/model"  // Update this line
    "GoDataOpsAPI/internal/store"   // Update this line
    "GoDataOpsAPI/pkg/response"     // Update this line
	
)

type ItemHandler struct {
	store *store.InMemoryStore
}

func NewItemHandler(store *store.InMemoryStore) *ItemHandler {
	return &ItemHandler{store: store}
}

func (h *ItemHandler) ItemsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		h.getItems(w, r)
	case "POST":
		h.createItem(w, r)
	default:
		response.JSONError(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

func (h *ItemHandler) ItemHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Path[len("/items/"):])
	if err != nil {
		response.JSONError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	switch r.Method {
	case "GET":
		h.getItem(w, r, id)
	case "PUT":
		h.updateItem(w, r, id)
	case "DELETE":
		h.deleteItem(w, r, id)
	default:
		response.JSONError(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

func (h *ItemHandler) getItems(w http.ResponseWriter, r *http.Request) {
	items := h.store.GetAll()
	response.JSON(w, http.StatusOK, items)
}

func (h *ItemHandler) createItem(w http.ResponseWriter, r *http.Request) {
	var item model.Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		response.JSONError(w, http.StatusBadRequest, "Invalid input")
		return
	}

	item = h.store.Create(item)
	response.JSON(w, http.StatusCreated, item)
}

func (h *ItemHandler) getItem(w http.ResponseWriter, r *http.Request, id int) {
	item, exists := h.store.GetByID(id)
	if !exists {
		response.JSONError(w, http.StatusNotFound, "Item not found")
		return
	}

	response.JSON(w, http.StatusOK, item)
}

func (h *ItemHandler) updateItem(w http.ResponseWriter, r *http.Request, id int) {
	var updatedItem model.Item
	if err := json.NewDecoder(r.Body).Decode(&updatedItem); err != nil {
		response.JSONError(w, http.StatusBadRequest, "Invalid input")
		return
	}

	item, updated := h.store.Update(id, updatedItem)
	if !updated {
		response.JSONError(w, http.StatusNotFound, "Item not found")
		return
	}

	response.JSON(w, http.StatusOK, item)
}

func (h *ItemHandler) deleteItem(w http.ResponseWriter, r *http.Request, id int) {
	if !h.store.Delete(id) {
		response.JSONError(w, http.StatusNotFound, "Item not found")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
