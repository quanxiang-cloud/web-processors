import fs from 'fs';
import sass from 'sass';

export default function compile(scssFile: string, target: string) {
  const sourceString = fs.readFileSync(scssFile, { encoding: 'utf-8' });
  const result = sass.compileString(sourceString)

  fs.writeFileSync(target, result.css);
}
