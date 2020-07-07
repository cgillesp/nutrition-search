import React from "react";
import { Food } from "../../Utility/data";
import { titleCase, valOrDash } from "../../Utility/text";

type resultProps = {
  resultsList: Food[] | undefined;
  onItemClick: (food: Food) => () => void;
};
export default function Results(props: resultProps) {
  return (
    <div className="results">
      {props.resultsList?.map((f, i) => (
        <div className="resultItem" key={i} onClick={props.onItemClick(f)}>
          <div className="leftStack">
            <div className="foodName">{titleCase(f.Description)}</div>
            <div className="brandName">
              {f.BrandOwner ? titleCase(f.BrandOwner) : ""}
            </div>
          </div>
          <div className="rightStack">
            <div className="caloriesNumber">{valOrDash(f.UnitCalories)}</div>
            <div className="caloriesLabel">cal/100{f.DefaultUnit}</div>
          </div>
        </div>
      ))}
    </div>
  );
}
