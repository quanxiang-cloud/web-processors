import fs from "fs";
import path from 'path';

const hostname = process.env.hostname || 'http://persona.alpha'; 

export function buildURL(targetPath: string): string {
  return `${hostname}${targetPath}`;
}

export function parseInput() {
  // // debug use
  // let inputFile = process.env.INPUT_FILE;
  // if (!inputFile) {
  //   inputFile = path.join(__dirname, './mock-task-input.txt');
  // }
  const inputFile = process.env.INPUT_FILE;
  if (!inputFile) {
    throw new Error("failed to get INPUT_FILE in process env");
  }

  const inputsStr = fs.readFileSync(inputFile, { encoding: 'utf-8' });
  if (!inputsStr) {
    throw new Error("failed to parse body params of fetch scss str in persona");
  }
  
  const inputs = JSON.parse(inputsStr);
  const keys = Object.entries(inputs).map(([_, key]) => key).map(({key, version}: any) => {
    if (Array.isArray(key)) {
      const keys = key.map((key) => ({key, version}))
      return { keys };
    }

    return { keys: [{ key, version }] };
  })

  return keys;
}

export function getScssStrMap(data: Record<string, string>[]) {
  return data.reduce((acc, keyValues) => {
    Object.assign(acc, keyValues);
    return acc;
  }, {});
}

export function combineScss(data: Record<string, string>[], finalScssFile: string) {
  const finalScssFileExisted = fs.existsSync(finalScssFile);
  if (finalScssFileExisted) {
    fs.writeFileSync(finalScssFile, '');
  }

  const options: fs.WriteFileOptions = { encoding: "utf8", flag: "a+" };

  Object.entries(getScssStrMap(data)).forEach(([_, scssStr]) => {
    fs.writeFileSync(finalScssFile, JSON.parse(JSON.stringify(scssStr)), options);
    fs.writeFileSync(finalScssFile, '\n', options);
  })
}
