// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"database/sql"
)

type BrandedFood struct {
	Index                    sql.NullInt64   `json:"index"`
	FdcID                    sql.NullInt64   `json:"fdc_id"`
	BrandOwner               sql.NullString  `json:"brand_owner"`
	GtinUpc                  sql.NullString  `json:"gtin_upc"`
	Ingredients              sql.NullString  `json:"ingredients"`
	ServingSize              sql.NullFloat64 `json:"serving_size"`
	ServingSizeUnit          sql.NullString  `json:"serving_size_unit"`
	HouseholdServingFulltext sql.NullString  `json:"household_serving_fulltext"`
	BrandedFoodCategory      sql.NullString  `json:"branded_food_category"`
	DataSource               sql.NullString  `json:"data_source"`
	ModifiedDate             sql.NullString  `json:"modified_date"`
	AvailableDate            sql.NullString  `json:"available_date"`
	MarketCountry            sql.NullString  `json:"market_country"`
	DiscontinuedDate         sql.NullFloat64 `json:"discontinued_date"`
}

type Food struct {
	Index           sql.NullInt64   `json:"index"`
	FdcID           sql.NullInt64   `json:"fdc_id"`
	DataType        sql.NullString  `json:"data_type"`
	Description     sql.NullString  `json:"description"`
	FoodCategoryID  sql.NullFloat64 `json:"food_category_id"`
	PublicationDate sql.NullString  `json:"publication_date"`
}

type FoodAttribute struct {
	Index               sql.NullInt64   `json:"index"`
	ID                  sql.NullInt64   `json:"id"`
	FdcID               sql.NullInt64   `json:"fdc_id"`
	SeqNum              sql.NullFloat64 `json:"seq_num"`
	FoodAttributeTypeID sql.NullInt64   `json:"food_attribute_type_id"`
	Name                sql.NullString  `json:"name"`
	Value               sql.NullString  `json:"value"`
}

type FoodNutrient struct {
	Index           sql.NullInt64   `json:"index"`
	ID              sql.NullInt64   `json:"id"`
	FdcID           sql.NullInt64   `json:"fdc_id"`
	NutrientID      sql.NullInt64   `json:"nutrient_id"`
	Amount          sql.NullFloat64 `json:"amount"`
	DataPoints      sql.NullFloat64 `json:"data_points"`
	DerivationID    sql.NullFloat64 `json:"derivation_id"`
	Min             sql.NullFloat64 `json:"min"`
	Max             sql.NullFloat64 `json:"max"`
	Median          sql.NullFloat64 `json:"median"`
	Footnote        sql.NullString  `json:"footnote"`
	MinYearAcquired sql.NullFloat64 `json:"min_year_acquired"`
}

type FoodPortion struct {
	Index              sql.NullInt64   `json:"index"`
	ID                 sql.NullInt64   `json:"id"`
	FdcID              sql.NullInt64   `json:"fdc_id"`
	SeqNum             sql.NullFloat64 `json:"seq_num"`
	Amount             sql.NullFloat64 `json:"amount"`
	MeasureUnitID      sql.NullInt64   `json:"measure_unit_id"`
	PortionDescription sql.NullString  `json:"portion_description"`
	Modifier           sql.NullString  `json:"modifier"`
	GramWeight         sql.NullFloat64 `json:"gram_weight"`
	DataPoints         sql.NullFloat64 `json:"data_points"`
	Footnote           sql.NullFloat64 `json:"footnote"`
	MinYearAcquired    sql.NullFloat64 `json:"min_year_acquired"`
}

type InputFood struct {
	Index              sql.NullInt64   `json:"index"`
	ID                 sql.NullInt64   `json:"id"`
	FdcID              sql.NullInt64   `json:"fdc_id"`
	FdcIDOfInputFood   sql.NullFloat64 `json:"fdc_id_of_input_food"`
	SeqNum             sql.NullFloat64 `json:"seq_num"`
	Amount             sql.NullFloat64 `json:"amount"`
	SrCode             sql.NullFloat64 `json:"sr_code"`
	SrDescription      sql.NullString  `json:"sr_description"`
	Unit               sql.NullString  `json:"unit"`
	PortionCode        sql.NullFloat64 `json:"portion_code"`
	PortionDescription sql.NullString  `json:"portion_description"`
	GramWeight         sql.NullFloat64 `json:"gram_weight"`
	RetentionCode      sql.NullFloat64 `json:"retention_code"`
	SurveyFlag         sql.NullFloat64 `json:"survey_flag"`
}

type MeasureUnit struct {
	Index sql.NullInt64  `json:"index"`
	ID    sql.NullInt64  `json:"id"`
	Name  sql.NullString `json:"name"`
}

type Nutrient struct {
	Index       sql.NullInt64   `json:"index"`
	ID          sql.NullInt64   `json:"id"`
	Name        sql.NullString  `json:"name"`
	UnitName    sql.NullString  `json:"unit_name"`
	NutrientNbr sql.NullFloat64 `json:"nutrient_nbr"`
	Rank        sql.NullFloat64 `json:"rank"`
}
