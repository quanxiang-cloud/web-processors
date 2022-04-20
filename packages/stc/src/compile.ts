import fs from 'fs';
import sass from 'sass';
import crypto from 'crypto';

function generateHash(value: string) {
  const cryptoCreate = crypto.createHash('md5');
  cryptoCreate.update(value);
  const hash = cryptoCreate.digest('hex');
  return hash;
}

export default function compile(scssFile: string, target: string) {
  const sourceString = fs.readFileSync(scssFile, { encoding: 'utf-8' });
  const result = sass.compileString(sourceString)
  fs.writeFileSync(target, result.css);
  
  const hash = generateHash(result.css);
  process.stdout.write(hash + ','+ target);
}
