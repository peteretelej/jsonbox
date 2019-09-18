package main

import (
	"fmt"
	"log"

	"github.com/peteretelej/jsonbox"
)

const demoBox = "demobox_6d9e326c183fde7b"

func main() {

	cl, err := jsonbox.NewClient("https://jsonbox.io")
	if err != nil {
		log.Fatalf("failed to create jsonbox client: %v", err)
	}

	out, err := cl.Read("demobox_6d9e326c183fde7b?limit=1")
	if err != nil {
		log.Fatalf("failed to READ first record from %s: %v", demoBox, err)
	}

	fmt.Printf("READ demobox first record: %s\n", out)

}
