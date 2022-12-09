const input = await Deno.readTextFile("2022/07/input.txt");

type FileName = string;
type FileSize = number;
type DirectoryName = string;

const sum = (values: number[]) => values.reduce((a, c) => a + c, 0);

class Directory {
  public name: string;
  public parent: Directory | null;
  public files: Map<FileName, FileSize>;
  public directories: Map<DirectoryName, Directory>;

  constructor(name: string, parent: Directory | null) {
    this.name = name;
    this.parent = parent;
    this.files = new Map();
    this.directories = new Map();
  }

  public addFile(name: FileName, size: FileSize) {
    this.files.set(name, size);
  }
  public addDir(name: DirectoryName) {
    const newDir = new Directory(name, this);
    this.directories.set(name, newDir);
    return newDir;
  }

  get size() {
    const fileSize = sum([...this.files.values()]);
    const dirSize = sum([...this.directories.values()].map((dir) => dir.size));
    return fileSize + dirSize;
  }

  get dirList(): { name: string; size: number }[] {
    const directories = [...this.directories.values()];
    const dirList = directories.map((dir) => ({
      name: dir.name,
      size: dir.size,
    }));
    const subDirs = directories.flatMap((dir) => dir.dirList);

    return [...dirList, ...subDirs];
  }
}

class FileSystem {
  public root: Directory;
  public currentDir: Directory;
  constructor() {
    this.root = new Directory("/", null);
    this.currentDir = this.root;
  }

  get size() {
    return this.root.size;
  }

  get dirList() {
    return this.root.dirList;
  }

  public cd(dirname: string) {
    dirname === "/"
      ? this.goToRoot()
      : dirname === ".."
      ? this.goUp()
      : this.goToOrCreateDir(dirname);
  }

  private goToRoot() {
    this.currentDir = this.root;
  }

  private goUp() {
    if (this.currentDir.parent === null) return;
    this.currentDir = this.currentDir.parent;
  }

  private goToOrCreateDir(dirname: string) {
    const dir = this.currentDir.directories.get(dirname);
    if (!dir) {
      const newDir = this.currentDir.addDir(dirname);
      this.currentDir = newDir;
    } else {
      this.currentDir = dir;
    }
  }
}

const fileSystem = new FileSystem();

const parseCommandLines = () => {
  const commands = input.split(`\n`);
  commands.forEach((command) => {
    if (command.startsWith("$ cd ")) {
      const dir = command.substring(5);
      fileSystem.cd(dir);
      return;
    }
    if (command.startsWith("$ ls") || command.startsWith("dir")) {
      return;
    }
    const [filesize, filename] = command.split(" ");
    fileSystem.currentDir.addFile(filename, Number(filesize));
  });
};

parseCommandLines();

const getDirsLessThan = (size: number) => {
  return fileSystem.dirList.filter((d) => d.size <= size);
};

// Part One: 1428881
const smallDirs = getDirsLessThan(100000);
const sumOfSmallDirs = sum(smallDirs.map(({ size }) => size));

console.log("Part One:", sumOfSmallDirs);

// Part Two
const getSmallestDir = () => {
  const diskSize = 70000000;
  const requiredSpace = 30000000;

  const currentAvailableSpace = diskSize - fileSystem.size;
  const minDirSize = requiredSpace - currentAvailableSpace;

  const dirsOverMinSize = fileSystem.dirList
    .filter((d) => d.size >= minDirSize)
    .map(({ size }) => size);
  const smallestDir = Math.min(...dirsOverMinSize);

  return smallestDir;
};

console.log("Part Two:", getSmallestDir());
