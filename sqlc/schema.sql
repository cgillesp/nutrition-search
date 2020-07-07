--
-- PostgreSQL database dump
--

-- Dumped from database version 12.3
-- Dumped by pg_dump version 12.3

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: branded_food; Type: TABLE; Schema: public; Owner: charlie
--

CREATE TABLE public.branded_food (
    index bigint,
    fdc_id bigint,
    brand_owner text,
    gtin_upc text,
    ingredients text,
    serving_size double precision,
    serving_size_unit text,
    household_serving_fulltext text,
    branded_food_category text,
    data_source text,
    modified_date text,
    available_date text,
    market_country text,
    discontinued_date double precision
);


ALTER TABLE public.branded_food OWNER TO charlie;

--
-- Name: food; Type: TABLE; Schema: public; Owner: charlie
--

CREATE TABLE public.food (
    index bigint,
    fdc_id bigint,
    data_type text,
    description text,
    food_category_id double precision,
    publication_date text
);


ALTER TABLE public.food OWNER TO charlie;

--
-- Name: food_attribute; Type: TABLE; Schema: public; Owner: charlie
--

CREATE TABLE public.food_attribute (
    index bigint,
    id bigint,
    fdc_id bigint,
    seq_num double precision,
    food_attribute_type_id bigint,
    name text,
    value text
);


ALTER TABLE public.food_attribute OWNER TO charlie;

--
-- Name: food_nutrient; Type: TABLE; Schema: public; Owner: charlie
--

CREATE TABLE public.food_nutrient (
    index bigint,
    id bigint,
    fdc_id bigint,
    nutrient_id bigint,
    amount double precision,
    data_points double precision,
    derivation_id double precision,
    min double precision,
    max double precision,
    median double precision,
    footnote text,
    min_year_acquired double precision
);


ALTER TABLE public.food_nutrient OWNER TO charlie;

--
-- Name: food_portion; Type: TABLE; Schema: public; Owner: charlie
--

CREATE TABLE public.food_portion (
    index bigint,
    id bigint,
    fdc_id bigint,
    seq_num double precision,
    amount double precision,
    measure_unit_id bigint,
    portion_description text,
    modifier text,
    gram_weight double precision,
    data_points double precision,
    footnote double precision,
    min_year_acquired double precision
);


ALTER TABLE public.food_portion OWNER TO charlie;

--
-- Name: measure_unit; Type: TABLE; Schema: public; Owner: charlie
--

CREATE TABLE public.measure_unit (
    index bigint,
    id bigint,
    name text
);


ALTER TABLE public.measure_unit OWNER TO charlie;

--
-- Name: portions_plus; Type: MATERIALIZED VIEW; Schema: public; Owner: charlie
--

CREATE MATERIALIZED VIEW public.portions_plus AS
 SELECT fp.index,
    fp.id,
    fp.fdc_id,
    fp.seq_num,
    fp.amount,
    fp.measure_unit_id,
    fp.portion_description,
    fp.modifier,
    fp.gram_weight,
    fp.data_points,
    fp.footnote,
    fp.min_year_acquired,
        CASE
            WHEN (mu.name = 'undetermined'::text) THEN NULL::text
            ELSE mu.name
        END AS unit_name
   FROM (public.food_portion fp
     LEFT JOIN public.measure_unit mu ON ((fp.measure_unit_id = mu.id)))
  WITH NO DATA;


ALTER TABLE public.portions_plus OWNER TO charlie;

--
-- Name: foods_plus; Type: MATERIALIZED VIEW; Schema: public; Owner: charlie
--

CREATE MATERIALIZED VIEW public.foods_plus AS
 SELECT fp.index,
    fp.fdc_id,
    fp.data_type,
    fp.description,
    fp.food_category_id,
    fp.publication_date,
    fp.default_size,
    fp.default_unit,
    fp.serving_size_description,
    fp.brand_owner,
    fp.gtin_upc,
    fp.default_calories,
    fp.unit_calories
   FROM ( SELECT food.index,
            food.fdc_id,
            food.data_type,
            food.description,
            food.food_category_id,
            food.publication_date,
            COALESCE(bf.serving_size, fp_1.gram_weight, (100)::double precision) AS default_size,
                CASE
                    WHEN (bf.serving_size IS NOT NULL) THEN bf.serving_size_unit
                    WHEN (fp_1.gram_weight IS NOT NULL) THEN 'g'::text
                    ELSE 'g'::text
                END AS default_unit,
                CASE
                    WHEN (bf.serving_size IS NOT NULL) THEN bf.household_serving_fulltext
                    WHEN (fp_1.gram_weight IS NOT NULL) THEN COALESCE(fp_1.portion_description, concat((fp_1.amount)::text, ' ', fp_1.unit_name))
                    ELSE NULL::text
                END AS serving_size_description,
            bf.brand_owner,
            bf.gtin_upc,
            (fn.amount * (COALESCE(bf.serving_size, fp_1.gram_weight, (100)::double precision) / (100)::double precision)) AS default_calories,
            fn.amount AS unit_calories
           FROM (((public.food
             LEFT JOIN public.branded_food bf ON ((food.fdc_id = bf.fdc_id)))
             LEFT JOIN ( SELECT DISTINCT ON (portions_plus.fdc_id) portions_plus.index,
                    portions_plus.id,
                    portions_plus.fdc_id,
                    portions_plus.seq_num,
                    portions_plus.amount,
                    portions_plus.measure_unit_id,
                    portions_plus.portion_description,
                    portions_plus.modifier,
                    portions_plus.gram_weight,
                    portions_plus.data_points,
                    portions_plus.footnote,
                    portions_plus.min_year_acquired,
                    portions_plus.unit_name
                   FROM public.portions_plus) fp_1 ON ((food.fdc_id = fp_1.fdc_id)))
             LEFT JOIN ( SELECT food_nutrient.index,
                    food_nutrient.id,
                    food_nutrient.fdc_id,
                    food_nutrient.nutrient_id,
                    food_nutrient.amount,
                    food_nutrient.data_points,
                    food_nutrient.derivation_id,
                    food_nutrient.min,
                    food_nutrient.max,
                    food_nutrient.median,
                    food_nutrient.footnote,
                    food_nutrient.min_year_acquired
                   FROM public.food_nutrient
                  WHERE (food_nutrient.nutrient_id = 1008)) fn ON ((food.fdc_id = fn.fdc_id)))) fp
  WHERE (fp.unit_calories IS NOT NULL)
  WITH NO DATA;


ALTER TABLE public.foods_plus OWNER TO charlie;

--
-- Name: input_food; Type: TABLE; Schema: public; Owner: charlie
--

CREATE TABLE public.input_food (
    index bigint,
    id bigint,
    fdc_id bigint,
    fdc_id_of_input_food double precision,
    seq_num double precision,
    amount double precision,
    sr_code double precision,
    sr_description text,
    unit text,
    portion_code double precision,
    portion_description text,
    gram_weight double precision,
    retention_code double precision,
    survey_flag double precision
);


ALTER TABLE public.input_food OWNER TO charlie;

--
-- Name: nutrient; Type: TABLE; Schema: public; Owner: charlie
--

CREATE TABLE public.nutrient (
    index bigint,
    id bigint,
    name text,
    unit_name text,
    nutrient_nbr double precision,
    rank double precision
);


ALTER TABLE public.nutrient OWNER TO charlie;

--
-- Name: food_nutrient_fdc_id_idx; Type: INDEX; Schema: public; Owner: charlie
--

CREATE INDEX food_nutrient_fdc_id_idx ON public.food_nutrient USING btree (fdc_id);


--
-- Name: food_nutrient_nutrient_id_idx; Type: INDEX; Schema: public; Owner: charlie
--

CREATE INDEX food_nutrient_nutrient_id_idx ON public.food_nutrient USING btree (nutrient_id);


--
-- Name: food_portion_fdc_id_idx; Type: INDEX; Schema: public; Owner: charlie
--

CREATE INDEX food_portion_fdc_id_idx ON public.food_portion USING btree (fdc_id);


--
-- Name: foods_plus_fdc_id_idx; Type: INDEX; Schema: public; Owner: charlie
--

CREATE UNIQUE INDEX foods_plus_fdc_id_idx ON public.foods_plus USING btree (fdc_id);


--
-- Name: ix_branded_food_index; Type: INDEX; Schema: public; Owner: charlie
--

CREATE INDEX ix_branded_food_index ON public.branded_food USING btree (index);


--
-- Name: ix_food_attribute_index; Type: INDEX; Schema: public; Owner: charlie
--

CREATE INDEX ix_food_attribute_index ON public.food_attribute USING btree (index);


--
-- Name: ix_food_index; Type: INDEX; Schema: public; Owner: charlie
--

CREATE INDEX ix_food_index ON public.food USING btree (index);


--
-- Name: ix_food_nutrient_index; Type: INDEX; Schema: public; Owner: charlie
--

CREATE INDEX ix_food_nutrient_index ON public.food_nutrient USING btree (index);


--
-- Name: ix_food_portion_index; Type: INDEX; Schema: public; Owner: charlie
--

CREATE INDEX ix_food_portion_index ON public.food_portion USING btree (index);


--
-- Name: ix_input_food_index; Type: INDEX; Schema: public; Owner: charlie
--

CREATE INDEX ix_input_food_index ON public.input_food USING btree (index);


--
-- Name: ix_measure_unit_index; Type: INDEX; Schema: public; Owner: charlie
--

CREATE INDEX ix_measure_unit_index ON public.measure_unit USING btree (index);


--
-- Name: ix_nutrient_index; Type: INDEX; Schema: public; Owner: charlie
--

CREATE INDEX ix_nutrient_index ON public.nutrient USING btree (index);


--
-- PostgreSQL database dump complete
--

