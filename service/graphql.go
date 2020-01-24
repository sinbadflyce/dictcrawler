package service

import (
	"fmt"

	"github.com/graphql-go/graphql"
	"github.com/sinbadflyce/dictcrawler/crawling"
	"github.com/sinbadflyce/dictcrawler/database"
)

var exampleType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Example",
		Fields: graphql.Fields{
			"AudioURL": &graphql.Field{
				Type: graphql.String,
			},
			"Text": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var senseType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Sense",
		Fields: graphql.Fields{
			"SignPost": &graphql.Field{
				Type: graphql.String,
			},
			"Definition": &graphql.Field{
				Type: graphql.String,
			},
			"Gram": &graphql.Field{
				Type: graphql.String,
			},
			"Examples": &graphql.Field{
				Type: graphql.NewList(exampleType),
			},
		},
	},
)

var entryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Sense",
		Fields: graphql.Fields{
			"Topics": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
			"Homnum": &graphql.Field{
				Type: graphql.String,
			},
			"Freqs": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
			"SpeakerURLs": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
			"Hyphenation": &graphql.Field{
				Type: graphql.String,
			},
			"Pron": &graphql.Field{
				Type: graphql.String,
			},
			"Poses": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
			"Senses": &graphql.Field{
				Type: graphql.NewList(senseType),
			},
		},
	},
)

var wordType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Word",
		Fields: graphql.Fields{
			"Name": &graphql.Field{
				Type: graphql.String,
			},
			"Entries": &graphql.Field{
				Type: graphql.NewList(entryType),
			},
		},
	},
)

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"Word": &graphql.Field{
				Type: wordType,
				Args: graphql.FieldConfigArgument{
					"Name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					idQuery, isOK := p.Args["Name"].(string)
					if isOK {
						w := filterByWord(idQuery)
						return w, nil
					}
					return nil, nil
				},
			},
		},
	})

var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query: queryType,
	},
)

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}
	return result
}

func filterByWord(name string) crawling.Word {
	var w crawling.Word = database.DictRepo.Find(name)
	if len(w.Name) == 0 {
		var c crawling.Crawler
		c.AtURL = "https://www.ldoceonline.com/dictionary/" + name
		w = c.Run()
		if len(w.Name) > 0 {
			database.DictRepo.Save(w)
		}
	}
	return w
}
