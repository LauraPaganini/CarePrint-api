// detectText gets text from the Vision API for an image at the given file path.
package main
import (
	"os"
        "fmt"
	//"strings"
	"reflect"
	// Imports the Google Cloud Vision API client package.
	vision "cloud.google.com/go/vision/apiv1"
	"golang.org/x/net/context"
)
func main() {
	fmt.Println(in_array("sick", [sick]))
        fmt.Println(detectText())
}

func in_array(val interface{}, array interface{}) (exists bool) {
    exists = false
    index := -1
    _ = index

    switch reflect.TypeOf(array).Kind() {
    case reflect.Slice:
        s := reflect.ValueOf(array)

        for i := 0; i < s.Len(); i++ {
            if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
                index = i
                exists = true
                return
            }
        }
    }

    return
}

func detectText() error {
        ctx := context.Background()

        client, err := vision.NewImageAnnotatorClient(ctx)
        if err != nil {
                return err
        }
        filename := "testdata/bill-2.jpg"

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
        annotations, err := client.DetectTexts(ctx, image, nil, 100)
        if err != nil {
                return err
        }

        if len(annotations) == 0 {
                println("No text found.")
        } else {
                println("Text:")
                for _, annotation := range annotations {
                        //fmt.Println("\n", annotation.Description)
			if in_array("and", annotation) {
                		fmt.Println("yea boi")
			} else {
				fmt.Println("bruh")	
			}
		}
        }

        return nil
}
