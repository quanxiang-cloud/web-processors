import path from 'path';
import fs from 'fs';
import compile from '../compile';

test('compile', () => {
  const source = path.join(__dirname, './demo.scss')
  const target = path.join(__dirname, './demo.css')
  compile(source, target)

  const targetSource = fs.readFileSync(target, { encoding: 'utf-8' })
  expect(targetSource).toMatchInlineSnapshot(`
".foo {
  color: red;
}"
`)
})
