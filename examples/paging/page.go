package main

import (
	"context"
	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"golang.org/x/oauth2/clientcredentials"
	"log"
	"os"
)

func main() {
	ctx := context.Background()
	config := &clientcredentials.Config{
		ClientID:     os.Getenv("SPOTIFY_ID"),
		ClientSecret: os.Getenv("SPOTIFY_SECRET"),
		TokenURL:     spotifyauth.TokenURL,
	}
	token, err := config.Token(ctx)
	if err != nil {
		log.Fatalf("couldn't get token: %v", err)
	}

	httpClient := spotifyauth.New("").Client(ctx, token)
	client := spotify.New(spotify.HTTPClientOpt(httpClient))

	tracks, err := client.GetPlaylistTracks(ctx, "37i9dQZF1DWWzVPEmatsUB")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Playlist has %d total tracks", tracks.Total)
	for page := 1; ; page++ {
		log.Printf("  Page %d has %d tracks", page, len(tracks.Tracks))
		err = client.NextPage(ctx, tracks)
		if err == spotify.ErrNoMorePages {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
	}
}
