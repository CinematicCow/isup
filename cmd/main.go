package main

import (
	"fmt"

	"github.com/CinematicCow/isup/internal/checker"
)

func main() {
	websites := []string{"youtube", "google", "duckduckgo", "instagram", "github"}

	result := checker.WebsiteChecker(websites)

	for website, status := range result {
		fmt.Printf("%s is %s\n", website, status)
	}
}
