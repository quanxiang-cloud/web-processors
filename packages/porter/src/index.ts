#!/usr/bin/env node

import yargs from 'yargs';
import { hideBin } from 'yargs/helpers';

import download from './cp/download';
import upload from './cp/upload';
import { buildURL, toAbsPath } from './utils';

const argv = yargs(hideBin(process.argv))
  .command('download', 'download file from file server')
  .command('upload', 'upload file to file server')
  .help('h').parseSync();

if (argv._.length < 3) {
  throw new Error('missing source or destination');
}

const sourcePath = argv._[1] as string;
const targetPath = argv._[2]  as string;

if (argv._[0] === 'download') {
  download(buildURL(sourcePath), toAbsPath(targetPath));
}

if (argv._[0] === 'upload') {
  upload(toAbsPath(sourcePath), buildURL(targetPath));
}
