package playlists

import (
	"context"

	"github.com/slatermorgan/playlist-builder/pkg/spotify"
	"go.uber.org/zap"
)

// UseService is the top level signature of this service
type PlaylistService interface {
	Get(ctx context.Context, id string) (*spotify.Playlist, error)
	GetAll(ctx context.Context) ([]*spotify.Playlist, error)
	Update(ctx context.Context, id string, playlist *spotify.UpdatePlaylist) error
	Create(ctx context.Context, playlist *CreatePlaylist) (*spotify.Playlist, error)
	Delete(ctx context.Context, id string) error
}

// Init sets up an instance of this domains
// usecase, pre-configured with the dependencies.
func Init(integration bool) (PlaylistService, error) {
	config := &spotify.Config{
		BaseURL: "https://api.spotify.com/v1",
	}
	client := spotify.New(config)
	logger, _ := zap.NewProduction()

	usecase := &LoggerAdapter{
		Logger:  logger,
		Usecase: &Usecase{Client: client},
	}
	return usecase, nil
}
