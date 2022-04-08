import undici from 'undici';
import fs from 'fs';
import { pipeline } from 'stream/promises';
import { lookup } from 'mime-types';
import { createGzip } from 'zlib';

async function upload(filePath: string, url: string): Promise<void> {
  const date = new Date().toUTCString();
  const contentType = lookup(filePath) || 'application/oct-stream';

  console.log(`uploading ${filePath}`, `to ${url}`);

  await pipeline(
    fs.createReadStream(filePath),
    createGzip(),
    undici.pipeline(
      url,
      {
        method: 'PUT',
        headers: {
          date,
          'content-type': contentType,
          'content-encoding': 'gzip',
          'cache-control': 'max-age=2592000',
        },
      },
      ({ statusCode, body }) => {
        if (statusCode >= 400) {
          console.error('failed to upload file', statusCode);
        }

        return body;
      },
    ),
  );

  console.log('file has been uploaded to:', url);
}

export default upload;
