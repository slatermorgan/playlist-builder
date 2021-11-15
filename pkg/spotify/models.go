package spotify

type ArtistTopTracks struct {
	tracks []*Track
}

type SearchResults struct {
	artists ArtistsSearchResults
}

type ArtistsSearchResults struct {
	Items       []*Artist `json:"items"`
	Limit       int       `json:"limit"`
	NextURL     string    `json:"limit"`
	previousURL string    `json:"previous"`
	Total       int       `json:"total"`
	Offset      int       `json:"offset"`
}

type Artist struct {
	ID         string `json:"id"`
	HRef       string `json:"href"`
	URI        string `json:"uri"`
	Popularity int64  `json:"popularity"`
}

type Track struct {
	ID         string `json:"id"`
	URI        string `json:"uri"`
	Name       string `json:"name"`
	Popularity int64  `json:"popularity"`
	HRef       string `json:"href"`
	IsPlayable bool   `json:"is_playable"`
}

type Playlist struct {
	ID          string `json:"id"`
	URI         string `json:"uri"`
	Name        string `json:"name"`
	HRef        string `json:"href"`
	IsPublic    bool   `json:"public"`
	Description string `json:"description"`
}

type UpdatePlaylist struct {
	URIs string `json:"uris"`
}
