package crawling

import (
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
)

// Crawler ...
type Crawler struct {
	AtURL string
}

// Run ...
func (crawler *Crawler) Run() Word {
	if len(crawler.AtURL) == 0 {
		return Word{}
	}

	c := colly.NewCollector()
	word := Word{}
	entries := make([]Entry, 0, 100)

	// Word Name
	c.OnHTML("h1[class=pagetitle]", func(e *colly.HTMLElement) {
		word.Name = e.Text
	})

	// Dictionary Entry
	c.OnHTML("span[class=dictentry]", func(e *colly.HTMLElement) {
		entry := Entry{}

		e.DOM.Find(".topics_container").Each(func(i int, s *goquery.Selection) {
			entry.Topics = parseTOPICs(s)
		})

		e.DOM.Find(".Head").Each(func(i int, s *goquery.Selection) {
			entry.Homnum = parseHOMNUM(s)
			entry.Hyphenation = parseHYPHENATION(s)
			entry.Freqs = parseFREQ(s)
			entry.Poses = parsePOS(s)
			entry.SpeakerURLs = parseSpeakerURLs(s)
			entry.Pron = parsePRON(s)
		})

		e.DOM.Find(".Sense").Each(func(i int, s *goquery.Selection) {
			sense := Sense{}
			sense.SignPost = parseSIGNPOST(s)
			sense.Definition = parseDEF(s)
			sense.Gram = parseGRAM(s)
			sense.Examples = parseEXAMPLES(s)
			entry.Senses = append(entry.Senses, sense)
		})

		entries = append(entries, entry)
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Println("Error:", r.StatusCode, err)
	})

	c.Visit(crawler.AtURL)
	word.Entries = entries
	return word
}

// Parsing Topic

func parseTOPICs(topicSelection *goquery.Selection) []string {
	var result []string
	topicSelection.Find(".topic").Each(func(i int, s *goquery.Selection) {
		topic, _ := s.Attr("title")
		result = append(result, topic)
	})
	return result
}

// Parsing of Sense Selection

func parseGRAM(senseSelection *goquery.Selection) string {
	var result string
	senseSelection.Find(".GRAM").Each(func(i int, s *goquery.Selection) {
		result = strings.TrimSpace(s.Text())
	})
	return result
}

func parseEXAMPLES(senseSelection *goquery.Selection) []Example {
	var result []Example
	senseSelection.Find(".EXAMPLE").Each(func(i int, s *goquery.Selection) {
		example := Example{}
		example.Text = strings.TrimSpace(s.Text())
		s.Find(".speaker").Each(func(i int, s1 *goquery.Selection) {
			example.AudioURL, _ = s1.Attr("data-src-mp3")
		})
		result = append(result, example)
	})
	return result
}

func parseDEF(senseSelection *goquery.Selection) string {
	var result string
	senseSelection.Find(".DEF").Each(func(i int, s *goquery.Selection) {
		result = s.Text()
	})

	if len(result) == 0 {
		senseSelection.Find(".REFHWD").Each(func(i int, s *goquery.Selection) {
			result = s.Text()
		})
	}
	result = strings.TrimSpace(result)
	return result
}

func parseSIGNPOST(senseSelection *goquery.Selection) string {
	var result string
	senseSelection.Find(".SIGNPOST").Each(func(i int, s *goquery.Selection) {
		result = s.Text()
	})
	return result
}

// Parsing of frequent Selection

func parseHOMNUM(frequentSelection *goquery.Selection) string {
	var result string
	frequentSelection.Find(".HOMNUM").Each(func(i int, s *goquery.Selection) {
		result = s.Text()
	})
	return result
}

func parseHYPHENATION(frequentSelection *goquery.Selection) string {
	var result string
	frequentSelection.Find(".HYPHENATION").Each(func(i int, s *goquery.Selection) {
		result = s.Text()
	})
	return result
}

func parsePRON(frequentSelection *goquery.Selection) string {
	var result string
	frequentSelection.Find(".PRON").Each(func(i int, s *goquery.Selection) {
		result = s.Text()
	})
	return result
}

func parseFREQ(frequentSelection *goquery.Selection) []string {
	var result []string
	frequentSelection.Find(".FREQ").Each(func(i int, s *goquery.Selection) {
		result = append(result, s.Text())
	})
	return result
}

func parsePOS(frequentSelection *goquery.Selection) []string {
	var result []string
	frequentSelection.Find(".POS").Each(func(i int, s *goquery.Selection) {
		result = append(result, strings.TrimSpace(s.Text()))
	})
	return result
}

func parseSpeakerURLs(frequentSelection *goquery.Selection) []string {
	var result []string
	frequentSelection.Find(".speaker").Each(func(i int, s *goquery.Selection) {
		surl, _ := s.Attr("data-src-mp3")
		result = append(result, surl)
	})
	return result
}
