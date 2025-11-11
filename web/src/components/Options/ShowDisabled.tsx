import { useState, useRef, useEffect } from 'react';
import { BiShow } from "react-icons/bi";
import { BiHide } from "react-icons/bi";
import { FaArrowLeft } from "react-icons/fa";
import { FaArrowDown } from "react-icons/fa";
import Cookies from 'js-cookie';

interface ShowDisabledProps {
    setShowDisabled: React.Dispatch<React.SetStateAction<boolean>>
}

const ShowDisabled: React.FC<ShowDisabledProps> = ({ setShowDisabled }) => {
    const [isOpen, setIsOpen] = useState<boolean>(false);
    const [selectedOption, setSelectedOption] = useState<string | null>(null);
    const dropdownRef = useRef<HTMLDivElement>(null);
    const iconSize = 100;
    const buttonSize = 20;

    const options = [
        {
            mode: "show",
            title: 'Show disabled',
            description: "This option will show disabled buttons that you don't have access."
        },
        {
            mode: "hide",
            title: 'Hide disabled',
            description: "This option will hide disabled buttons that you don't have access. (Recommended)"
        }

    ];

    const toggleDropdown = () => {
        setIsOpen(!isOpen);
    };

    const handleOptionClick = (option: {
        mode: string;
        title: string;
        description: string;
    }) => {
        if (option.mode === "show") {
            Cookies.set("show-disabled", "true");
        } else {
            Cookies.set("show-disabled", "false");
        }
        setSelectedOption(option.title);
        setIsOpen(false);
        setShowDisabled(Boolean(Cookies.get("show-disabled")) ?? false)
    };

    const handleOutsideClick = (event: MouseEvent) => {
        if (dropdownRef.current && !dropdownRef.current.contains(event.target as Node)) {
            setIsOpen(false);
        }
    };

    useEffect(() => {
        document.addEventListener('mousedown', handleOutsideClick);
        return () => {
            document.removeEventListener('mousedown', handleOutsideClick);
        };
    }, []);

    useEffect(() => {
        console.log(Cookies.get("show-disabled"));
        console.log(typeof Cookies.get("show-disabled"));
        let option = Cookies.get("show-disabled");


        if (option !== undefined) {
            if (option === "true") {
                setSelectedOption(options[0]?.title ?? "Show disabled")
            } else {
                setSelectedOption(options[1]?.title ?? "Hide disabled")
            }
        } else {
            setSelectedOption(options[1]?.title ?? "Hide disabled")
            Cookies.set("show-disabled", "false")
        }

    }, [])

    return (
        <div className="relative inline-block text-left pl-5" ref={dropdownRef}>
            <button
                onClick={toggleDropdown}
                className="inline-flex justify-between items-center w-48 px-3 py-2 text-sm font-medium text-white bg-gray-700 rounded-md hover:bg-gray-600 focus:outline-none"
            >
                {selectedOption === "Show disabled" ?
                    <BiShow size={buttonSize} /> :
                    <BiHide size={buttonSize} />
                }
                {selectedOption}
                {isOpen ? <FaArrowDown size={buttonSize - 5} /> : <FaArrowLeft size={buttonSize - 5} />}
            </button>

            {isOpen && (
                <div className="origin-top-left absolute left-0 mt-2 w-72 rounded-md shadow-lg bg-gray-800 ring-1 ring-black ring-opacity-5 z-10">
                    <div className="py-1">
                        {options.map((option, index) => (
                            <div
                                key={index}
                                className="flex items-start px-4 py-2 text-sm text-white hover:bg-gray-900 cursor-pointer border-2 border-sky-700"
                                onClick={() => handleOptionClick(option)}
                            >
                                {option.mode === "show" ? (
                                    <BiShow className="p-2" size={iconSize} />
                                ) : (
                                    <BiHide className="p-2" size={iconSize} />
                                )}
                                <div className="ml-4">
                                    <h1 className="text-lg">{option.title}</h1>
                                    <p className="text-sm">{option.description}</p>
                                </div>
                            </div>
                        ))}
                    </div>
                </div>
            )}
        </div>
    );
};

export default ShowDisabled;
