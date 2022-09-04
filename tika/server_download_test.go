package tika

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"
	"time"
)

func TestDownloadServer(t *testing.T) {
	err := DownloadServer(context.Background(), "2.4.1", "tika-server-2.4.1.jar")
	if err != nil {
		t.Fatalf("can not download server: %v", err)
	}
}

func TestServer_Start(t *testing.T) {
	// Optionally pass a port as the second argument.
	s, err := NewServer("tika-server-2.4.1.jar", "")
	if err != nil {
		t.Fatalf("can not init server: %v", err)
	}
	err = s.Start(context.Background())
	if err != nil {
		t.Fatalf("can not start server: %v", err)
	}
	time.Sleep(time.Second * 20)
	err = s.Stop()
	if err != nil {
		t.Fatalf("can not stop server: %v", err)
	}
}

func TestNewClient(t *testing.T) {
	// Optionally pass a port as the second argument.
	s, err := NewServer("tika-server-2.4.1.jar", "")
	if err != nil {
		t.Fatalf("can not init server: %v", err)
	}
	err = s.Start(context.Background())
	if err != nil {
		t.Fatalf("can not start server: %v", err)
	}
	defer s.Stop()

	// import "os"
	f, err := os.Open("./test-data/test.pdf")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	client := NewClient(nil, s.URL())
	body, err := client.Parse(context.Background(), f)
	fmt.Println(body)

}
