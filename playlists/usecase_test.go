package playlists

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCanGetPlaylist(t *testing.T) {
	expected := &Playlist{Name: "Ewan"}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := NewMockrepository(ctrl)
	repo.EXPECT().Get(context.Background(), "abc123").Return(expected, nil)

	uc := Usecase{repo}

	playlist, err := uc.Get(context.Background(), "abc123")

	assert.NoError(t, err)
	assert.Equal(t, expected, playlist)
}

func TestCanGetAllPlaylists(t *testing.T) {
	expected := []*Playlist{
		&Playlist{Name: "test1", Age: 1},
		&Playlist{Name: "test2", Age: 2},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := NewMockrepository(ctrl)
	repo.EXPECT().GetAll(context.Background()).Return(expected, nil)

	uc := Usecase{repo}

	playlists, err := uc.GetAll(context.Background())
	assert.NoError(t, err)
	assert.Len(t, playlists, 2)
	assert.Equal(t, expected, playlists)
}

func TestCanCreatePlaylist(t *testing.T) {
	expected := &Playlist{
		Name:  "testing",
		Email: "test@test.com",
		Age:   30,
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := NewMockrepository(ctrl)
	repo.EXPECT().Create(context.Background(), expected).Return(nil)

	uc := Usecase{repo}
	err := uc.Create(context.Background(), expected)

	assert.NoError(t, err)
}

func TestCanValidatePlaylist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := NewMockrepository(ctrl)

	uc := Usecase{repo}

	playlists := []*Playlist{
		&Playlist{},                      // No required fields
		&Playlist{Name: "", Age: 0},      // Blank name
		&Playlist{Name: "123", Age: 200}, // Integers as name, age too high
		&Playlist{Email: "nope"},
	}
	for _, val := range playlists {
		err := uc.Create(context.Background(), val)
		assert.Error(t, err)
	}
}

func TestCanUpdatePlaylist(t *testing.T) {
	playlist := &UpdatePlaylist{
		Name:  "new name",
		Email: "test@test.com",
		Age:   20,
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := NewMockrepository(ctrl)
	repo.EXPECT().Update(context.Background(), "abc123", playlist).Return(nil)
	uc := Usecase{repo}
	err := uc.Update(context.Background(), "abc123", playlist)
	assert.NoError(t, err)
}

func TestCanDeletePlaylist(t *testing.T) {
	playlist := &Playlist{
		ID: "abc123",
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := NewMockrepository(ctrl)
	repo.EXPECT().Delete(context.Background(), playlist.ID).Return(nil)
	uc := Usecase{repo}
	err := uc.Delete(context.Background(), playlist.ID)
	assert.NoError(t, err)
}
