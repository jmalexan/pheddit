package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/jzelinskie/geddit"
	"github.com/notdisliked/pheddit/perspective"
)

func main() {
	session := geddit.NewSession("gedditAgent v1")
	subOpts := geddit.ListingOptions{
		Limit: 100,
	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	text = text[:len(text)-1]
	submissions, _ := session.SubredditSubmissions(text, geddit.TopSubmissions, subOpts)
	var total float64
	for _, submission := range submissions {
		title := submission.Title
		fmt.Println(title)
		toxicity := perspective.GetToxicity(title)
		fmt.Println(toxicity * 100)
		total += toxicity
	}
	average := total / float64(len(submissions))
	fmt.Println()
	fmt.Println(average * 100)
}
