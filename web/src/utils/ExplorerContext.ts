import { createContext, type Context } from "react"
import { type ExplorerContextType } from "../types/ExplorerContextType.js";
import { type ContextMenuType } from "../types/ContextMenuType.js";
import type { DirectoryItem } from "../types/DirectoryItem.js";

const ExplorerContext: Context<ExplorerContextType> = createContext<ExplorerContextType>({
    path: "",
    setPath: ((value: string | ((prev: string) => string)) => {}) as React.Dispatch<React.SetStateAction<string>>,
    readDir: () => {},
    error: "",
    response: "",
    setShowDisabled: ((value: boolean | ((prev: boolean) => boolean)) => {}) as React.Dispatch<React.SetStateAction<boolean>>,
    directory: [],
    setDirectory: ((value: DirectoryItem[] | ((prev: DirectoryItem[]) => DirectoryItem[])) => {}) as React.Dispatch<React.SetStateAction<DirectoryItem[]>>,
    itemInfo: null,
    setItemInfo: ((value: DirectoryItem | ((prev: DirectoryItem) => DirectoryItem)) => {}) as React.Dispatch<React.SetStateAction<DirectoryItem | null>>,
    isEmpty: false,
    moveItem: () => {},
    getParent: () => "",
    directoryInfo: null,
    downloading: false,
    unzipping: false,
    waitingResponse: false,
    permissions: null,
    unzipProgress: "",
    createCopy: () => {},
    startUnzipping: () => {},
    createItem: () => {},
    deleteItem: () => {},
    renameItem: () => {},
    downloadFile: () => {},
    contextMenu: {
        show: false,
        x: 0,
        y: 0
    },
    setContextMenu: ((value: ContextMenuType) => {}) as React.Dispatch<React.SetStateAction<ContextMenuType>>,
    setMessageBoxMsg: ((value: string | ((prev: string) => string)) => {}) as React.Dispatch<React.SetStateAction<string>>,
    setError: ((value: string | ((prev: string) => string)) => {}) as React.Dispatch<React.SetStateAction<string>>,
    setRes: ((value: string | ((prev: string) => string)) => {}) as React.Dispatch<React.SetStateAction<string>>,
    showDisabled: false,
    downloadProgress: 0,
    scrollIndex: {current: 0},
    isDirLoading: false
});

export default ExplorerContext;