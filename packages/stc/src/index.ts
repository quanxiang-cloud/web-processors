#!/usr/bin/env node
import fs from 'fs';
import compile from './compile';
import { fetchScssStr } from './http';
import { combineScss, parseParams } from './utils';

try {
  const params = parseParams();
  // post to get all scss file content
  fetchScssStr(params).then((res) => {
    const finalScssFilePath = './final.scss'
    const finalCssFilePath = './final.css';
    combineScss(res, finalScssFilePath);
    compile(finalScssFilePath, finalCssFilePath);
    fs.rm(finalScssFilePath, () => null);
  });
} catch(error) {
  process.stderr.write(String(error));
}
