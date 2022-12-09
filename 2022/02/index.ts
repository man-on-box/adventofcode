const input = await Deno.readTextFile("2022/02/input.txt");

const OPP_ROCK = "A";
const OPP_PAPER = "B";
const OPP_SCISSOR = "C";
const YOU_ROCK = "X";
const YOU_PAPER = "Y";
const YOU_SCISSOR = "Z";

// const TO_LOSE = "X";
const TO_DRAW = "Y";
const TO_WIN = "Z";

const scoreMap = {
  X: 1,
  Y: 2,
  Z: 3,
  lose: 0,
  draw: 3,
  win: 6,
};

type Result = "win" | "draw" | "lose";
type Opponent = "A" | "B" | "C";
type You = "X" | "Y" | "Z";
type Round = [Opponent, You];

const rounds = input.split(`\n`).map((str) => str.split(" ") as Round);

const getResult = ([opponent, you]: Round): Result => {
  switch (opponent) {
    case OPP_ROCK:
      return you === YOU_ROCK ? "draw" : you === YOU_PAPER ? "win" : "lose";
    case OPP_PAPER:
      return you === YOU_PAPER ? "draw" : you === YOU_SCISSOR ? "win" : "lose";
    case OPP_SCISSOR:
      return you === YOU_SCISSOR ? "draw" : you === YOU_ROCK ? "win" : "lose";
    default:
      throw new Error(`Unknown opponent: ${opponent}`);
  }
};

const getHands = ([opponent, need]: Round): Round => {
  let you: You = "Z";
  switch (opponent) {
    case OPP_ROCK:
      need === TO_WIN
        ? (you = YOU_PAPER)
        : need === TO_DRAW
        ? (you = YOU_ROCK)
        : (you = YOU_SCISSOR);
      break;
    case OPP_PAPER:
      need === TO_WIN
        ? (you = YOU_SCISSOR)
        : need === TO_DRAW
        ? (you = YOU_PAPER)
        : (you = YOU_ROCK);
      break;
    case OPP_SCISSOR:
      need === TO_WIN
        ? (you = YOU_ROCK)
        : need === TO_DRAW
        ? (you = YOU_SCISSOR)
        : (you = YOU_PAPER);
      break;
  }
  return [opponent, you];
};

const calculateScoreForRound = (round: Round): number => {
  const hands = getHands(round);
  const result = getResult(hands);
  const score = scoreMap[result] + scoreMap[hands[1]];

  return score;
};

const scores = rounds.map(calculateScoreForRound);
const totalScore = scores.reduce((total, score) => total + score, 0);

console.log(totalScore);
