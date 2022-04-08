import path from 'path';

const hostname = process.env.hostname;
const port = process.env.port || 80;

export function buildURL(targetPath: string): string {
  return `http://${hostname}:${port}${targetPath}`;
}

export function toAbsPath(targetPath: string): string {
  return path.isAbsolute(targetPath) ? targetPath : path.join(process.cwd(), targetPath);
}
