package http

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/slatermorgan/playlist-builder/pkg/spotify"
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

func (d *delivery) Get(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), fiveSecondsTimeout)
	defer cancel()

	vars := mux.Vars(r)
	id := vars["id"]

	playlist, err := d.usecase.Get(ctx, id)
	if err != nil {
		writeErr(w, err)
		return
	}

	data, err := json.Marshal(playlist)
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

	playlists, err := d.usecase.GetAll(ctx)
	if err != nil {
		writeErr(w, err)
		return
	}

	data, err := json.Marshal(playlists)
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
	playlist := &spotify.UpdatePlaylist{}
	if err := decoder.Decode(&playlist); err != nil {
		writeErr(w, err)
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]

	if err := d.usecase.Update(ctx, id, playlist); err != nil {
		writeErr(w, err)
		return
	}
}

func (d *delivery) Create(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), fiveSecondsTimeout)
	defer cancel()

	decoder := json.NewDecoder(r.Body)
	playlist := &spotify.Playlist{}
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
	w.Write([]byte("method not found"))
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
