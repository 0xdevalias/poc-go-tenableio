package main

import (
	"fmt"
	"log"

	"github.com/0xdevalias/poc-go-tenableio/api"
)

func main() {
	c := api.DefaultClient(
		"TODO",
		"TODO",
	).WithDebug()

	TestAssets(c)
}

func TestAssets(c *api.Client) {
	assets, err := c.Workbenches.Assets()
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	for _, a := range assets.Assets {
		fmt.Printf("%#v\n\n", a)
	}
}
