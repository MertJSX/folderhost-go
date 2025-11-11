import {
  FaFolder,
  FaFileAlt,
  FaFileImage,
  FaDatabase,
  FaFilePdf,
  FaFileArchive,
  FaHtml5,
  FaCss3,
  FaJava,
  FaMusic,
  FaPython,
  FaFileWord,
  FaFileExcel,
  FaFilePowerpoint,
  FaFolderOpen
} from "react-icons/fa";
import { FaGear } from "react-icons/fa6";
import { IoLogoJavascript } from "react-icons/io";
import { BiMoviePlay } from "react-icons/bi";
import { TbBrandCSharp, TbBrandPowershell } from "react-icons/tb";
import { VscJson } from "react-icons/vsc";
import { LuCodeXml } from "react-icons/lu";
import { RiPhpFill, RiPhpLine } from "react-icons/ri";
import { GrDocumentConfig } from "react-icons/gr";
import type { DirectoryItem } from "../types/DirectoryItem";

interface DirectoryItemIconProps {
  logoSize: number;
  itemInfo: DirectoryItem;
}

export const DirectoryItemIcon: React.FC<DirectoryItemIconProps> = ({ logoSize, itemInfo }) => {
  const icons: { [key: string]: JSX.Element } = {
    folder: <FaFolder size={logoSize} className='mx-2' style={{ color: "blue" }} />,
    folderOpen: <FaFolderOpen size={logoSize} className='mx-2' />,
    image: <FaFileImage size={logoSize} className='mx-2' />,
    pdf: <FaFilePdf size={logoSize} className='mx-2' />,
    archive: <FaFileArchive size={logoSize} className='mx-2' />,
    html: <FaHtml5 size={logoSize} className='mx-2' />,
    css: <FaCss3 size={logoSize} className='mx-2' />,
    audio: <FaMusic size={logoSize} className='mx-2' />,
    video: <BiMoviePlay size={logoSize} className='mx-2' />,
    java: <FaJava size={logoSize} className='mx-2' />,
    javascript: <IoLogoJavascript size={logoSize} className='mx-2' />,
    csharp: <TbBrandCSharp size={logoSize} className="mx-2" />,
    exe: <FaGear size={logoSize} className="mx-2" />,
    database: <FaDatabase size={logoSize} className="mx-2" />,
    python: <FaPython size={logoSize} className="mx-2" />,
    json: <VscJson size={logoSize} className="mx-2" />,
    msword: <FaFileWord size={logoSize} className="mx-2" />,
    msexcel: <FaFileExcel size={logoSize} className="mx-2" />,
    mspowerpoint: <FaFilePowerpoint size={logoSize} className="mx-2" />,
    code: <LuCodeXml size={logoSize} className="mx-2" />,
    php: <RiPhpFill size={logoSize + 5} className="mx-2" />,
    shell: <TbBrandPowershell size={logoSize} className="mx-2" />,
    config: <GrDocumentConfig size={logoSize} className="mx-2" />,
    default: <FaFileAlt size={logoSize} className='mx-2' />
  };

  const extensionMap: { [key: string]: string } = {
    png: 'image',
    jpg: 'image',
    jpeg: 'image',
    svg: 'image',
    webp: 'image',
    ico: 'image',
    pdf: 'pdf',
    rar: 'archive',
    zip: 'archive',
    tar: 'archive',
    gz: 'archive',
    '7z': 'archive',
    html: 'html',
    css: 'css',
    mp3: 'audio',
    opus: 'audio',
    wav: 'audio',
    flac: 'audio',
    alac: 'audio',
    mp4: 'video',
    java: 'java',
    jar: 'java',
    js: 'javascript',
    ts: 'javascript',
    tsx: 'javascript',
    jsx: 'javascript',
    cs: 'csharp',
    exe: 'exe',
    db: 'database',
    db3: 'database',
    mdb: 'database',
    accdb: 'database',
    dbf: 'database',
    sqlite: 'database',
    sqlite3: 'database',
    csv: 'database',
    tsv: 'database',
    sql: 'database',
    json: 'json',
    sh: 'shell',
    bat: 'shell',
    cmd: 'shell',
    bash: 'shell',
    docx: 'msword',
    dotx: 'msword',
    doc: 'msword',
    xls: 'msexcel',
    xlsx: 'msexcel',
    pptx: 'mspowerpoint',
    ppt: 'mspowerpoint',
    odt: 'mspowerpoint',
    xml: 'code',
    htmx: 'code',
    php: 'php',
    config: 'config',
    ini: 'config',
    yml: 'config',
    yaml: 'config',
    py: 'python'
  };


  if (!itemInfo) return icons.default;

  const { isDirectory, path, name } = itemInfo;
  const currentPath = path.slice(-1) === "/" ? path : path + "/";

  if (isDirectory) {
    return path === currentPath ? icons.folderOpen : icons.folder;
  }

  const extension = name?.split('.').pop()?.toLowerCase();
  if (!extension) {
    return null
  }
  const iconType = extensionMap[extension] || 'default';

  return icons[iconType];
};