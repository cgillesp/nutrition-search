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

  if (input == undefined) {
    return "—";
  }

  if (typeof input === "number") {
    return input.toFixed(0) + followString;
  }

  return input.toString() + followString;
}
