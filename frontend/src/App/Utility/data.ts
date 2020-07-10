export interface Food {
  Index: number;
  FdcID: number;
  DataType: string;
  Description: string;
  FoodCategoryID?: number;
  PublicationDate?: string;
  DefaultSize?: number;
  DefaultUnit?: string;
  ServingSizeDescription?: string;
  BrandOwner?: string;
  GtinUPC?: string;
  DefaultCalories?: number;
  UnitCalories?: number;
}

export interface FoodNutrients {
  Description: string;
  Calories: number;
  Fat: number;
  SatFat: number;
  TransFat: number;
  Cholesterol: number;
  Sodium: number;
  Carbohydrates: number;
  Fiber: number;
  Sugars: number;
  AddedSugars: number;
  Protein: number;
}

export interface Result {
  GoodResponse: boolean;
  Hits?: Food[];
}

const apiDomain =
  process.env.NODE_ENV === "production"
    ? "https://demos.charliegillespie.com"
    : "http://localhost:4321";

export async function makeSearchRequest(
  query: string
): Promise<Food[] | undefined> {
  let queryURL = new URL(apiDomain + "/food/query");
  queryURL.searchParams.append("q", query);

  const response = fetch(queryURL.toString());

  const responseString = await (await response).text();

  const responseObject: Result = JSON.parse(responseString);

  return responseObject.Hits;
}

export async function makeDetailsRequest(
  fdcID: Number
): Promise<FoodNutrients | undefined> {
  let queryURL = new URL(apiDomain + "/food/nutrients");
  queryURL.searchParams.append("q", fdcID.toString());

  const response = fetch(queryURL.toString());

  const responseString = await (await response).text();

  const responseObject: FoodNutrients = JSON.parse(responseString);

  return responseObject;
}
