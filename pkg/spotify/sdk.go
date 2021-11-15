package spotify

type SDK struct {
	client Client
}

func New(config *Config) *SDK {
	client := Client{config}
	return &SDK{client: client}
}

func (s *SDK) GetArtistsTopTracks(artistID string) ([]*Track, error) {
	var artistTopTracks *ArtistTopTracks

	err := s.client.Get(artistTopTracks, "/artists/"+artistID+"/top-tracks")
	if err != nil {
		return nil, err
	}

	return artistTopTracks.tracks, nil
}

func (s *SDK) SearchForArtist(q string) ([]*Artist, error) {
	var searchResults *SearchResults

	err := s.client.Get(searchResults, "/search/?type=artist&q="+q)
	if err != nil {
		return nil, err
	}

	return searchResults.artists.Items, nil
}

func (s *SDK) AddItemsToPlaylist(updatePlaylist UpdatePlaylist, playlistID string) error {
	err := s.client.Post(
		updatePlaylist,
		nil,
		"/users/"+playlistID+"/playlists",
	)

	return err
}

func (s *SDK) CreatePlaylist(playlistReq *Playlist, userID string) (*Playlist, error) {
	var playlistRes *Playlist

	err := s.client.Post(
		playlistReq,
		playlistRes,
		"/playlists/"+userID+"/tracks",
	)
	if err != nil {
		return nil, err
	}

	return playlistRes, nil
}
