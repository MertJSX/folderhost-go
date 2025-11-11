import Cookies from 'js-cookie';
import { useNavigate, Link, useLocation } from 'react-router-dom';
import { MdExplore, MdMiscellaneousServices } from "react-icons/md";
import { FaUserFriends, FaPencilAlt, FaFolder, FaUser, FaBars, FaTimes } from "react-icons/fa"
import { FaArrowRotateLeft } from "react-icons/fa6"
import { useCallback, useLayoutEffect, useState } from 'react';
import axiosInstance from '../../utils/axiosInstance';

const Header = () => {
    let navigate = useNavigate();
    let location = useLocation();
    const [username, setUsername] = useState<string>(Cookies.get("username") as string);
    const [isMobileMenuOpen, setIsMobileMenuOpen] = useState(false);

    useLayoutEffect(() => {
        if (Cookies.get("username")) return

        fetchUserInfo()
    }, []);

    const fetchUserInfo = useCallback(() => {
        axiosInstance.get('/user-info'
        ).then((response) => {
            if (response.data.username) {
                setUsername(response.data.username);
            }
        }).catch((err) => {
            console.error('Failed to fetch user info:', err);
        });
    }, [])

    const isActiveLink = (path: string) => {
        return location.pathname.startsWith(path);
    }

    const getNavLinkClass = (path: string) => {
        const baseClasses = 'text-base flex items-center gap-2 px-5 py-3 transition-all border-b-2';
        const activeClasses = 'text-sky-400 border-sky-500';
        const inactiveClasses = 'text-white border-transparent hover:bg-slate-700 hover:text-sky-400 hover:border-sky-500';
        
        return `${baseClasses} ${isActiveLink(path) ? activeClasses : inactiveClasses}`;
    }

    const closeMobileMenu = () => {
        setIsMobileMenuOpen(false);
    }

    return (
        <div className='flex flex-col items-center justify-center bg-slate-800 sticky left-0 right-0 top-0 w-full border-b-2 border-slate-700 shadow-lg z-50'>
            {/* Desktop Header */}
            <section className='hidden md:flex flex-row items-center justify-between w-full px-6 py-2'>
                {/* Logo Section */}
                <div className="flex items-center gap-3">
                    <div className="p-2 bg-blue-500 rounded-lg">
                        <FaFolder className="w-5 h-5 text-white" />
                    </div>
                    <h1 className="font-extrabold text-2xl italic select-none text-white">
                        FolderHost
                    </h1>
                </div>

                {/* User Info Section */}
                <div className="flex items-center gap-3 px-4 py-3 rounded-lg">
                    <div className="flex items-center gap-2">
                        <div className="p-1.5 bg-blue-500 rounded-full">
                            <FaUser className="w-3.5 h-3.5 text-sky-400" />
                        </div>
                        <span className="text-white font-medium">{username}</span>
                    </div>
                </div>

                {/* Logout Button */}
                <button
                    className='bg-slate-700 px-5 py-2 rounded-lg border border-slate-600 hover:border-sky-500 hover:bg-slate-600 active:bg-slate-700 font-semibold transition-all text-white'
                    onClick={() => {
                        Cookies.remove("token");
                        navigate("/login");
                    }}
                >
                    Logout
                </button>
            </section>

            {/* Mobile Header */}
            <section className='flex md:hidden flex-col items-center w-full px-6 py-3'>
                {/* Logo */}
                <div className="flex items-center gap-3 mb-4">
                    <div className="p-2 bg-blue-500 rounded-lg">
                        <FaFolder className="w-5 h-5 text-white" />
                    </div>
                    <h1 className="font-extrabold text-2xl italic select-none text-white">
                        FolderHost
                    </h1>
                </div>

                {/* User Info and Logout */}
                <div className="flex items-center justify-between w-full">
                    {/* User Info */}
                    <div className="flex items-center gap-3 px-4 py-2 rounded-lg">
                        <div className="flex items-center gap-2">
                            <div className="p-1.5 bg-blue-500 rounded-full">
                                <FaUser className="w-3.5 h-3.5 text-sky-400" />
                            </div>
                            <span className="text-white font-medium">{username}</span>
                        </div>
                    </div>

                    {/* Logout Button and Mobile Menu Button */}
                    <div className="flex items-center gap-2">
                        <button
                            className='bg-slate-700 px-4 py-2 rounded-lg border border-slate-600 hover:border-sky-500 hover:bg-slate-600 active:bg-slate-700 font-semibold transition-all text-white text-sm'
                            onClick={() => {
                                Cookies.remove("token");
                                navigate("/login");
                            }}
                        >
                            Logout
                        </button>
                        
                        {/* Mobile Menu Button */}
                        <button 
                            className="p-2 text-white hover:bg-slate-700 rounded-lg transition-all border border-slate-600"
                            onClick={() => setIsMobileMenuOpen(!isMobileMenuOpen)}
                        >
                            {isMobileMenuOpen ? <FaTimes className="w-5 h-5" /> : <FaBars className="w-5 h-5" />}
                        </button>
                    </div>
                </div>
            </section>

            {/* Desktop Navigation */}
            <nav className='hidden md:flex flex-row justify-center items-center gap-1 w-full border-t border-slate-700/50'>
                <Link
                    className={getNavLinkClass('/explorer')}
                    to={"/explorer/.%2F"}>
                    <MdExplore className="w-5 h-5" />
                    Explorer
                </Link>
                <Link
                    className={getNavLinkClass('/services')}
                    to={"/services"}>
                    <MdMiscellaneousServices className="w-5 h-5" />
                    Services
                </Link>
                <Link
                    className={getNavLinkClass('/recovery')}
                    to={"/recovery"}>
                    <FaArrowRotateLeft className="w-4 h-4" />
                    Recovery
                </Link>
                <Link
                    className={getNavLinkClass('/users')}
                    to={"/users"}>
                    <FaUserFriends className="w-5 h-5" />
                    Users
                </Link>
                <Link
                    className={getNavLinkClass('/logs')}
                    to={"/logs"}>
                    <FaPencilAlt className="w-4 h-4" />
                    Logs
                </Link>
            </nav>

            {/* Mobile Navigation Menu */}
            {isMobileMenuOpen && (
                <div className="md:hidden w-full bg-slate-800 border-t border-slate-700/50 animate-fadeIn">
                    <nav className='flex flex-col w-full'>
                        <Link
                            className={getNavLinkClass('/explorer')}
                            to={"/explorer/.%2F"}
                            onClick={closeMobileMenu}>
                            <MdExplore className="w-5 h-5" />
                            Explorer
                        </Link>
                        <Link
                            className={getNavLinkClass('/services')}
                            to={"/services"}
                            onClick={closeMobileMenu}>
                            <MdMiscellaneousServices className="w-5 h-5" />
                            Services
                        </Link>
                        <Link
                            className={getNavLinkClass('/recovery')}
                            to={"/recovery"}
                            onClick={closeMobileMenu}>
                            <FaArrowRotateLeft className="w-4 h-4" />
                            Recovery
                        </Link>
                        <Link
                            className={getNavLinkClass('/users')}
                            to={"/users"}
                            onClick={closeMobileMenu}>
                            <FaUserFriends className="w-5 h-5" />
                            Users
                        </Link>
                        <Link
                            className={getNavLinkClass('/logs')}
                            to={"/logs"}
                            onClick={closeMobileMenu}>
                            <FaPencilAlt className="w-4 h-4" />
                            Logs
                        </Link>
                    </nav>
                </div>
            )}
        </div>
    )
}

export default Header