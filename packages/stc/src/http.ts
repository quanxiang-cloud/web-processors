import { request } from 'undici';

// post to get all scss file content
export async function fetch(api: string, params: Record<string, unknown>): Promise<Record<string, string>> {
  const { body, statusCode, } = await request(api, {
    method: 'POST',
    body: JSON.stringify(params),
    headers: { 'Content-Type': 'application/json' }
  });
  
  if (statusCode !== 200) {
    throw new Error('failed to fetch persona scss');
  }

  const { data } = await body.json();
  
  return Promise.resolve(data.result);
}

export async function fetchAllScssStr(api: string, params: Record<string, unknown>[]): Promise<any> {
  return Promise.all(params.map((keys) => fetch(api, keys)));
}
