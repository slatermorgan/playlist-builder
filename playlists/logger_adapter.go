package playlists

import (
	"context"

	"go.uber.org/zap"
)

// LoggerAdapter wraps the usecase interface
// with a logging adapter which can be swapped out
type LoggerAdapter struct {
	Logger  *zap.Logger
	Usecase PlaylistService
}

func (a *LoggerAdapter) logErr(err error) {
	if err != nil {
		a.Logger.Error(err.Error())
	}
}

// Get a single playlist
func (a *LoggerAdapter) Get(ctx context.Context, id string) (*Playlist, error) {
	defer a.Logger.Sync()
	a.Logger.With(zap.String("id", id))
	a.Logger.Info("getting a single playlist")
	playlist, err := a.Usecase.Get(ctx, id)
	a.logErr(err)
	return playlist, err
}

// GetAll gets all playlists
func (a *LoggerAdapter) GetAll(ctx context.Context) ([]*Playlist, error) {
	defer a.Logger.Sync()
	a.Logger.Info("getting all playlists")
	playlists, err := a.Usecase.GetAll(ctx)
	a.logErr(err)
	return playlists, err
}

// Update a single playlist
func (a *LoggerAdapter) Update(ctx context.Context, id string, playlist *UpdatePlaylist) error {
	defer a.Logger.Sync()
	a.Logger.With(zap.String("id", id))
	a.Logger.Info("updating a single playlist")
	err := a.Usecase.Update(ctx, id, playlist)
	a.logErr(err)
	return err
}

// Create a single playlist
func (a *LoggerAdapter) Create(ctx context.Context, playlist *Playlist) error {
	defer a.Logger.Sync()
	a.Logger.Info("creating a single playlist")
	err := a.Usecase.Create(ctx, playlist)
	a.logErr(err)
	return err
}

// Delete a single playlist
func (a *LoggerAdapter) Delete(ctx context.Context, id string) error {
	defer a.Logger.Sync()
	a.Logger.Info("deleting a single playlist")
	err := a.Usecase.Delete(ctx, id)
	a.logErr(err)
	return err
}
