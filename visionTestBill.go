// detectText gets text from the Vision API for an image at the given file path.
package main
import (
	"os"
        "fmt"
	// Imports the Google Cloud Vision API client package.
	vision "cloud.google.com/go/vision/apiv1"
	"golang.org/x/net/context"
)
func main() {
        fmt.Println(detectText())
}

func detectText() error {
        ctx := context.Background()

        client, err := vision.NewImageAnnotatorClient(ctx)
        if err != nil {
                return err
        }
        filename := "testdata/test-bill.jpg"

        f, err := os.Open(filename)
        println(filename)
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
                println("No text found.")
        } else {
                println("Text:")
                for _, annotation := range annotations {
                        fmt.Println("%q\n", annotation.Description)
                }
        }

        return nil
}