import fs from "fs";
import path from 'path';

export function parseParams() {
  let inputFile = process.env.INPUT_FILE;
  if (!inputFile) {
    // throw new Error("failed to get INPUT_FILE in process env");
    inputFile = path.join(__dirname, './mock-task-input.txt');
  }

  const inputsStr = fs.readFileSync(inputFile, { encoding: 'utf-8' });
  if (!inputsStr) {
    throw new Error("failed to parse body params of fetch scss str in persona");
  }
  
  const inputs = JSON.parse(inputsStr);
  const keys = Object.entries(inputs).reduce((acc: any, [_, {key, version}]: any) => {
    if (typeof key === 'string') {
      acc.push({ key, version })
    }

    if (Array.isArray(key)) {
      key.forEach((_key) => {
        acc.push({ key: _key, version })
      })
    }

    return acc;
  }, []);

  return { keys };
}

export function combineScss(data: Record<string, string>, finalScssFile: string) {
  const finalScssFileExisted = fs.existsSync(finalScssFile);
  if (finalScssFileExisted) {
    fs.writeFileSync(finalScssFile, '');
  }

  const options: fs.WriteFileOptions = { encoding: "utf8", flag: "a+" };
  const scssStrMap = Object.entries(data);
  scssStrMap.forEach(([_, scssStr], index) => {
    fs.writeFileSync(finalScssFile, JSON.parse(JSON.stringify(scssStr)), options);
    if (index === scssStrMap.length - 1) return;
    fs.writeFileSync(finalScssFile, '\n', options);
  })
}
