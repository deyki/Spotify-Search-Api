package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)


type ExternalUrls struct {
     Spotify	string
}


type Artist struct {
     Name		string
     ExternalUrls 	ExternalUrls
}


type Album struct {
     Name			string
     Artists			[]Artist
     ExternalUrls		ExternalUrls
}


type Track struct {
     Name			string
     Href			string
     Popularity			int
     Album			Album
     ExternalUrl		ExternalUrls
}


type ApiResponse struct {
     Tracks		Items[Track]
}


type Items[T any] struct {
     Items []T
}


func printTracks(tracks []Track) {
	for _, track := range tracks {
		for _, artist := range track.Album.Artists {
			fmt.Printf("%s\n%s\n%s\n%s\n", track.Name, track.Album.Name, artist.Name, track.ExternalUrl.Spotify)
		}
	}
}


func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter search query: ")
	scanner.Scan()
	searchQuery := scanner.Text()

	fmt.Print("Enter authentication token: ")
	scanner.Scan()
	authToken := scanner.Text()

	url := "https://api.spotify.com/v1/search?q=" + searchQuery + "&type=track,artist"

	client := http.Client{}
	request, _ := http.NewRequest("GET", url, nil)

	request.Header.Set("Authorization", "Bearer " + authToken)
	request.Header.Set("CONTENT_TYPE", "application/json")
	request.Header.Set("ACCEPT", "application/json")

	response, _ := client.Do(request)

	switch response.StatusCode {
	case 200:
		var apiResponse ApiResponse
		json.NewDecoder(response.Body).Decode(&apiResponse)
		printTracks(apiResponse.Tracks.Items)

	case 401:
		fmt.Println("Unauthorized, try another auth token!")

	default:
		fmt.Println("Unexpected error!")
	}
}
