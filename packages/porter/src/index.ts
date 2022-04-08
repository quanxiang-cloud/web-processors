#!/usr/bin/env node

import yargs from 'yargs';
import { hideBin } from 'yargs/helpers';

import commandCP from './cp/command-cp';

const argv = yargs(hideBin(process.argv))
  .options({
    private: {
      type: 'boolean',
      description: 'use private bucket',
      global: true,
    },
  })
  .command('cp', 'copy file between local and file-server')
  .command('has', 'check if file exist')
  .help('h').parseSync();

if (argv._[0] === 'cp') {
  if (argv._.length < 3) {
    throw new Error('missing source or destination');
  }

  const sourcePath = argv._[1];
  const targetPath = argv._[2];

  commandCP(String(sourcePath), String(targetPath), !!argv.private);
}
