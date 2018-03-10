package main

import (
        "fmt"
        "os"
        "io"
        // Imports the Google Cloud Vision API client package.
        vision "cloud.google.com/go/vision/apiv1"
        "golang.org/x/net/context"
)
func detectText(w io.Writer, file string) error {
        ctx := context.Background()

        client, err := vision.NewImageAnnotatorClient(ctx)
        if err != nil {
                return err
        }

        f, err := os.Open(file)
        if err != nil {
                return err
        }
        defer f.Close()

        image, err := vision.NewImageFromReader(f)
        if err != nil {
                return err
        }
        annotations, err := client.DetectTexts(ctx, image, nil, 10)
        if err != nil {
                return err
        }

        if len(annotations) == 0 {
                fmt.Fprintln(w, "No text found.")
        } else {
                fmt.Fprintln(w, "Text:")
                for _, annotation := range annotations {
                        fmt.Fprintf(w, "%q\n", annotation.Description)
                }
        }

        return nil
}
