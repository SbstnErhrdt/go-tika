package tika_service

import (
	"context"
	"github.com/google/go-tika/tika"
	"log"
	"os"
)

func Start() {

	err := tika.DownloadServer(context.Background(), "1.21", "tika-server-1.21.jar")
	if err != nil {
		log.Fatal(err)
	}

	// Optionally pass a port as the second argument.
	s, err := tika.NewServer("tika-server-1.21.jar", "")
	if err != nil {
		log.Fatal(err)
	}
	err = s.Start(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer s.Stop()

	// import "os"
	f, err := os.Open("./test-data/80315_HRA103253_UUS_R_2015-06-23_24995452_38243.tiff")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	client := tika.NewClient(nil, s.URL())

	det, err := client.Detectors(context.Background())
	log.Println(det)

	body, err := client.
		fmt.Println(body)

}
