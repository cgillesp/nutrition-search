import React, { useEffect, useState } from "react";
import { FoodNutrients, makeDetailsRequest, Food } from "../../Utility/data";
import "./index.css";
import { titleCase, valOrDash } from "../../Utility/text";

type detailProps = {
  food: Food;
};

export default function FoodDetails(props: detailProps) {
  const [foodDetail, setFoodDetail] = useState<FoodNutrients>();

  useEffect(() => {
    makeDetailsRequest(props.food.FdcID).then((r) => setFoodDetail(r));
  }, [props.food.FdcID]);
  return (
    <div className="results detailView">
      <div className="foodDetailName">{titleCase(props.food.Description)}</div>
      <div className="foodDetailBrandName">
        {valOrDash(titleCase(props.food.BrandOwner))}
      </div>
      <div>Serving Size: 100g (3.5oz)</div>
      <hr />
      <div className="nutritionRow caloriesRow">
        <b>Calories</b> {valOrDash(foodDetail?.Calories)}
      </div>
      <div className="nutritionRow subNutritionRow">
        <b>Total Fat</b> {valOrDash(foodDetail?.Fat, "g")}
      </div>
      <div className="nutritionRow subNutritionRow">
        <b>Saturated Fat</b> {valOrDash(foodDetail?.SatFat, "g")}
      </div>
      <div className="nutritionRow subNutritionRow">
        <b>Trans Fat</b> {valOrDash(foodDetail?.TransFat, "g")}
      </div>
      <div className="nutritionRow">
        <b>Cholesterol</b> {valOrDash(foodDetail?.Cholesterol, "mg")}
      </div>
      <div className="nutritionRow">
        <b>Sodium</b> {valOrDash(foodDetail?.Sodium, "mg")}
      </div>
      <div className="nutritionRow">
        <b>Total Carbohydrate</b> {valOrDash(foodDetail?.Carbohydrates, "g")}
      </div>
      <div className="nutritionRow subNutritionRow">
        <b>Dietary Fiber</b> {valOrDash(foodDetail?.Fiber, "g")}
      </div>
      <div className="nutritionRow subNutritionRow">
        <b>Total Sugars</b> {valOrDash(foodDetail?.Sugars, "g")}
      </div>
      <div className="nutritionRow subNutritionRow">
        <b>Added Sugars</b> {valOrDash(foodDetail?.AddedSugars, "g")}
      </div>
      <div className="nutritionRow">
        <b>Protein</b> {valOrDash(foodDetail?.Protein, "g")}
      </div>
    </div>
  );
}
