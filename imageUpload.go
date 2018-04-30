package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	vision "cloud.google.com/go/vision/apiv1"
	"golang.org/x/net/context"
)

//BillData represents the cleaned data from an image upload
type BillData struct {
	HospitalName       string `json:"hospitalName"`
	TreatmentDate      string `json:"treatmentDate"`
	MedicalCode        string `json:"medicalCode"`
	ServiceDescription string `json:"serviceDescription"`
	Balance            string `json:"balance"`
}

// ImageUploadHandler is called from /upload
func ImageUploadHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	fmt.Println("upload")

	decoder := base64.NewDecoder(base64.StdEncoding, r.Body)

	billData, err := detectText(ctx, decoder)
	if err != nil {
		fmt.Println(err)
	}

	jData, err := json.Marshal(billData)
	if err != nil {
		fmt.Print(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Write(jData)
}

func detectText(ctx context.Context, reader io.Reader) (*BillData, error) {
	build := ""
	_ = build

	testBill := BillData{HospitalName: "", TreatmentDate: "", MedicalCode: "", ServiceDescription: "", Balance: ""}
	client, err := vision.NewImageAnnotatorClient(ctx)

	if err != nil {
		return nil, err
	}

	image, err := vision.NewImageFromReader(reader)
	if err != nil {
		if err != nil {
			return nil, err
		}
	}
	annotations, err := client.DetectTexts(ctx, image, nil, 100)
	if err != nil {
		return nil, err
	}

	if len(annotations) == 0 {
		println("No text found.")
	} else {
		for index, annotation := range annotations {
			//fmt.Println("\n", annotation.Description)
			if annotation.Description == "to:" {
				for index2, annotation2 := range annotations {
					if index2 == index+1 {
						build = annotation2.Description
					}
					if index2 == index+2 {
						build += " " + annotation2.Description
					}
					if index2 == index+3 {
						build += " " + annotation2.Description
					}
					if index2 == index+4 {
						build += " " + annotation2.Description
					}
					testBill.HospitalName = build
				}
			}
			if annotation.Description == "Service:" {
				for index2, annotation2 := range annotations {
					if index2 == index+1 {
						testBill.TreatmentDate = annotation2.Description
					}
				}
			}
			if annotation.Description == "Code:" {
				for index2, annotation2 := range annotations {
					if index2 == index+1 {
						testBill.MedicalCode = annotation2.Description
					}
				}
			}
			if annotation.Description == "Description:" {
				for index2, annotation2 := range annotations {
					if index2 == index+1 {
						testBill.ServiceDescription = annotation2.Description + "I"
					}
				}
			}
			if annotation.Description == "Due" {
				for index2, annotation2 := range annotations {
					if index2 == index+1 {
						testBill.Balance = annotation2.Description
					}
				}
			}
		}
	}

	return &testBill, nil
}
