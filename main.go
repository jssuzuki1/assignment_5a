package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gocolly/colly/v2"
)

// Data structure to hold URL and scraped text
type pageText struct {
	URL  string `json:"url"`
	Text string `json:"text"`
}

func main() {

	// Slice to store the scraped data
	var scrapedText []pageText

	// urls is an array. We will loop through this array to scrape.
	urls := []string{
		"https://en.wikipedia.org/wiki/Robotics",
		"https://en.wikipedia.org/wiki/Robot",
		"https://en.wikipedia.org/wiki/Reinforcement_learning",
		"https://en.wikipedia.org/wiki/Robot_Operating_System",
		"https://en.wikipedia.org/wiki/Intelligent_agent",
		"https://en.wikipedia.org/wiki/Software_agent",
		"https://en.wikipedia.org/wiki/Robotic_process_automation",
		"https://en.wikipedia.org/wiki/Chatbot",
		"https://en.wikipedia.org/wiki/Applications_of_artificial_intelligence",
		"https://en.wikipedia.org/wiki/Android_(robot)",
	}

	// Declare new collector
	c := colly.NewCollector()

	// This OnHTML callback will extract the text, append the results to the scrapedText slice.
	c.OnHTML(".mw-parser-output > p", func(e *colly.HTMLElement) {
		// Concatenate the text from all paragraphs into a single string
		scrapedText[len(scrapedText)-1].Text += e.Text + "\n"
	})

	// Error handling in case things go wrong
	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("failed with response:", err)
	})

	for _, url := range urls {
		// Append a new PageData with an empty text field for each URL
		scrapedText = append(scrapedText, pageText{URL: url, Text: ""})
		err := c.Visit(url)
		if err != nil {
			fmt.Println("Error visiting the URL:", err)
		}
	}

	// Declare the creation of a new JSON file.
	// This JSON file will be created once the function finishes running
	file, err := os.Create("scraped_wikis.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Encode the scraped data into JSON and write it to the file
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")
	err = encoder.Encode(scrapedText)
	if err != nil {
		fmt.Println("Error encoding data to JSON:", err)
		return
	}

}
