package main

import (
	"bufio"
	"log"
	"os"
)

var ProgName = "netmet"

func main() {
	// handle args
	if len(os.Args) < 2 {
		log.Fatalf("%s: no filename provided, halting", ProgName)
	}
	var filename = os.Args[0]

	// open file or fail
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("%s nable to open %s", ProgName, filename)
	}
	defer file.Close()


	_, err = reporter(bufio.NewReader(file))
	// send to reporter
}

type NetMet struct {
	Server struct {
		Machine   string `json:"machine"`
		Locations struct {
			City    string `json:"city"`
			Country string `json:"country"`
		} `json:"locations"`
	} `json:"Server"`
	Download struct {
		ConnectionInfo struct {
			Client string `json:"Client"`
			Server string `json:"Server"`
			UUID   string `json:"UUID"`
		} `json:"ConnectionInfo"`
	} `json:"Download"`
	Upload struct {
		ConnectionInfo struct {
			Client string `json:"Client"`
			Server string `json:"Server"`
			UUID   string `json:"UUID"`
		} `json:"ConnectionInfo"`
	} `json:"Upload"`
	DownloadMsm []struct {
		ElapsedTime    float64 `json:"ElapsedTime"`
		NumBytes       int     `json:"NumBytes"`
		MeanClientMbps float64 `json:"MeanClientMbps"`
	} `json:"download_msm"`
	UploadMsm []struct {
		ElapsedTime    float64 `json:"ElapsedTime"`
		NumBytes       int     `json:"NumBytes"`
		MeanClientMbps float64 `json:"MeanClientMbps"`
	} `json:"upload_msm"`
	Timestamp int64 `json:"timestamp"`
}

func reporter(file *bufio.Reader) (int, error) {
	return 0, nil//
}
