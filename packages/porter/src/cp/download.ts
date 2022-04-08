import fs from 'fs';
import path from 'path';
import { request } from 'undici';
import { pipeline } from 'stream/promises';

async function download(url: string, filePath: string): Promise<void> {
  const targetFolder = path.dirname(filePath);
  fs.mkdirSync(targetFolder)

  const response = await request(url)

  await pipeline(
    response.body,
    fs.createWriteStream(filePath),
  );

  console.log('file has been downloaded to:', filePath);
}

export default download;
