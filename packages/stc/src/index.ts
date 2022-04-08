#!/usr/bin/env node

import { hideBin } from 'yargs/helpers';

const argv = hideBin(process.argv);

if (argv.length !== 2) {
  throw new Error("invalid arguments");
}

console.log(argv[0], argv[1])
