package playlists

import (
	"context"

	"github.com/slatermorgan/playlist-builder/pkg/spotify"
	"go.uber.org/zap"
)

// UseService is the top level signature of this service
type PlaylistService interface {
	Create(ctx context.Context, playlist *CreatePlaylist, token string) (*spotify.Playlist, error)
}

// Init sets up an instance of this domains
// usecase, pre-configured with the dependencies.
func Init(integration bool) (PlaylistService, error) {
	config := &spotify.Config{
		BaseURL: "https://api.spotify.com/v1",
	}
	sdk := spotify.New(config)
	logger, _ := zap.NewProduction()

	usecase := &LoggerAdapter{
		Logger:  logger,
		Usecase: &Usecase{SDK: sdk},
	}
	return usecase, nil
}
