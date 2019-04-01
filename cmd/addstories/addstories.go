package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"cloud.google.com/go/firestore"
)

func readStories(filename string) []interface{} {

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("failed to read stories: %v", err)
	}

	stories := make([]interface{}, 0, 0)
	err = json.Unmarshal(content, &stories)
	if err != nil {
		log.Fatalf("failed to parse stories: %v", err)
	}

	return stories
}

func main() {

	if len(os.Args) != 2 {
		log.Fatalf("usage: addstories path/to/stories.json")
	}

	stories := readStories(os.Args[1])

	ctx := context.Background()

	projectID := "poc-ns-api"
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("failed to create firestore client: %v", err)
	}
	defer client.Close()

	for _, story := range stories {
		_, _, err = client.Collection("stories").Add(ctx, story)
		if err != nil {
			log.Fatalf("failed to add story: %v", err)
		}
	}
}
