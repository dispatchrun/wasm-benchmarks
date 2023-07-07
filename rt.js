import { readFile } from 'node:fs/promises';
import { WASI } from 'wasi';
import { argv } from 'node:process';

const wasi = new WASI({
  version: 'preview1',
  args: argv.slice(2),
});

const wasm = await WebAssembly.compile(
  await readFile(new URL(argv[2], import.meta.url)),
);
const instance = await WebAssembly.instantiate(wasm, wasi.getImportObject());

wasi.start(instance);
