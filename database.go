package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"strconv"

	"nutrition/db"

	"github.com/blevesearch/bleve"
	"gopkg.in/guregu/null.v4"
)

// Food - Stores the profile of a food retrieved from the DB.
type Food struct {
	Index                  int64       `db:"index"`
	FdcID                  int64       `db:"fdc_id"`
	DataType               string      `db:"data_type"`
	Description            string      `db:"description"`
	FoodCategoryID         null.Float  `db:"food_category_id"`
	PublicationDate        null.String `db:"publication_date"`
	DefaultSize            null.Float  `db:"default_size"`
	DefaultUnit            null.String `db:"default_unit"`
	ServingSizeDescription null.String `db:"serving_size_description"`
	BrandOwner             null.String `db:"brand_owner"`
	GtinUPC                null.String `db:"gtin_upc"`
	DefaultCalories        null.Float  `db:"default_calories"`
	UnitCalories           null.Float  `db:"unit_calories"`
}

func search(index bleve.Index, query string) (*bleve.SearchResult, error) {
	bleveQuery := bleve.NewMatchQuery(query)
	search := bleve.NewSearchRequest(bleveQuery)
	searchResults, err := index.Search(search)
	if err != nil {
		return nil, errors.New("Error searching")
	}
	return searchResults, nil
}

func getSearchResultsJSON(index bleve.Index, query string) (string, error) {
	results, err := search(index, query)
	if err != nil {
		return "", err
	}
	topTen := results.Hits[0:10]

	var foodResults []Food
	for i := 0; i < topTen.Len(); i++ {
		if topTen[i] == nil {
			break
		}
		id := topTen[i].ID
		row := dbc.QueryRowx("SELECT * FROM foods_plus WHERE fdc_id = $1", id)
		var food Food
		err := row.StructScan(&food)
		if err != nil {
			return "", errors.New("error reading results from database")
		}
		foodResults = append(foodResults, food)
	}

	response := struct {
		GoodResponse bool
		Hits         []Food
	}{
		true,
		foodResults,
	}

	jsonResult, err := json.Marshal(response)
	if err != nil {
		return "", errors.New("Could not encode JSON")
	}

	return string(jsonResult), nil
}

func getNutrientsJSON(fdcID string) (string, error) {
	q := db.New(dbc)
	fdcInt, err := strconv.ParseInt(fdcID, 10, 64)
	if err != nil {
		return "", errors.New("fdc_id must be an integer")
	}
	nullableFDC := sql.NullInt64{
		Int64: fdcInt,
		Valid: true,
	}
	rowResult, err := q.GetNutrients(context.Background(), nullableFDC)
	if err != nil {
		return "", errors.New("Could not find food with that ID")
	}

	jsonResult, err := json.Marshal(rowResult)

	if err != nil {
		return "", errors.New("Could not encode JSON")
	}

	return string(jsonResult), nil
}
