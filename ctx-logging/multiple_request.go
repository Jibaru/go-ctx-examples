package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
)

const baseURL = "http://localhost:8080/api/pokemon"

func makeRequest(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	if id%2 == 0 {
		id = -id
	}

	reqID := fmt.Sprintf("%v-%v", id, strconv.FormatInt(time.Now().UnixNano(), 10))

	url := fmt.Sprintf("%s/%d", baseURL, id)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("Error creating request for ID %d: %v", id, err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("request-id", reqID)

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error making request for ID %d: %v", id, err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("[request-id=\"%v\"] [id=%v]: failed with status %s", reqID, id, resp.Status)
		return
	}

	log.Printf("[request-id=\"%v\"] [id=%v]: succeeded with status %s", reqID, id, resp.Status)
}

func main() {
	var wg sync.WaitGroup

	wg.Add(100)

	for id := 1; id <= 100; id++ {
		go makeRequest(id, &wg)
	}

	wg.Wait()

	log.Println("All requests completed")
}
