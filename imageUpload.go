package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"

	vision "cloud.google.com/go/vision/apiv1"
	"golang.org/x/net/context"
)

// ImageUploadHandler is called from /upload
func ImageUploadHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	fmt.Println("upload")

	decoder := base64.NewDecoder(base64.StdEncoding, r.Body)

	// Creates a client.
	client, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	image, err := vision.NewImageFromReader(decoder)
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
