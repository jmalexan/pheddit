package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jzelinskie/geddit"
	"github.com/notdisliked/pheddit/perspective"
)

func main() {
	session := geddit.NewSession("gedditAgent v1")
	subOpts := geddit.ListingOptions{
		Limit: 100,
		Time:  geddit.ThisMonth,
	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Select subreddit for scanning: /r/")
	text, _ := reader.ReadString('\n')
	text = strings.Trim(text, "\r\n")
	submissions, err := session.SubredditSubmissions(text, geddit.TopSubmissions, subOpts)
	if err != nil {
		panic(err)
	}
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
