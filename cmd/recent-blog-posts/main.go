package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/mmcdole/gofeed"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %v <XML/Atom URL>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("# Most recent blog posts :pencil:")
	for _, i := range feed.Items {
		fmt.Printf("* [%v](%v)\n", i.Title, i.Link)
	}
}
