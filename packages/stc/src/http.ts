import { request } from 'undici';

const hostname = process.env.PERSONA_HOSTNAME || 'persona.alpha';

// post to get all scss file content
export async function fetchScssStr(params: any): Promise<Record<string, string>> {
  const api = `http://${hostname}/api/v1/persona/batchGetValue`;
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
