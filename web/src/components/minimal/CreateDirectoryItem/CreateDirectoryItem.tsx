import { useContext, useEffect } from "react";
import ExplorerContext from "../../../utils/ExplorerContext";
import { AiOutlineFileAdd, AiOutlineFolderAdd } from "react-icons/ai";
import { IoClose } from "react-icons/io5";
import { useState } from "react";

const CreateDirectoryItem: React.FC = () => {
    const { createItem, path, showCreateItemMenu, setShowCreateItemMenu } = useContext(ExplorerContext)
    const [itemName, setItemName] = useState<string>("")

    useEffect(() => {
        setItemName("");
    }, [showCreateItemMenu])

    return showCreateItemMenu && (
        <section className='bg-black fixed inset-0 flex items-center justify-center w-full bg-opacity-60 z-30 animate-in fade-in duration-200'>
            <div className='flex flex-col bg-slate-800 border border-slate-700 rounded-xl w-full max-w-lg p-6 shadow-2xl animate-in zoom-in-95 duration-200'>
                {/* Header */}
                <div className="flex items-center justify-between mb-6">
                    <div>
                        <h2 className="text-2xl font-bold text-white">Create New Item</h2>
                        <p className="text-sm text-slate-400 mt-1">Enter a name for your file or folder</p>
                    </div>
                    <button
                        onClick={() => setShowCreateItemMenu?.(false)}
                        className="p-2 hover:bg-slate-700 rounded-lg transition-all text-slate-400 hover:text-white"
                        aria-label="Close"
                    >
                        <IoClose size={24} />
                    </button>
                </div>

                {/* Input Field */}
                <div className="mb-6">
                    <label htmlFor="itemName" className="text-slate-300 text-sm font-medium pl-1 mb-2 block">
                        Item Name
                    </label>
                    <input
                        id="itemName"
                        type="text"
                        className='bg-slate-700 border border-slate-600 focus:border-sky-500 focus:ring-2 focus:ring-sky-500/30 rounded-lg w-full px-4 py-3 text-white placeholder-slate-400 transition-all outline-none'
                        placeholder='example.txt or my-folder'
                        value={itemName}
                        onChange={(e) => {
                            setItemName(e.target.value)
                        }}
                        autoComplete="off"
                        autoFocus
                    />
                </div>

                {/* Action Buttons */}
                <div className="flex flex-col gap-3">
                    <div className="flex gap-3">
                        <button
                            className='flex gap-2 items-center justify-center flex-1 py-3 px-4 font-semibold rounded-lg transition-all bg-green-600 hover:bg-green-500 active:scale-[0.98] text-white shadow-lg hover:shadow-green-500/20'
                            onClick={() => {
                                setShowCreateItemMenu?.(false)
                                createItem(path, false, itemName)
                            }}
                        >
                            <AiOutlineFileAdd size={22} />
                            Create File
                        </button>
                        <button
                            className='flex gap-2 items-center justify-center flex-1 py-3 px-4 font-semibold rounded-lg transition-all bg-sky-600 hover:bg-sky-500 active:scale-[0.98] text-white shadow-lg hover:shadow-sky-500/20'
                            onClick={() => {
                                setShowCreateItemMenu?.(false)
                                createItem(path, true, itemName)
                            }}
                        >
                            <AiOutlineFolderAdd size={22} />
                            Create Folder
                        </button>
                    </div>
                    
                    <button
                        className='w-full py-3 px-4 font-semibold rounded-lg transition-all bg-slate-700 hover:bg-slate-600 border border-slate-600 text-white active:scale-[0.98]'
                        onClick={() => setShowCreateItemMenu?.(false)}
                    >
                        Cancel
                    </button>
                </div>
            </div>
        </section>
    )
}

export default CreateDirectoryItem