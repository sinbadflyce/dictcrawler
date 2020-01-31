package resolvers

import (
	"context"

	"github.com/sinbadflyce/dictcrawler/crawling"
	"github.com/sinbadflyce/dictcrawler/database"
	"github.com/sinbadflyce/dictcrawler/generates"
	"github.com/sinbadflyce/dictcrawler/models"
)

// LMResolver ...
type LMResolver struct{}

// CrawlerQuery ...
func (r *LMResolver) CrawlerQuery() generates.CrawlerQueryResolver {
	return &crawlerQueryLMResolver{r}
}

type crawlerQueryLMResolver struct{ *LMResolver }

// LookupWord ...
func (r *crawlerQueryLMResolver) LookupWord(ctx context.Context, name string) (*models.Word, error) {
	var w models.Word = database.DictRepo.Find(name)
	if len(w.Name) == 0 {
		var c crawling.Crawler
		c.AtURL = "https://www.ldoceonline.com/dictionary/" + name
		w = c.Run()
		if len(w.Name) > 0 {
			database.DictRepo.Save(w)
		}
	}

	return &w, nil
}
