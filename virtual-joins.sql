DROP MATERIALIZED VIEW IF EXISTS portions_plus CASCADE;

CREATE MATERIALIZED VIEW portions_plus AS
    SELECT fp.*,
           CASE
               WHEN mu.name = 'undetermined' THEN null
               ELSE mu.name
               END
               as unit_name
        FROM food_portion fp
            LEFT OUTER JOIN measure_unit mu
            on fp.measure_unit_id = mu.id;


DROP MATERIALIZED VIEW IF EXISTS foods_plus CASCADE;

CREATE MATERIALIZED VIEW foods_plus AS
    SELECT * FROM (SELECT food.*,
           coalesce(bf.serving_size, fp.gram_weight, 100) as default_size,
           (CASE
               WHEN bf.serving_size IS NOT NULL THEN bf.serving_size_unit
               WHEN fp.gram_weight IS NOT NULL THEN 'g'
               ELSE 'g'
                END)
               as default_unit,
           (CASE
               WHEN bf.serving_size IS NOT NULL THEN bf.household_serving_fulltext
               WHEN fp.gram_weight IS NOT NULL THEN coalesce(fp.portion_description,
                   concat(fp.amount::text, ' ', fp.unit_name))
               END)
               as serving_size_description,
           brand_owner,
           gtin_upc,
           fn.amount * (coalesce(bf.serving_size, fp.gram_weight, 100)/100) as default_calories,
           fn.amount as unit_calories
    FROM food
    LEFT OUTER JOIN branded_food bf
        on food.fdc_id = bf.fdc_id
    LEFT OUTER JOIN (SELECT DISTINCT on (fdc_id) * FROM portions_plus) fp
        on food.fdc_id = fp.fdc_id
    LEFT JOIN (SELECT * FROM food_nutrient WHERE nutrient_id = 1008) fn
    on food.fdc_id = fn.fdc_id) as fp WHERE unit_calories IS NOT NULL ;

DROP INDEX IF EXISTS foods_plus_fdc_id_idx;
CREATE UNIQUE INDEX on foods_plus (fdc_id);

DROP INDEX IF EXISTS food_nutrient_fdc_id_idx;
CREATE INDEX on food_nutrient (fdc_id);

DROP INDEX IF EXISTS food_nutrient_nutrient_id_idx;
CREATE INDEX on food_nutrient (nutrient_id);

DROP INDEX IF EXISTS food_portion_fdc_id_idx;
CREATE INDEX on food_portion (fdc_id);

UPDATE branded_food SET brand_owner = regexp_replace(brand_owner, '\s*-*\s*\(?\d{5,}\)?', '') WHERE brand_owner ~ '\d{5,}'



