package playlists

import (
	"context"
	"strings"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/slatermorgan/playlist-builder/pkg/spotify"
	"gopkg.in/go-playground/validator.v9"
)

var (
	validate *validator.Validate
)

const tracksPerArtist = 5

type SDK interface {
	GetArtistsTopTracks(artistID string) ([]*spotify.Track, error)
	SearchForArtist(q string) ([]*spotify.Artist, error)
	AddItemsToPlaylist(updatePlaylist spotify.UpdatePlaylist, userID string) error
	CreatePlaylist(playlistReq *spotify.Playlist, userID string) (*spotify.Playlist, error)
	GetPlaylist(playlistID string) (*spotify.Playlist, error)
	SetClientToken(string)
}

// Usecase for interacting with playlists
type Usecase struct {
	SDK SDK
}

// Create a single playlist
func (u *Usecase) Create(
	ctx context.Context,
	createPlaylist *CreatePlaylist,
	token string,
) (*spotify.Playlist, error) {
	u.SDK.SetClientToken(token)
	validate = validator.New()
	if err := validate.Struct(*createPlaylist); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		// Here we should create custom returns
		return nil, validationErrors
	}

	newPlaylist, err := u.SDK.CreatePlaylist(
		&spotify.Playlist{
			Name:        createPlaylist.Name,
			IsPublic:    createPlaylist.IsPublic,
			Description: createPlaylist.Description,
		},
		"userID",
	)
	if err != nil {
		return nil, errors.Wrap(nil, "failed to create playlist")
	}

	var playlistTrackURIs string
	for _, name := range createPlaylist.ArtistNames {
		artists, err := u.SDK.SearchForArtist(name)
		if err != nil {
			return nil, errors.Wrap(nil, "failed to search artist: "+name)
		}

		for _, artist := range artists {
			if artist.Name == name {
				topTracks, err := u.SDK.GetArtistsTopTracks(artist.ID)
				if err != nil {
					return nil, errors.Wrap(nil, "failed to search artist: "+name)
				}

				for i := 0; i < tracksPerArtist; i++ {
					playlistTrackURIs += topTracks[i].URI + ","
				}
			}
		}
		playlistTrackURIs = strings.TrimRight(playlistTrackURIs, ",")
	}

	err = u.SDK.AddItemsToPlaylist(
		spotify.UpdatePlaylist{
			URIs: playlistTrackURIs,
		},
		newPlaylist.ID,
	)
	if err != nil {
		return nil, errors.Wrap(nil, "failed to add items to playlist ID: "+newPlaylist.ID)
	}

	newPlaylist, err = u.SDK.GetPlaylist(newPlaylist.ID)
	if err != nil {
		return nil, errors.Wrap(nil, "failed to get playlist ID: "+newPlaylist.ID)
	}

	return newPlaylist, nil
}

func (u *Usecase) newID() string {
	uid := uuid.New()
	return uid.String()
}
