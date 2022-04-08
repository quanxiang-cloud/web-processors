import uploadFile from './upload';
import download from './download';
import { buildURL, toAbsPath } from '../utils';

function commandCP(sourcePath: string, targetPath: string, isPrivate: boolean): Promise<void> {
  console.log('isPrivate:', isPrivate)

  if (!sourcePath.startsWith('fs:')) {
    return uploadFile(toAbsPath(sourcePath), buildURL(targetPath));
  }


  return download(buildURL(sourcePath), toAbsPath(targetPath));
}

export default commandCP;
