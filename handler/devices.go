package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/themanojk/reflekt/models"
	"github.com/themanojk/reflekt/store"
)

type Device struct {
	store store.Store
}

func NewDevice(s store.Store) *Device {
	return &Device{store: s}
}

func (d *Device) Create(w http.ResponseWriter, r *http.Request) {
	var device models.Device
	if err := json.NewDecoder(r.Body).Decode(&device); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := d.store.Insert(r.Context(), &device)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}

func (d *Device) GetByMacAddress(w http.ResponseWriter, r *http.Request) {
	macAddress := chi.URLParam(r, "macAddress")

	res, err := d.store.GetByMacAddress(r.Context(), macAddress)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}
