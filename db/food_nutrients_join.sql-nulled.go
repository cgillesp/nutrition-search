package db

import (
	"context"
	"database/sql"

	"gopkg.in/guregu/null.v4"
)

const getNutrients = `-- name: GetNutrients :one
SELECT description, cal.amount as calories,
                       fat.amount as fat, sfat.amount as sat_fat,
                       tfat.amount as trans_fat, clst.amount as cholesterol,
                       sodm.amount as sodium, carb.amount as carbohydrates,
                       fibr.amount as fiber, sugr.amount as sugars,
                       asgr.amount as added_sugars, prtn.amount as protein
                FROM
                    food fp
                        LEFT JOIN food_nutrient cal
                                  ON cal.nutrient_id = 1008
                                      AND cal.fdc_id = fp.fdc_id
                        LEFT JOIN food_nutrient fat
                                  ON fat.nutrient_id = 1004
                                      AND fp.fdc_id = fat.fdc_id
                        LEFT JOIN food_nutrient sfat
                                  ON sfat.nutrient_id = 1258
                                      AND fp.fdc_id = sfat.fdc_id
                        LEFT JOIN food_nutrient tfat
                                  ON tfat.nutrient_id = 1257
                                      AND fp.fdc_id = tfat.fdc_id
                        LEFT JOIN food_nutrient clst
                                  ON clst.nutrient_id = 1253
                                      AND fp.fdc_id = clst.fdc_id
                        LEFT JOIN food_nutrient sodm
                                  ON sodm.nutrient_id = 1093
                                      AND fp.fdc_id = sodm.fdc_id
                        LEFT JOIN food_nutrient carb
                                  ON carb.nutrient_id = 1005
                                      AND fp.fdc_id = carb.fdc_id
                        LEFT JOIN food_nutrient fibr
                                  ON fibr.nutrient_id = 1079
                                      AND fp.fdc_id = fibr.fdc_id
                        LEFT JOIN food_nutrient sugr
                                  ON sugr.nutrient_id = 2000
                                      AND fp.fdc_id = sugr.fdc_id
                        LEFT JOIN food_nutrient asgr
                                  ON asgr.nutrient_id = 1235
                                      AND fp.fdc_id = asgr.fdc_id
                        LEFT JOIN food_nutrient prtn
                                  ON prtn.nutrient_id = 1003
                                      AND fp.fdc_id = prtn.fdc_id
                   WHERE fp.fdc_id = $1
`

// GetNutrientsRow - Result of Get Nutrient Query
type GetNutrientsRow struct {
	Description   null.String
	Calories      null.Float
	Fat           null.Float
	SatFat        null.Float
	TransFat      null.Float
	Cholesterol   null.Float
	Sodium        null.Float
	Carbohydrates null.Float
	Fiber         null.Float
	Sugars        null.Float
	AddedSugars   null.Float
	Protein       null.Float
}

// GetNutrients - Query to Get Nutrient Row
func (q *Queries) GetNutrients(ctx context.Context, fdcID sql.NullInt64) (GetNutrientsRow, error) {
	row := q.queryRow(ctx, q.getNutrientsStmt, getNutrients, fdcID)
	var i GetNutrientsRow
	err := row.Scan(
		&i.Description,
		&i.Calories,
		&i.Fat,
		&i.SatFat,
		&i.TransFat,
		&i.Cholesterol,
		&i.Sodium,
		&i.Carbohydrates,
		&i.Fiber,
		&i.Sugars,
		&i.AddedSugars,
		&i.Protein,
	)
	return i, err
}
