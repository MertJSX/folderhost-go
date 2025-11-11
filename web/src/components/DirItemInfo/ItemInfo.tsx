import { useRef, useContext } from 'react'
import moment from 'moment'
import { FaFolder } from "react-icons/fa";
import { FaFileAlt } from "react-icons/fa";
import { FaFileImage } from "react-icons/fa";
import { FaFilePdf } from "react-icons/fa6";
import { FaFileArchive } from "react-icons/fa";
import { FaHtml5 } from "react-icons/fa";
import { FaCss3 } from "react-icons/fa";
import { IoLogoJavascript } from "react-icons/io";
import { FaFolderOpen } from "react-icons/fa6";
import { FaJava } from "react-icons/fa";
import { FaMusic } from "react-icons/fa";
import { BiMoviePlay } from "react-icons/bi";
import Cookies from 'js-cookie';
import ExplorerContext from '../../utils/ExplorerContext';
import { type ExplorerContextType } from '../../types/ExplorerContextType';
import { DirectoryItemIcon } from '../../utils/DirectoryItemIcon';

const ItemInfo = () => {
  const renameInput = useRef<HTMLInputElement>(null)
  const logoSize = 75;
  const {
    itemInfo, setItemInfo, renameItem, downloadFile, downloadProgress, deleteItem, createCopy, path, createItem, unzipProgress, permissions, showDisabled, startUnzipping, directoryInfo
  } = useContext<ExplorerContextType>(ExplorerContext)

  return (
    <div className='flex flex-col items-center justify-center w-1/3 mx-auto min-w-[320px] max-w-[30%] min-h-[600px] h-[700px] max-h-[800px]'>
      <div className='flex flex-col bg-gray-800 items-center justify-center gap-3 rounded-xl shadow-2xl w-full h-auto p-4 min-h-[400px]'>
        { itemInfo ? <DirectoryItemIcon itemInfo={itemInfo} logoSize={logoSize} /> : null}
        {
          itemInfo?.path === "./" ?
            <h1
              title='The base folder cannot be renamed!'
              className='bg-transparent text-xl text-amber-300 font-bold text-wrap break-all text-center'
            >
              {itemInfo?.name}
            </h1> :
            !permissions?.rename ?
              <h1
                title='No permission to rename this one!'
                className='bg-transparent text-xl text-amber-300 font-bold text-wrap break-all text-center'
              >
                {itemInfo?.name}
              </h1> :
              <input
                type='text'
                ref={renameInput}
                value={itemInfo?.name}
                title='Rename item'
                onChange={(e) => {
                  setItemInfo((prev) => {
                    if (!prev) return prev
                    return { ...prev, name: e.target.value }
                  });
                }}
                onKeyDown={(e) => {
                  if (e.key === "/" ||
                    e.key === "\\" ||
                    e.key === ":" ||
                    e.key === "*" ||
                    e.key === "?" ||
                    e.key === "<" ||
                    e.key === ">" ||
                    e.key === "|"
                  ) {
                    e.preventDefault()
                  }
                  if (e.key === "Enter") {
                    renameItem(itemInfo, itemInfo?.name)
                    renameInput.current?.blur();
                  }
                }}
                className='bg-transparent text-xl text-amber-300 font-bold text-wrap break-all text-center'
              />
        }


        <h1 className='text-gray-400 text-center break-all text-wrap'>
          Path: <span className="text-amber-200">{itemInfo?.path}</span>
        </h1>
        {
          (itemInfo?.size && (Cookies.get("mode") !== "Quality mode" && Cookies.get("mode") !== "Balanced mode")) && (itemInfo?.size && itemInfo.isDirectory) ? (
            null
          ) : <h1 className='text-gray-400'>Size:
            <span className='text-amber-200'>
              {
                (itemInfo?.size === "N/A" && Cookies.get("mode") === "Quality mode") || (itemInfo?.size === "N/A" && !itemInfo?.isDirectory && Cookies.get("mode") === "Balanced mode") ?
                  " 0 Bytes" : ` ${itemInfo?.size}`
              }
            </span>
          </h1>
        }
        {
          itemInfo?.storage_limit && path == "./" ?
            <h1 className='text-gray-400'>
              {"Limit: "}<span className='text-emerald-300'>{itemInfo?.storage_limit}</span>
            </h1> : null
        }
        <h1 className='text-gray-400'>
          Modified: <span className='text-gray-300'>{moment(itemInfo?.dateModified).format("Do MMMM YYYY HH:mm")}</span>
        </h1>
        {!itemInfo?.isDirectory ? (
          <div className='flex flex-col gap-2 w-5/6'>
            {
              permissions?.delete ?
                <button
                  className='bg-red-600 px-6 font-bold rounded-xl'
                  title='Click to delete.'
                  onClick={() => {
                    if (!window.confirm("Are you sure you want to delete this file?")) {
                      return;
                    }
                    deleteItem(itemInfo)
                  }}
                >Delete file</button> : showDisabled === true ?
                  <button
                    className='bg-red-600 px-6 font-bold rounded-xl opacity-50'
                    title='No permission!'
                    disabled
                  >Delete file</button> : null
            }
            {
              permissions?.copy ?
                <button
                  className='bg-sky-600 px-6 font-bold rounded-xl'
                  title='Click to create copy.'
                  onClick={() => {
                    createCopy(itemInfo)
                  }}
                >Create copy</button> : showDisabled === true ?
                  <button
                    className='bg-sky-600 px-6 font-bold rounded-xl opacity-50'
                    title='No permission!'
                    disabled
                  >Create copy</button> : null
            }
            {itemInfo?.name.split(".").pop() === "zip" && unzipProgress === "" ?
              (permissions?.extract ?
                <button
                  className='bg-yellow-600 px-6 font-bold rounded-xl'
                  title='Click to unzip.'
                  onClick={() => {
                    startUnzipping()
                  }}
                >Unzip</button> : showDisabled === true ?
                  <button
                    className='bg-yellow-600 px-6 font-bold rounded-xl opacity-50'
                    title='No permission!'
                    disabled
                  >Unzip</button> : null)
              : itemInfo?.name.split(".").pop() === "zip" && unzipProgress !== "" ?
                <div>
                  <h1 className='text-center'>Unzipping...</h1>
                  <h1 className='text-center text-xl'>Progress: <span className='text-sky-300'>{unzipProgress}</span></h1>
                </div> : null
            }
            {!downloadProgress ?
              permissions?.download_files ?
                <button
                  className='bg-emerald-600 px-6 font-bold rounded-xl'
                  title='Click to download.'
                  onClick={() => {
                    downloadFile(itemInfo?.path)
                  }}
                >Download</button> : showDisabled === true ?
                  <button
                    className='bg-emerald-600 px-6 font-bold rounded-xl opacity-50'
                    title='No permission!'
                    disabled
                  >Download</button> : null
              :
              <div>
                <h1 className="text-center">{downloadProgress === 100 ? "Downloaded" : "Downloading..."}</h1>
                <div className="w-full bg-gray-200 rounded-full h-2.5 dark:bg-gray-800">
                  <div className="bg-emerald-600 h-2.5 rounded-full" style={{ width: `${downloadProgress}%` }} />
                </div>
              </div>

            }
            {
              permissions?.read_files ?
                <a
                  href={`/editor/${encodeURIComponent(itemInfo?.path ?? "./")}`}
                  target="_blank" rel="noreferrer"
                  className='bg-sky-700 px-6 font-bold text-center rounded-xl'
                >Open Editor</a> : showDisabled === true ?
                  <button
                    className='bg-sky-700 px-6 font-bold text-center rounded-xl opacity-50'
                    disabled
                  >Open Editor</button> : null

            }
          </div>

        ) :
          <div className="flex flex-col gap-2 w-5/6">
            {
              itemInfo?.path !== "./" && permissions?.delete && itemInfo.path !== directoryInfo?.path ?
                <button
                  className='bg-red-600 hover:bg-red-700 px-6 font-bold rounded-xl'
                  title='Click to delete.'
                  onClick={() => {
                    if (!window.confirm("Are you sure you want to delete this directory?")) {
                      return;
                    }
                    deleteItem(itemInfo)
                  }}
                >Delete directory</button> :
                itemInfo?.path !== "./" && showDisabled === true ?
                  <button
                    className='bg-red-600 px-6 font-bold rounded-xl opacity-50'
                    title='No permission to delete!'
                    disabled
                  >Delete directory</button> : null
            }
            {
              itemInfo?.path !== "./" && permissions?.copy && itemInfo?.path !== "./" && itemInfo.path !== directoryInfo?.path ?
                <button
                  className='bg-sky-600 px-6 font-bold rounded-xl'
                  title='Click to create copy.'
                  onClick={() => {
                    createCopy(itemInfo)
                  }}
                >Create copy</button>
                : itemInfo?.path !== "./" && itemInfo?.path !== (path.slice(-1) === "/" ? path : path + "/") && showDisabled === true ?
                  <button
                    className='bg-sky-600 px-6 font-bold rounded-xl opacity-50'
                    title='No permission!'
                    disabled
                  >Create copy</button> : null
            }
            {
              itemInfo?.path === directoryInfo?.path ?
                <div className='w-full flex flex-col gap-2'>
                  {
                    permissions?.upload_files ?
                      <a
                        className='bg-purple-600 px-6 font-bold rounded-xl text-center'
                        title='Click to upload.'
                        target='_blank' rel="noreferrer"
                        href={`/upload/${encodeURIComponent(itemInfo.path)}`}
                      >Upload new file</a> : showDisabled === true ?
                        <button
                          className='bg-purple-600 px-6 font-bold rounded-xl text-center opacity-50 cursor-not-allowed'
                          title='No permission to upload!'
                        >Upload new file</button> : null
                  }

                </div> : null
            }
          </div>
        }
      </div>
    </div>
  )
}

export default ItemInfo