import fs from 'fs';
import path from 'path';
import { combineScss } from '../utils';

test('combine_scss', () => {
  // mock response data
  const inputsStr = fs.readFileSync(path.join(__dirname, './mock/demo.scss'), { encoding: 'utf-8' });
  const _inputsStr = fs.readFileSync(path.join(__dirname, './mock/demo2.scss'), { encoding: 'utf-8' });
  const _inputsStr_ = fs.readFileSync(path.join(__dirname, './mock/demo3.scss'), { encoding: 'utf-8' });
  const { result } = {
    result: {
      GLOBAL_BASE_STYLE_CONFIG_KEY: inputsStr,
      GLOBAL_BASE_STYLE_KEY: _inputsStr,
      TEST: _inputsStr_,
    }
  }

  const finalScssFilePath = path.join(__dirname, '../final.scss'); 
  combineScss(result, finalScssFilePath);
  const finalScss = fs.readFileSync(finalScssFilePath);
  expect(finalScss).not.toBeNull();
})
