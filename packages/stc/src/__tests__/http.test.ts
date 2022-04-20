import { fetch, fetchAllScssStr } from '../http';
import { buildURL, parseInput } from '../utils';

const apiURL = buildURL('/api/v1/persona/batchGetValue');

it('works with url', () => {
  expect.assertions(1);
  return fetch(apiURL, {keys:[{design_token_key:{key:'GLOBAL_BASE_STYLE_CONFIG',version:'0.1.0'}}]}).then((res) => expect(res).not.toBeNull());
});

test('works with promises', done => {
  function callback(data: unknown) {
    expect(data).not.toBeNull();
    done();
  }

  const params = parseInput();
  fetchAllScssStr(apiURL, params).then((res) => {
    const isArray = Array.isArray(res);
    expect(isArray).toBeTruthy();
  }).then(callback);
});
