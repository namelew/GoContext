package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "https://localhost:8080", nil)
	if err != nil {
		log.Panic(err.Error())
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Panic(err.Error())
	}

	defer res.Body.Close()

	io.Copy(os.Stdout, res.Body)
}
