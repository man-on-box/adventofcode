const input = await Deno.readTextFile("2022/05/input.txt");

const splitLines = input.split(`\n`);
const stackNumberLineIndex = splitLines.findIndex((line) =>
  line.startsWith(" 1")
);
const rawStacks = [...splitLines].slice(0, stackNumberLineIndex);
const rawCommands = [...splitLines].splice(stackNumberLineIndex + 2);

interface Command {
  fromStack: number;
  toStack: number;
  noOfCrates: number;
}
type StackMap = Map<number, string[]>;

const parseCommand = (str: string): Command => {
  const [noOfCrates, fromStack, toStack] = (str.match(/\d+/g) || []).map(
    Number
  );
  return {
    noOfCrates,
    fromStack,
    toStack,
  };
};
const commands = rawCommands.map(parseCommand);

const createStackMap = () => {
  const stackMap: StackMap = new Map();
  const stackIndexes: number[] = [];

  const rawStackStr = splitLines[stackNumberLineIndex];
  rawStackStr.split("").forEach((c, i) => {
    if (c === " ") return;
    stackMap.set(Number(c), []);
    stackIndexes.push(i);
  });

  const reversedRawStacks = [...rawStacks].reverse();

  reversedRawStacks.forEach((rawStack) => {
    stackIndexes.forEach((no, i) => {
      const stackNo = i + 1;
      const crate = rawStack[no];
      if (crate === " ") return;
      stackMap.set(stackNo, [...stackMap.get(stackNo)!, crate]);
    });
  });

  return stackMap;
};

class CraneOperator {
  stackMap: StackMap;
  constructor(stackMap: StackMap) {
    this.stackMap = stackMap;
    this.doCommand = this.doCommand.bind(this);
  }

  public doCommand({ fromStack, toStack, noOfCrates }: Command) {
    for (let i = 1; i <= noOfCrates; i++) {
      this.moveCrate(fromStack, toStack);
    }
  }

  public doCommandList(commands: Command[]) {
    commands.forEach(this.doCommand);
  }

  public getTopCrates() {
    const stacks = [...this.stackMap.values()];
    return stacks.map((stack) => stack.at(-1));
  }

  get stacks() {
    return this.stackMap;
  }

  private moveCrate(from: number, to: number) {
    const fromStack = this.stackMap.get(from)!;
    const toStack = this.stackMap.get(to)!;
    const crate = fromStack.pop()!;
    toStack.push(crate);
  }
}

class CraneOperator9001 extends CraneOperator {
  constructor(stackMap: StackMap) {
    super(stackMap);
    this.doCommand = this.doCommand.bind(this);
  }

  public doCommand(command: Command): void {
    this.moveCrates(command);
  }

  private moveCrates(command: Command) {
    const fromStack = this.stackMap.get(command.fromStack)!;
    const toStack = this.stackMap.get(command.toStack)!;
    const crates = fromStack.splice(fromStack.length - command.noOfCrates);
    this.stackMap.set(command.toStack, [...toStack, ...crates]);
  }
}

const craneOperator = new CraneOperator(createStackMap());
const craneOperator9001 = new CraneOperator9001(createStackMap());

// Part One
craneOperator.doCommandList(commands);
console.log(craneOperator.getTopCrates().join(""));

// Part Two
craneOperator9001.doCommandList(commands);
console.log(craneOperator9001.getTopCrates().join(""));
