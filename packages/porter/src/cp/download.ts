import fs from 'fs';
import path from 'path';
import { request } from 'undici';
import { pipeline } from 'stream/promises';

async function download(url: string, downloadTo: string): Promise<void> {
  const targetFolder = path.dirname(downloadTo);
  if (!fs.existsSync(targetFolder)) {
    console.log('create target folder recursively:', targetFolder)
    fs.mkdirSync(targetFolder, { recursive: true })
  }
  console.log('downloading file:', url);
  const response = await request(url)

  if (response.statusCode !== 200) {
    throw new Error(`download failed: ${response.statusCode}`);
  }

  await pipeline(
    response.body,
    fs.createWriteStream(downloadTo),
  );

  console.log('file has been downloaded to:', downloadTo);
}

export default download;
