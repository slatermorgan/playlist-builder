package http

import (
	"context"
	"encoding/json"
	"go-care/shifts"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/slatermorgan/go-care/playlists"
)

const fiveSecondsTimeout = time.Second * 5

type delivery struct {
	usecase playlists.ShiftService
}

func writeErr(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(err.Error()))
}

func (d *delivery) Get(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), fiveSecondsTimeout)
	defer cancel()

	vars := mux.Vars(r)
	id := vars["id"]

	shift, err := d.usecase.Get(ctx, id)
	if err != nil {
		writeErr(w, err)
		return
	}

	data, err := json.Marshal(shift)
	if err != nil {
		writeErr(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (d *delivery) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), fiveSecondsTimeout)
	defer cancel()

	shifts, err := d.usecase.GetAll(ctx)
	if err != nil {
		writeErr(w, err)
		return
	}

	data, err := json.Marshal(shifts)
	if err != nil {
		writeErr(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (d *delivery) Update(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), fiveSecondsTimeout)
	defer cancel()

	decoder := json.NewDecoder(r.Body)
	shift := &shifts.UpdateShift{}
	if err := decoder.Decode(&shift); err != nil {
		writeErr(w, err)
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]

	if err := d.usecase.Update(ctx, id, shift); err != nil {
		writeErr(w, err)
		return
	}
}

func (d *delivery) Create(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), fiveSecondsTimeout)
	defer cancel()

	decoder := json.NewDecoder(r.Body)
	shift := &shifts.Shift{}
	if err := decoder.Decode(&shift); err != nil {
		writeErr(w, err)
		return
	}

	if err := d.usecase.Create(ctx, shift); err != nil {
		writeErr(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Ok"))
}

func (d *delivery) Delete(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), fiveSecondsTimeout)
	defer cancel()

	vars := mux.Vars(r)
	id := vars["id"]

	if err := d.usecase.Delete(ctx, id); err != nil {
		writeErr(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("Deleted"))
}

// Routes -
func Routes() (*mux.Router, error) {
	usecase, err := shifts.Init(true)
	if err != nil {
		log.Panic(err)
	}

	delivery := &delivery{usecase}

	r := mux.NewRouter()
	r.HandleFunc("/shifts", delivery.Create).Methods("POST")
	r.HandleFunc("/shifts", delivery.GetAll).Methods("GET")
	r.HandleFunc("/shifts/{id}", delivery.Get).Methods("GET")
	r.HandleFunc("/shifts/{id}", delivery.Update).Methods("PUT")
	r.HandleFunc("/shifts/{id}", delivery.Delete).Methods("DELETE")

	return r, nil
}
