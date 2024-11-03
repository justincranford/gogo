package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	mode := flag.String("mode", "server", "Mode to run: 'server' or 'client'")
	hostname := flag.String("host", "localhost", "Hostname to bind or connect to")
	port := flag.Int("port", 8080, "Port to bind or connect to")
	path := flag.String("path", "/", "Path to request (for client mode)")

	flag.Parse()

	log.Printf("Starting\n- mode: %s\n- hostname: %s\n- port: %d\n- path: %s", *mode, *hostname, *port, *path)

	if *mode == "server" {
		go func() {
			sigChannel := make(chan os.Signal, 1)
			signal.Notify(sigChannel, os.Interrupt, syscall.SIGTERM)
			<-sigChannel
			log.Println("Received shutdown signal, exiting...")
			os.Exit(0)
		}()
		startServer(*hostname, *port)
	} else if *mode == "client" {
		startClient(*hostname, *port, *path)
	} else {
		log.Fatalf("Invalid mode: %s. Use 'server' or 'client'.\n", *mode)
	}
}

func startServer(hostname string, port int) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/":
			_, err := fmt.Fprintf(w, "Welcome to the main page!")
			if err != nil {
				log.Printf("Failed to handle path %s: %v\n", r.URL.Path, err)
				return
			}
		case "/helloworld":
			_, err := fmt.Fprintf(w, "Hello, World!")
			if err != nil {
				log.Printf("Failed to handle path %s: %v\n", r.URL.Path, err)
				return
			}
		case "/info":
			_, err := fmt.Fprintf(w, "This is the info page.")
			if err != nil {
				log.Printf("Failed to handle path %s: %v\n", r.URL.Path, err)
				return
			}
		case "/shutdown":
			_, err := fmt.Fprintf(w, "Server shutting down...")
			if err != nil {
				log.Printf("Failed to handle path %s: %v\n", r.URL.Path, err)
				return
			}
			go func() {
				time.Sleep(1 * time.Second)
				os.Exit(0)
			}()
		default:
			http.NotFound(w, r) // Return a 404 for unrecognized paths
		}
	})

	address := fmt.Sprintf("%s:%d", hostname, port)
	log.Printf("Starting server at %s\n", address)
	if err := http.ListenAndServe(address, nil); err != nil {
		log.Fatalf("Server failed: %s\n", err)
	}
}

func startClient(hostname string, port int, path string) {
	url := fmt.Sprintf("http://%s:%d%s", hostname, port, path)
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Failed to make GET request to %s: %v\n", url, err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("Failed to close GET request to %s: %v\n", url, err)
			return
		}
	}(resp.Body)

	body, _ := io.ReadAll(resp.Body)
	log.Printf("Response from %s:\n%s\n", path, string(body))
}
