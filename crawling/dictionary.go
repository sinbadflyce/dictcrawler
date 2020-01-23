package crawling

// Example ...
type Example struct {
	AudioURL string
	Text     string
}

// Sense ...
type Sense struct {
	SignPost   string
	Definition string
	Gram       string
	Examples   []Example
}

// Entry ...
type Entry struct {
	Topics      []string
	Homnum      string
	Freqs       []string
	SpeakerURLs []string
	Hyphenation string
	Pron        string
	Poses       []string
	Senses      []Sense
}

// Word ...
type Word struct {
	Name    string
	Entries []Entry
}
