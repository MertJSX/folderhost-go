import React from 'react'
import { Link } from 'react-router-dom'
import { BsEmojiFrown } from "react-icons/bs";

const NoPage = () => {
    return (
        <div>
            <div className="flex flex-col justify-center items-center bg-gray-700 w-1/2 mx-auto rounded-lg p-5 gap-4 mt-32">
                <BsEmojiFrown size={45} />
                <h1 className='text-3xl text-center'>Sorry! This page doesn't exist!</h1>
                <Link className='text-center font-bold text-2xl bg-sky-700 w-1/2 rounded-xl m-4 border-2 border-sky-500' to="/">Return to Home</Link>
            </div>
        </div>
    )
}

export default NoPage