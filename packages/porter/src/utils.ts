import path from 'path';

export function buildURL(targetPath: string): string {
  return 'todo_hostname:port/path';
}

export function toAbsPath(targetPath: string): string {
  return path.isAbsolute(targetPath) ? targetPath : path.join(process.cwd(), targetPath);
}
