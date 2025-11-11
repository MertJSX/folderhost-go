import { useContext } from "react"
import ExplorerContext from "../../utils/ExplorerContext"
import { RiDeleteBin6Fill } from "react-icons/ri";
import { FaDownload, FaCopy, FaFileArchive } from "react-icons/fa";
import ExplorerRMItem from "./ExplorerRMItem";

interface ExplorerRightclickMenuProps {
    x: number, y: number
}

const ExplorerRightclickMenu: React.FC<ExplorerRightclickMenuProps> = ({ x, y }) => {
    const { itemInfo,
        permissions,
        showDisabled,
        deleteItem,
        createCopy,
        unzipProgress,
        startUnzipping,
        downloadFile,
        downloadProgress } = useContext(ExplorerContext)
    const iconSize = 20;
    return (
        <div
            style={{ top: `${y}px`, left: `${x}px` }}
            className='flex flex-col items-start bg-slate-900 rounded-lg text-white p-1 fixed z-20 w-44'
        >
            {
                permissions?.delete ?
                    <ExplorerRMItem
                        title='Click to delete.'
                        onClick={() => {
                            if (!window.confirm("Are you sure you want to delete this file?")) {
                                return;
                            }
                            deleteItem(itemInfo)
                        }}>
                        <RiDeleteBin6Fill size={iconSize} />Delete
                    </ExplorerRMItem>
                    : showDisabled === true ?
                        <ExplorerRMItem
                            isDisabled={true}
                            title="No permission">
                            <RiDeleteBin6Fill size={iconSize} />Delete
                        </ExplorerRMItem>
                        : null
            }

            {!downloadProgress && !itemInfo?.isDirectory ?
                permissions?.download_files ?
                    <ExplorerRMItem
                        title='Click to download.'
                        onClick={() => {
                            downloadFile(itemInfo?.path)
                        }}
                    ><FaDownload size={iconSize} /> Download</ExplorerRMItem> : showDisabled === true ?
                        <ExplorerRMItem
                            title='No permission!'
                            isDisabled={true}
                        ><FaDownload size={iconSize} />Download</ExplorerRMItem> : null
                : !itemInfo?.isDirectory && downloadProgress ?
                    <ExplorerRMItem
                        isDisabled={true}
                    ><FaDownload size={iconSize} />Downloading...</ExplorerRMItem> : null}

            <ExplorerRMItem
                title='Click to create a copy.'
                onClick={() => { createCopy(itemInfo) }}>
                <FaCopy size={iconSize} />Create Copy
            </ExplorerRMItem>
            {(itemInfo?.name.split(".").pop() === "zip" && unzipProgress === "") && !itemInfo?.isDirectory ?
                (permissions?.extract ?
                    <ExplorerRMItem
                        title='Click to unzip.'
                        onClick={() => {
                            startUnzipping()
                        }}
                    ><FaFileArchive size={iconSize} />Unzip</ExplorerRMItem> : showDisabled === true ?
                        <ExplorerRMItem
                            title='No permission!'
                            isDisabled={true}
                        ><FaFileArchive size={iconSize} />Unzip</ExplorerRMItem> : null)
                : (itemInfo?.name.split(".").pop() === "zip" && unzipProgress !== "") && !itemInfo?.isDirectory ?
                    <ExplorerRMItem
                        title='Unzipping...'
                        isDisabled={true}
                    ><FaFileArchive size={iconSize} />Unzipping...</ExplorerRMItem> : null
            }


        </div>
    )
}

export default ExplorerRightclickMenu