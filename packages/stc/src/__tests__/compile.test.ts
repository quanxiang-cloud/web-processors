import path from 'path';
import fs from 'fs';
import compile from '../compile';


// utils combine scss test before this test
test('compile', () => {
  const source = path.join(__dirname, '../final.scss')
  const target = path.join(__dirname, '../final.css')
  compile(source, target)

  const targetSource = fs.readFileSync(target, { encoding: 'utf-8' })
  expect(targetSource).toMatchInlineSnapshot(`
".foo {
  color: red;
}

.test {
  width: \\"100px\\";
  color: var(\\"--red-600\\");
}

.unkonw {
  padding-top: 20%;
}"
`)
})
