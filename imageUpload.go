package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	vision "cloud.google.com/go/vision/apiv1"
	"golang.org/x/net/context"
)

// ImageUploadHandler is called from /upload
func ImageUploadHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	fmt.Println("upload")
	body, err := ioutil.ReadAll(r.Body)

	fmt.Println(body)
	// Creates a client.
	client, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	image, err := vision.NewImageFromReader(r.Body)
	if err != nil {
		log.Fatalf("Failed to create image: %v", err)
	}

	labels, err := client.DetectLabels(ctx, image, nil, 10)
	if err != nil {
		log.Fatalf("Failed to detect labels: %v", err)
	}

	fmt.Println("Labels:")
	for _, label := range labels {
		fmt.Println(label.Description)
	}
}
