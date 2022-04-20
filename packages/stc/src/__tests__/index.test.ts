import fs from 'fs';
import path from "path";
import compile from "../compile";
import { fetchAllScssStr } from "../http";
import { combineScss, parseInput, buildURL } from "../utils";

test('works with promises', done => {
  function callback(data: any) {
    expect(data).not.toBeNull();
    done();
  }

  const params = parseInput();
  fetchAllScssStr(buildURL('/api/v1/persona/batchGetValue'),params).then((res) => {
    const isArray = Array.isArray(res);
    expect(isArray).toBeTruthy();
    expect(res).not.toBeNull();
    const finalScssFilePath = path.join(__dirname, '../final.scss'); 
    const finalCssFilePath = path.join(__dirname, '../final.css');
    combineScss(res, finalScssFilePath);
    compile(finalScssFilePath, finalCssFilePath);
    fs.rm(finalScssFilePath, () => null);
  }).then(callback);
});
