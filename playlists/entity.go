package playlists

// Playlist -
type CreatePlaylist struct {
	Name        string   `json:"name" validate:"required"`
	IsPublic    bool     `json:"isPublic"`
	Description string   `json:"description"`
	ArtistNames []string `json:"artistNames" validate:"required"`
}
