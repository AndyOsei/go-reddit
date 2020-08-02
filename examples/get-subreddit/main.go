package main

import (
	"context"
	"fmt"
	"log"

	"github.com/vartanbeno/go-reddit"
)

var ctx = context.Background()

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() (err error) {
	withCredentials := reddit.WithCredentials("id", "secret", "username", "password")

	client, err := reddit.NewClient(nil, withCredentials)
	if err != nil {
		return
	}

	sr, _, err := client.Subreddit.Get(ctx, "golang")
	if err != nil {
		return
	}

	fmt.Printf("%s was created on %s and has %d subscribers.\n", sr.NamePrefixed, sr.Created.Local(), sr.Subscribers)

	return
}