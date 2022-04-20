import fs from 'fs';
import path from 'path';
import { fetchScssStr } from '../http';
import { combineScss, parseParams } from '../utils';
import compile from '../compile';

it('works with promises', () => {
  const params = parseParams();
  expect.assertions(1);
  return fetchScssStr(params).then((res) => {
    expect(res).not.toBeNull();
    const finalScssFilePath = path.join(__dirname, '../final.scss'); 
    const finalCssFilePath = path.join(__dirname, '../final.css');
    combineScss(res, finalScssFilePath);
    compile(finalScssFilePath, finalCssFilePath);
    fs.rm(finalScssFilePath, () => null);
  });
});
