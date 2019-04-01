package main

import (
	"context"
	"encoding/json"
	"log"

	"cloud.google.com/go/firestore"
)

var oneStory = `  {
    "url": "",
    "mp3Url": "https://s3-us-west-1.amazonaws.com/nativestories/media/Hiiaka.mp3",
    "mp3Duration": 467,
    "imageUrls": [
      "https://s3-us-west-1.amazonaws.com/nativestories/images/Princess+Ruths+estate+nawahi+place.png"
    ],
    "storyTeller": {
      "name": "Hi\u02BBohia"
    },
    "name": "Saved by Hi\u02BBiaka",
    "description": "Location: K\u012Blauea Cove, at the end of Keawa\u02BBula Bay\n\nTo Purchase books: Kamehameha Publishing, www.kamehamehapublishing.org \n\nName of Series: The Ka\u02BBena Aloha Series\n\nDate Recorded: 9/29/2018\n\nDescription: \nMalolo kai \u0113! Malolo kai \u0113! If you hear the call, \"Low tide! Low tide!\" while passing through the valley of M\u0101kaha, run.  Run as fast as you can!\nE ka pua o ka \u02BBilima e,\nH\u014Dmai ana ho\u02BBi he ola\nE M\u0101kua i ka nu\u02BBa o ke kai-e\nHa\u02BBawi mai ana ho\u02BBi ua ola-e\nE ola ku\u02BBu kama i ka hu\u02BBa o ke kai-e\nA ola ho\u02BBi i\u0101 K\u0101ne i ka wai ola-e\n\nThis is among one of the many chants that was composed by Hi\u02BBiakaikapoliopele, a beloved Hawaiian goddess.  With this particular chant, she revived a young girl, giving her a second chance at life...\n\nBook Design: D. Ka\u02BBiulani Kauihou\nProject Directors: D. Ka\u02BBiulani Kauihou & Leilani Kapoi\nAuthor: A.Ka. Mahiai in the Hoku o Hawaii newspaper on November 23, 1926\nPhoto credits: Pono Chang\nShout outs: The Hawaiian people of Wai\u02BBanae, Board of Water Supply, Ulukau, CZM Hawai\u02BBi\nPublisher: Penmar Publishing\nVoice: Kiele Gonzalez",
    "place": "K\u012Blauea Cove",
    "location": {
      "address": "",
      "city": "Honolulu",
      "ahupuaa": "",
      "latitude": 21.556022,
      "longitude": -158.248486
    },
    "tags": [
      {
        "type": "type",
        "name": "Location"
      },
      {
        "type": "category",
        "name": "Ancient"
      },
      {
        "type": "category",
        "name": "New Arrivals"
      },
      {
        "type": "category",
        "name": "English"
      }
    ]
  }`

func main() {

	ctx := context.Background()

	projectID := "poc-ns-api"
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("failed to create firestore client: %v", err)
	}
	defer client.Close()

	story := make(map[string]interface{})
	err = json.Unmarshal([]byte(oneStory), &story)
	if err != nil {
		log.Fatalf("can't parse story: %v", err)
	}

	_, _, err = client.Collection("stories").Add(ctx, story)
	if err != nil {
		log.Fatalf("failed to add story: %v", err)
	}
}
