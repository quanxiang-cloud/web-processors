#!/usr/bin/env node
import fs from 'fs';
import compile from './compile';
import { fetchAllScssStr } from './http';
import { buildURL, combineScss, parseInput } from './utils';

try {
  const params = parseInput();
  const apiUrl = buildURL('/api/v1/persona/batchGetValue');
  // post to get all scss file content
  fetchAllScssStr(apiUrl, params).then((res) => {
    const finalScssFilePath = './final.scss'
    const finalCssFilePath = './final.css';
    combineScss(res, finalScssFilePath);
    compile(finalScssFilePath, finalCssFilePath);
    fs.rm(finalScssFilePath, () => null);
  });
} catch(error) {
  throw error;
}
