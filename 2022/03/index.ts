const input = await Deno.readTextFile("2022/03/input.txt");
const rucksacks = input.split(`\n`);

const alphLower = "abcdefghijklmnopqrstuvwxyz".split("");
const alphUpper = alphLower.map((c) => c.toUpperCase());

const getCharacterScore = (char: string) => {
  if (alphLower.includes(char)) {
    return alphLower.indexOf(char) + 1;
  }
  if (alphUpper.includes(char)) {
    return alphUpper.indexOf(char) + 27;
  }
  throw new Error(`could not get score of: ${char}`);
};

const splitStringInHalf = (str: string) => {
  const divider = str.length / 2;
  const firstHalf = str.substring(0, divider);
  const secondHalf = str.substring(divider, divider * 2);
  const splitString = [firstHalf, secondHalf];
  return splitString;
};

const getCommonCharacters = (arrOfStrings: string[]) => {
  const splitStrings = arrOfStrings.map((str) => str.split(""));
  const [firstString, ...rest] = splitStrings;

  const commonChars = firstString.filter((char) => {
    return rest.every((arr) => arr.includes(char));
  });

  return commonChars;
};

const parseRucksack = (rucksack: string, index: number) => {
  const compartments = splitStringInHalf(rucksack);
  const commonChars = getCommonCharacters(compartments);
  const groupNumber = Math.ceil((index + 1) / 3);
  return {
    rucksack,
    compartments,
    commonChars,
    groupNumber,
    score: getCharacterScore(commonChars[0]),
  };
};

type ParsedRucksack = ReturnType<typeof parseRucksack>;

const parsedRucksacks = rucksacks.map(parseRucksack);

const groupParsedRucksacks = (parsedRucksacks: ParsedRucksack[]) => {
  const groupedRucksacks = new Map<number, ParsedRucksack[]>();

  parsedRucksacks.forEach((rucksack) => {
    const { groupNumber } = rucksack;
    const rucksackGroup = groupedRucksacks.get(groupNumber);
    if (rucksackGroup) {
      groupedRucksacks.set(groupNumber, [...rucksackGroup, rucksack]);
    } else {
      groupedRucksacks.set(groupNumber, [rucksack]);
    }
  });
  const groups = [...groupedRucksacks.values()];

  return groups.map(parseRucksackGroup);
};

const parseRucksackGroup = (rucksackGroup: ParsedRucksack[]) => {
  const rucksacks = rucksackGroup.map(({ rucksack }) => rucksack);
  const badge = getCommonCharacters(rucksacks)[0];
  const score = getCharacterScore(badge);

  return {
    rucksackGroup,
    badge,
    score,
  };
};

const groupedRucksacks = groupParsedRucksacks(parsedRucksacks);

const totalScores = parsedRucksacks.reduce((acc, { score }) => {
  return acc + score;
}, 0);

const totalBadgeScores = groupedRucksacks.reduce((acc, { score }) => {
  return acc + score;
}, 0);

console.log({ totalScores, totalBadgeScores });
