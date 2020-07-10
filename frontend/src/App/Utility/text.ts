import { titleCase as tc } from "title-case";

export function titleCase(input: String | undefined) {
  if (!input) {
    return undefined;
  }
  return tc(input.toLowerCase());
}

export function valOrDash(
  input: String | Number | undefined,
  follow?: String | undefined
): string {
  const followString = follow ? follow : "";

  // Okay so eslint complains about not using isNaN
  // But that would require casting to a number :(
  // eslint-disable-next-line
  if (input == null || input === "" || input === NaN) {
    return "—";
  }

  if (typeof input === "number") {
    return input.toFixed(0) + followString;
  }

  return input.toString() + followString;
}
