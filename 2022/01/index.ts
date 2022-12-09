const input = await Deno.readTextFile("2022/01/input.txt");

const caloriesByElves = input
  .replaceAll(`\n`, ".")
  .split("..")
  .map((string) => string.split("."));

const parseElves = (elves: string[][]) => {
  return elves.map((caloriesStrings) => {
    const calories = caloriesStrings.map(Number);
    const totalCalories = calories.reduce((acc, cal) => cal + acc, 0);

    return {
      calories,
      total: totalCalories,
    };
  });
};

const parsedElves = parseElves(caloriesByElves);
const arrOfTotals = parsedElves.map(({ total }) => total)

const sortedTotals = arrOfTotals.sort((a, b) => b-a)
const top3elves = sortedTotals.slice(0, 3)

console.log(top3elves.reduce((acc, cal) => cal + acc, 0));
