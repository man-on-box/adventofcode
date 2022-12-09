const input = await Deno.readTextFile("2022/04/input.txt");

const assignmentPairs = input
  .split(`\n`)
  .map((s) =>
    s.split(",").map((s) => s.split("-").map(Number))
  ) as AssignmentPair[];

type Assignment = [number, number];
type AssignmentPair = [Assignment, Assignment];

const pairContainsTheOther = (pair: AssignmentPair) => {
  const minOverlaps = pair[0][0] >= pair[1][0];
  const maxOverlaps = pair[0][1] <= pair[1][1];
  return minOverlaps && maxOverlaps;
};

const overlappingPair = (pair: AssignmentPair) => {
  return (
    pairContainsTheOther(pair) ||
    pairContainsTheOther([...pair].reverse() as AssignmentPair)
  );
};

const pairOverlaps = ([one, two]: AssignmentPair) => {
  const overlaps = one[0] <= two[1] && two[0] <= one[1];
  return overlaps;
};

const containingPairs = assignmentPairs.filter(overlappingPair);
const overlappingPairs = assignmentPairs.filter(pairOverlaps);

console.log(containingPairs.length, overlappingPairs.length);
