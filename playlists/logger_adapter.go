package playlists

import (
	"context"

	"github.com/slatermorgan/playlist-builder/pkg/spotify"
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

// Create a single playlist
func (a *LoggerAdapter) Create(ctx context.Context, playlist *CreatePlaylist, token string) (*spotify.Playlist, error) {
	defer a.Logger.Sync()
	a.Logger.Info("creating a single playlist")
	createdPlaylist, err := a.Usecase.Create(ctx, playlist, token)
	a.logErr(err)

	return createdPlaylist, err
}
