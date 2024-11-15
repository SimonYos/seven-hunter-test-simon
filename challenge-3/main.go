package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"
)

var meatType = []string{
	"t-bone", "fatback", "pastrami", "pork", "meatloaf", "jowl", "enim", "bresaola",
}

// BeefSum for response
type BeefSum struct {
	Beef map[string]int `json:"beef"`
}

// FetchTextData from url
func FetchTextData() (string, error) {
	resp, err := http.Get("https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text")
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	boff, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(boff), nil
}

// CountMeats from text
func CountMeats(text string) BeefSum {
	meats := make(map[string]int)

	re := regexp.MustCompile(`[,.]`)

	textSplit := re.ReplaceAllString(text, "")
	words := strings.Fields(strings.ToLower(textSplit))

	for _, word := range words {
		for _, meat := range meatType {
			if word == meat {
				meats[meat]++
			}
		}
	}

	return BeefSum{Beef: meats}
}

// Handler for beef summary
func Handler(w http.ResponseWriter, r *http.Request) {
	text, err := FetchTextData()
	if err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	summary := CountMeats(text)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(summary)
}

func main() {
	http.HandleFunc("/beef/summary", Handler)
	fmt.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
