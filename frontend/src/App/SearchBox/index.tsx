import React, { useState } from "react";
import "./index.css";
import { makeSearchRequest, Food } from "../Utility/data";
import Results from "./Results";
import FoodDetails from "./FoodDetails";

export default function SearchBox() {
  const [query, setQuery] = useState("");
  const [resultsList, setResultsList] = useState<Food[]>();
  const [detailFood, setDetailFood] = useState<Food | undefined>();

  const onItemClick = (food: Food) => () => setDetailFood(food);
  const closeDetail = () => setDetailFood(undefined);
  return (
    <div className="searchArea">
      <input
        value={query}
        onChange={(e) => {
          const val = e.currentTarget.value;
          setDetailFood(undefined);
          setQuery(val);
          if (!val) {
            setResultsList([]);
          } else {
            makeSearchRequest(e.currentTarget.value).then((r) =>
              setResultsList(r ? r : resultsList)
            );
          }
        }}
        autoFocus
        placeholder="Search Foods..."
        className="searchBox"
      />
      {detailFood ? (
        <FoodDetails food={detailFood} closeDetail={closeDetail} />
      ) : (
        <Results resultsList={resultsList} onItemClick={onItemClick} />
      )}
    </div>
  );
}
