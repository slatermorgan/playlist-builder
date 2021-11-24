package http

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/slatermorgan/playlist-builder/playlists"
)

const fiveSecondsTimeout = time.Second * 5

type delivery struct {
	usecase playlists.PlaylistService
}

func writeErr(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(err.Error()))
}

func writeInvalidMethod(w http.ResponseWriter) {
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte("method not supported"))
}

func (d *delivery) Get(w http.ResponseWriter, r *http.Request) {
	writeInvalidMethod(w)
	return
}

func (d *delivery) GetAll(w http.ResponseWriter, r *http.Request) {
	writeInvalidMethod(w)
	return
}

func (d *delivery) Update(w http.ResponseWriter, r *http.Request) {
	writeInvalidMethod(w)
	return
}

func (d *delivery) Create(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), fiveSecondsTimeout)
	defer cancel()

	decoder := json.NewDecoder(r.Body)
	playlist := &playlists.CreatePlaylist{}
	if err := decoder.Decode(&playlist); err != nil {
		writeErr(w, err)
		return
	}

	createdPlaylist, err := d.usecase.Create(ctx, playlist)
	if err != nil {
		writeErr(w, err)
		return
	}

	data, err := json.Marshal(createdPlaylist)
	if err != nil {
		writeErr(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (d *delivery) Delete(w http.ResponseWriter, r *http.Request) {
	writeInvalidMethod(w)
	return
}

// Routes -
func Routes() (*mux.Router, error) {
	usecase, err := playlists.Init(true)
	if err != nil {
		log.Panic(err)
	}

	delivery := &delivery{usecase}

	r := mux.NewRouter()
	r.HandleFunc("/playlists", delivery.Create).Methods("POST")
	r.HandleFunc("/playlists", delivery.GetAll).Methods("GET")
	r.HandleFunc("/playlists/{id}", delivery.Get).Methods("GET")
	r.HandleFunc("/playlists/{id}", delivery.Update).Methods("PUT")
	r.HandleFunc("/playlists/{id}", delivery.Delete).Methods("DELETE")

	return r, nil
}
