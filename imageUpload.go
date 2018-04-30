package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"os"
	"encoding/json"

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
func detectText() error {
        
    type billData struct {
          hospitalName string `json:"hospitalName"`
          treatmentDate string `json:"treatmentDate"`
          medicalCode string `json:"medicalCode"`
          serviceDescription string `json:"serviceDescription"`
          balance string `json:"balance"`
    }
    build := ""
    _ = build

    ctx := context.Background()
    testBill := billData{hospitalName: "", treatmentDate: "", medicalCode: "", serviceDescription: "", balance: ""}
    client, err := vision.NewImageAnnotatorClient(ctx)
            
    if err != nil {
        return err
    }

    filename := "CarePrint-api/testdata/realBill.jpg"
    
    f, err := os.Open(filename)
    println(filename)

    if err != nil {
        return err
    }

    defer f.Close()
    
    image, err := vision.NewImageFromReader(f)
    if err != nil {
        if err != nil {
            return err
        }
    }
    annotations, err := client.DetectTexts(ctx, image, nil, 100)
    if err != nil {
        return err
    }

    if len(annotations) == 0 {
        println("No text found.")
    } else {
        for index, annotation := range annotations {
            //fmt.Println("\n", annotation.Description)
            if (annotation.Description == "to:") {
                for index2, annotation2 := range annotations {
                    if index2 == index + 1 {
                        build = annotation2.Description
                    }
                    if index2 == index + 2 {
                        build += " " +  annotation2.Description
                    }
                    if index2 == index + 3 {
                        build += " " + annotation2.Description
                    }
                    if index2 == index + 4 {
                        build += " " + annotation2.Description
                    }
                    testBill.hospitalName = build
                }
            }
            if (annotation.Description == "Service:") {
                for index2, annotation2 := range annotations {
                        if index2 == index + 1 {
                            testBill.treatmentDate = annotation2.Description
                        }
                }
            }
            if (annotation.Description == "Code:") {
                for index2, annotation2 := range annotations {
                    if index2 == index + 1 {
                        testBill.medicalCode = annotation2.Description
                    }
                }
            }
            if (annotation.Description == "Description:") {
                for index2, annotation2 := range annotations {
                    if index2 == index + 1 {
                        testBill.serviceDescription = annotation2.Description + "I"
                    }
                }
            }
            if (annotation.Description == "Due") {
                for index2, annotation2 := range annotations {
                    if index2 == index + 1 {
                        testBill.balance = annotation2.Description
                    }
                }
            }
        }
    }
//      fmt.Println(testBill)
    b, err := json.Marshal(testBill)
    if err != nil {
        fmt.Println("error:", err)
    }
    fmt.Println(b)
    return  nil
}    
