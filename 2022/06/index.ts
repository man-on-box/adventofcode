const input = await Deno.readTextFile("2022/06/input.txt");

const SIGNAL_MARKER = 4;
const MESSAGE_MARKER = 14;

class SignalFinder {
  private buffer: string[];
  private marker: number;
  private signalIndexes: number[] = [];
  private foundSignals: string[] = [];

  constructor(buffer: string, marker: number) {
    this.buffer = buffer.split("");
    this.marker = marker;
    this.processBuffer()
  }

  get signals() {
    return {
      indexes: this.signalIndexes,
      signals: this.foundSignals,
    };
  }

  private processBuffer() {
    this.buffer.forEach((_c, i) => {
      if (i < this.marker || (this.signalIndexes.at(-1) || 0) > i - this.marker) {
        return;
      }
      const processing = this.buffer.slice(i - this.marker, i);
      const set = new Set(processing);
      if (processing.length === [...set.values()].length) {
        this.signalIndexes.push(i);
        this.foundSignals.push(processing.join());
      }
    });
  }
}


// Part one: 1760
const signalFinder = new SignalFinder(input, SIGNAL_MARKER);
console.log(signalFinder.signals.indexes[0]);

// Part two: 2974
const messageFinder = new SignalFinder(input, MESSAGE_MARKER)
console.log(messageFinder.signals.indexes[0]);
