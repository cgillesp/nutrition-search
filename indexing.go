package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/blevesearch/bleve"
)

type searchFields struct {
	Description string
	BrandOwner  string
}

func buildIndex(index bleve.Index) bleve.Index {
	countRow := dbc.QueryRowx("SELECT count(*) FROM foods_plus")
	var rowsCount int64
	err := countRow.Scan(&rowsCount)
	if err != nil {
		fmt.Println(err)
		panic("invalid count")
	}

	i := 0
	chunksize := 10000
	for true {
		rows, err := dbc.Queryx("SELECT * FROM foods_plus LIMIT $1 OFFSET $2",
			chunksize, i)
		if err != nil {
			fmt.Println(err)
			panic("Error reading foods")
		}

		j := 0
		batchOp := index.NewBatch()
		for rows.Next() {
			var food Food
			rows.StructScan(&food)
			batchOp.Index(strconv.FormatInt(food.FdcID, 10), getSearchFields(food))

			j++
		}

		err = index.Batch(batchOp)
		if err != nil {
			panic("index insertion failed")
		}

		printProgress(i, j, rowsCount)

		if j == 0 {
			break
		}
		i += chunksize
	}

	return index
}

func printProgress(i int, j int, totalRows int64) {
	fmt.Printf("\r %d of %d", i+j, totalRows)
}

func getSearchFields(food Food) searchFields {
	description := food.Description
	brandOwner := food.BrandOwner.String
	return searchFields{description, brandOwner}
}

func openIndex() bleve.Index {
	index, err := bleve.Open("./fts.bleve")
	if err != nil {
		index = initIndex()
	}
	return index
}

func initIndex() bleve.Index {
	err := os.RemoveAll("./fts.bleve")
	if err != nil {
		fmt.Println(err)
		panic("Failed to clear index")
	}
	mapping := bleve.NewIndexMapping()
	mapping.DefaultAnalyzer = "en"
	index, err := bleve.New("./fts.bleve", mapping)
	if err != nil {
		fmt.Println(err)
		panic("Could not initialize search database.")
	}
	buildIndex(index)
	return index
}
