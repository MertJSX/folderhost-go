import Header from '../../components/Header/Header';
import FileExplorer from '../../components/FileExplorer/FileExplorer';
import OptionsBar from '../../components/Options/OptionsBar';
import ItemInfo from '../../components/DirItemInfo/ItemInfo';
import { useState, useEffect, useRef, useLayoutEffect } from 'react';
import Cookies from 'js-cookie';
import fileDownload from 'js-file-download';
import { useParams, useNavigate } from 'react-router-dom';
import ExplorerContext from '../../utils/ExplorerContext';
import MessageBox from '../../components/minimal/MessageBox/MessageBox';
import useWebSocket from '../../utils/useWebSocket';
import axiosInstance from '../../utils/axiosInstance';
import { type ExplorerContextType } from '../../types/ExplorerContextType';
import { type DirectoryItem } from '../../types/DirectoryItem';
import { type AccountPermissions } from '../../types/AccountPermissions';
import type { WebSocketResponseType } from '../../types/CodeEditorTypes';
import { getUserPermissions } from '../../utils/getUserPermissions';
import CreateDirectoryItem from '../../components/minimal/CreateDirectoryItem/CreateDirectoryItem';
const API_BASE_URL: string = import.meta.env.VITE_API_BASE_URL;

const ExplorerPage: React.FC = () => {
  const params = useParams();
  const navigate = useNavigate();
  const [path, setPath] = useState(params.path || './');
  const [permissions, setPermissions] = useState<AccountPermissions | null>(null);
  const [showDisabled, setShowDisabled] = useState(Cookies.get("show-disabled") === "true");
  const [disableCaching, setDisableCaching] = useState(Cookies.get("disable-caching") === "true");
  const [directory, setDir] = useState<DirectoryItem[]>([]);
  const [isEmpty, setIsEmpty] = useState(false);
  const [directoryInfo, setDirectoryInfo] = useState<DirectoryItem | null>(null);
  const [itemInfo, setItemInfo] = useState<DirectoryItem | null>(null);
  const [response, setRes] = useState<string>("");
  const [error, setError] = useState<string>("");
  const [downloadProgress, setDownloadProgress] = useState<number>(0);
  const [unzipProgress, setUnzipProgress] = useState<string>(""); // formatted unzip size (2.5 GB for example)
  const [messageBoxMsg, setMessageBoxMsg] = useState<string>("")
  const [messageBoxIsErr, setMessageBoxIsErr] = useState(false)
  const scrollIndex = useRef<number>(0)
  const [isDirLoading, setIsDirLoading] = useState<boolean>(false)
  const [showCreateItemMenu, setShowCreateItemMenu] = useState<boolean>(false);
  const [contextMenu, setContextMenu] = useState({
    show: false,
    x: 0,
    y: 0
  })
  const socket = useRef<WebSocket | null>(null);
  // Work states
  const [downloading, setDownloading] = useState(false);
  const [unzipping, setUnzipping] = useState(false);
  const [waitingResponse, setWaitingResponse] = useState(false);

  const {
    isConnected,
    isConnectedRef,
    messages,
    sendMessage,
    connectionError
  } = useWebSocket(path.slice(1), true)

  function getParent(filePath: string): string {
    let lastIndex = filePath.lastIndexOf('/');
    if (lastIndex === -1) return filePath;

    let item = filePath.slice(0, lastIndex);

    if (item.length > 1) {
      return item;
    } else {
      return filePath.slice(0, lastIndex + 1);
    }
  }

  useEffect(() => {
    navigate(`/explorer/${encodeURIComponent(path)}`, { replace: true });
  }, [path])

  useEffect(() => {
    if (isConnected && error == "WebSocket connection error!") {
      setError("")
      return
    }

    if (connectionError) {
      setError("WebSocket connection error!")
    }
  }, [isConnected])

  useEffect(() => {
    if (isConnectedRef.current) {
      let message: WebSocketResponseType;

      try {
        message = JSON.parse(messages[messages.length - 1] ?? "");
      } catch (err) {
        console.warn(messages[messages.length - 1]);

        console.error(err);
        return
      }

      switch (message.type) {
        case "unzip-progress":
          setUnzipProgress(message.totalSize ?? "")
          if (message.abortMsg) {
            readDir()
            setUnzipProgress("")
            setUnzipping(false)
            setError(message.abortMsg)
          }
          if (message.isCompleted) {
            readDir()
            setUnzipProgress("")
            setUnzipping(false)
            setRes("Unzip completed successfully!")
          }
          break;
        case "directory-update":
          if (unzipping) {
            return
          }
          readDir()
          break;
        case "error":
          setError(message.error ?? "Unknown error")
          break;
      }
    }
  }, [messages])

  useEffect(() => {
    if (Cookies.get("token")) {
      readDir()
      getUserPermissions((perms) => {
        setPermissions(perms)
      });
    } else {
      navigate("/login")
    }
  }, [])

  function declareError(error: string): void {
    setError(`${error}`);
  }

  function handleError(err: any, isErrorData?: boolean): void {
    setWaitingResponse(false);
    if (isErrorData && err.err) {
      declareError(err.err)
      return
    }
    if (isErrorData) {
      declareError("Unknown error!")
      return
    }
    if (err.response) {
      declareError(err.response.data.err)
    } else {
      declareError("Cannot connect to the server!")
    }
  }

  useEffect(() => {
    if (error) {
      setMessageBoxMsg(error)
      setMessageBoxIsErr(true)
    } else {
      setMessageBoxMsg(response)
      setMessageBoxIsErr(false)
    }
  }, [error, response])

  function waitPreviousAction(): void {
    setError("You must wait previus action to be completed!")
  }

  function moveItem(oldPath: string, newPath: string): void {
    if (downloading || waitingResponse || unzipping) {
      waitPreviousAction();
      return
    } else {
      setWaitingResponse(true);
    }
    axiosInstance.put(`/explorer/rename?oldFilepath=${oldPath.slice(1)}&newFilepath=${newPath.slice(1)}&type=move`)
      .then(() => {
        setWaitingResponse(false)
        readDir();
      }).catch((err) => {
        handleError(err)
      })
  }

  function renameItem(item: DirectoryItem, newName: string): void {
    if (downloading || waitingResponse || unzipping) {
      waitPreviousAction();
      return
    } else {
      setWaitingResponse(true);
    }
    const itemWithPath = item;
    let oldPath = itemWithPath.path.slice(1);
    let newPath = `${getParent(itemWithPath.path.slice(0, -1))}`; // /${newName}
    if (newPath.slice(-1) === "/") {
      newPath = newPath + newName
    } else {
      newPath = newPath + "/" + newName
    }
    axiosInstance.put(`/explorer/rename?oldFilepath=${oldPath}&newFilepath=${newPath.slice(1)}&type=rename`)
      .then(() => {
        setWaitingResponse(false)
        if (itemWithPath.isDirectory) {
          if (itemWithPath.path === `${path}/`) {
            readDir(false, newPath)

          } else {
            readDir()
          }
        } else {
          readDir()
        }
      }).catch((err) => {
        handleError(err)
      })
  }

  function downloadFile(filepath: string): void {
    if (downloading || waitingResponse || unzipping) {
      waitPreviousAction();
      return;
    }

    const encodedPath = encodeURIComponent(filepath.slice(1));

    axiosInstance.get(`/explorer/get-download-link?filepath=${encodedPath}`).then((data) => {
      if (!data.data.id) {
        return
      }
      const downloadUrl = `${API_BASE_URL}/download?id=${data.data.id}`;

      const a = document.createElement('a');
      a.href = downloadUrl;
      a.download = itemInfo?.name || '';
      a.style.display = 'none';
      document.body.appendChild(a);
      a.click();
      document.body.removeChild(a);
    }).catch((err) => {
      handleError(err)
    })

  }

  function deleteItem(item: DirectoryItem): void {
    if (downloading || waitingResponse || unzipping) {
      waitPreviousAction();
      return
    } else {
      setWaitingResponse(true);
    }
    const itemWithPath = item;
    let newPath = `${getParent(itemWithPath.path.slice(0, -1))}`;
    axiosInstance.delete(`/explorer/delete?path=${itemWithPath.path.slice(1)}`
    ).then((data) => {
      setWaitingResponse(false)
      if (itemWithPath.isDirectory) {
        if (itemWithPath.path === `${path}/`) {
          readDir(false, newPath)

        } else {
          readDir()
        }
      } else {
        readDir()
      }

      if (data.data.response) {
        setRes(data.data.response)
      }

    }).catch((err) => {
      handleError(err)
    })
  }

  function createCopy(item: DirectoryItem): void {
    if (downloading || waitingResponse || unzipping) {
      waitPreviousAction();
      return
    } else {
      setWaitingResponse(true);
    }
    const itemWithPath = item;
    axiosInstance.post(`/explorer/create-copy?path=${itemWithPath.path.slice(1)}`
    ).then((data) => {
      setWaitingResponse(false)
      readDir()
      if (data.data.response) {
        setRes(data.data.response)
      }
    }).catch((err) => {
      handleError(err)
    })
  }

  function createItem(itempath: string, isFolder: boolean, itemName: string): void {
    if (downloading || waitingResponse || unzipping) {
      waitPreviousAction();
      return
    } else {
      setWaitingResponse(true);
    }
    axiosInstance.post(`/explorer/create-item?path=${itempath.slice(1)}&isFolder=${isFolder}&itemName=${itemName}`)
      .then((data) => {
        setWaitingResponse(false)
        readDir()

        if (data.data.response) {
          setRes(data.data.response)
        }

      }).catch((err) => {
        handleError(err)
      })
  }

  function readDir(asParentPath?: boolean, pathInput?: string): void {
    setWaitingResponse(false);
    setDownloading(false);
    setIsDirLoading(true)
    if (asParentPath && path !== "./") {
      scrollIndex.current = 0;
      setPath(getParent(path));
      axiosInstance.get(`/explorer/read-dir?folder=${getParent(path).slice(1)}&mode=${Cookies.get("mode") || "Optimized mode"}${disableCaching ? "&caching=false" : ""}`
      )
        .then((data) => {
          if (data.data.items != null) {
            setDir(data.data.items)
          } else {
            setDir([])
          }
          setRes("");
          setIsEmpty(data.data.items == null);
          setDirectoryInfo(data.data.directoryInfo)
          setItemInfo(data.data.directoryInfo)
        }).catch((err) => {
          handleError(err)
        }).finally(() => {
          setIsDirLoading(false)
        })
      return;
    } else if (pathInput === undefined && !asParentPath) {
      axiosInstance.get(`/explorer/read-dir?folder=${path.slice(1)}&mode=${Cookies.get("mode") || "Optimized mode"}${disableCaching ? "&caching=false" : ""}`
      ).then((data) => {
        setIsEmpty(data.data.items == null);
        if (data.data.items != null) {
          setDir(data.data.items)
        } else {
          setDir([])
        }
        setDirectoryInfo(data.data.directoryInfo)
        setItemInfo(data.data.directoryInfo);
      }).catch((err) => {
        handleError(err)
      }).finally(() => {
        setIsDirLoading(false)
      })
      return;
    } else if (pathInput) {
      scrollIndex.current = 0;
      axiosInstance.get(`/explorer/read-dir?folder=${pathInput.slice(1)}&mode=${Cookies.get("mode") || "Optimized mode"}${disableCaching ? "&caching=false" : ""}`
      ).then((data) => {
        setPath(pathInput)
        setIsEmpty(data.data.items == null);
        if (data.data.items != null) {
          setDir(data.data.items)
        } else {
          setDir([])
        }
        setRes("")
        setDirectoryInfo(data.data.directoryInfo)
        setItemInfo(data.data.directoryInfo)
      }).catch((err) => {
        handleError(err)
      }).finally(() => {
        setIsDirLoading(false)
      })
    }
  }

  function startUnzipping(): void {
    if (downloading || waitingResponse || unzipping) {
      waitPreviousAction();
      return
    }
    if (socket !== null) {
      setUnzipping(true);
      sendMessage(JSON.stringify({
        type: "unzip",
        path: itemInfo?.path.slice(1)
      }))
    }
  }

  const contextValue: ExplorerContextType = {
    path: path,
    setPath: setPath,
    readDir: readDir,
    error: error,
    response: response,
    setShowDisabled: setShowDisabled,
    directory: directory,
    setDirectory: setDir,
    itemInfo: itemInfo,
    setItemInfo: setItemInfo,
    isEmpty: isEmpty,
    moveItem: moveItem,
    getParent: getParent,
    directoryInfo: directoryInfo,
    downloading: downloading,
    unzipping: unzipping,
    waitingResponse: waitingResponse,
    permissions: permissions,
    unzipProgress: unzipProgress,
    createCopy: createCopy,
    startUnzipping: startUnzipping,
    createItem: createItem,
    deleteItem: deleteItem,
    renameItem: renameItem,
    downloadFile: downloadFile,
    contextMenu: contextMenu,
    setContextMenu: setContextMenu,
    setMessageBoxMsg: setMessageBoxMsg,
    setError: setError,
    setRes: setRes,
    showDisabled: showDisabled,
    downloadProgress: downloadProgress,
    scrollIndex: scrollIndex,
    isDirLoading: isDirLoading,
    showCreateItemMenu: showCreateItemMenu,
    setShowCreateItemMenu: setShowCreateItemMenu,
    disableCaching: disableCaching,
    setDisableCaching: setDisableCaching
  };

  return (
    <ExplorerContext.Provider
      value={contextValue}>
      {/* <Header /> */}
      <div className='relative'>
        <OptionsBar />
        <main className="flex flex-row w-full justify-center pt-4 flex-wrap min-h-[calc(100vh-190px)]"
          onClick={(e) => {
            setContextMenu({ show: false, x: e.pageX, y: e.pageY })
          }}
        >
          <MessageBox message={messageBoxMsg} isErr={messageBoxIsErr} />
          <CreateDirectoryItem />
          <FileExplorer />
          {
            itemInfo && (
              <ItemInfo />
            )
          }
        </main>
      </div>
    </ExplorerContext.Provider>
  )
}

export default ExplorerPage