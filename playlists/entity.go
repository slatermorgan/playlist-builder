package playlists

// Playlist -
type Playlist struct {
	ID         string `json:"id"`
	TimeStart  int    `json:"timeStart" validate:"required,numeric"`
	TimeEnd    int    `json:"timeEnd" validate:"required,numeric"`
	ClientID   string `json:"clientID" validate:"required"`
	AssignedTo string `json:"assignedTo"`
}

// UpdatePlaylist -
type UpdatePlaylist struct {
	TimeStart  int    `json:"timeStart" validate:"required,numeric"`
	TimeEnd    int    `json:"timeEnd" validate:"required,numeric"`
	ClientID   string `json:"clientID" validate:"required"`
	AssignedTo string `json:"assignedTo"`
}
