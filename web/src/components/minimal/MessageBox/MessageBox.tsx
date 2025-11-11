import { FaCheckCircle } from "react-icons/fa";
import { MdError } from "react-icons/md";
import { useContext } from "react";
import ExplorerContext from "../../../utils/ExplorerContext";

interface MessageBoxProps {
    message: string,
    isErr: boolean,
    setMessage?: React.Dispatch<React.SetStateAction<string>>
}

const MessageBox: React.FC<MessageBoxProps> = ({ message, isErr, setMessage }) => {
    const { setMessageBoxMsg, setError, setRes } = useContext(ExplorerContext)
    return message && (
        <section className='bg-black fixed inset-0 flex items-center justify-center w-full bg-opacity-60 z-30'>
            <div className={`flex flex-col justify-center items-center bg-slate-800 border-2 rounded-xl w-[550px] p-5 ${isErr ? "border-red-400" : "border-sky-400"}`}>
                {isErr ?
                    <MdError className="m-5" size={100} />
                    :
                    <FaCheckCircle className="m-5" size={100} />}
                <p className="text-center text-xl">{message}</p>
                <button
                    className={`w-full mt-5 p-1 transition-all ${isErr ? "bg-red-700 hover:bg-red-800" : "bg-sky-700 hover:bg-sky-800"}`}
                    onClick={() => {
                        setError("")
                        setRes("")
                        setMessageBoxMsg("")
                        if (setMessage) {
                            setMessage("")
                        }
                    }}
                >Okay</button>
            </div>
        </section>
    )
}

export default MessageBox