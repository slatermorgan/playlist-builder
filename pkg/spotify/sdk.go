package spotify

type silly interface {
	SearchForArtist(q string) ([]*Artist, error)
	GetArtistsTopTrack(artistName string) ([]*Track, error)
	CreatePlaylist() error
	AddTrackToPlaylist() error
}

type SDK struct {
	client Client
}

func (s *SDK) GetArtistsTopTrack(artistID string) ([]*Track, error) {
	var artistTopTracks *ArtistTopTracks

	s.client.Get(artistTopTracks, "/artists/"+artistID+"/top-tracks")

	return artistTopTracks.tracks, nil
}

func (s *SDK) SearchForArtist(q string) ([]*Artist, error) {
	var searchResults *SearchResults

	s.client.Get(searchResults, "/search/?type=artist&q="+q)

	return searchResults.artists.Items, nil
}
