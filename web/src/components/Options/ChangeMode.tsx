import { useState, useRef, useEffect } from 'react';
import { PiSpeedometerBold } from "react-icons/pi";
import { IoDiamond } from "react-icons/io5";
import { FaBalanceScale } from "react-icons/fa";
import { FaArrowLeft } from "react-icons/fa";
import { FaArrowDown } from "react-icons/fa";
import Cookies from 'js-cookie';

const ChangeMode = () => {
    const [isOpen, setIsOpen] = useState<boolean>(false);
    const [selectedOption, setSelectedOption] = useState<string>(Cookies.get("mode") || "Optimized mode");
    const dropdownRef = useRef<HTMLDivElement>(null);
    const iconSize = 100;
    const onButtonSize = 20;

    const options = [
        {
            mode: "optimized",
            title: 'Optimized mode',
            description: "You will get faster loading, but you won't be able to get folder size information!"
        },
        {
            mode: "quality",
            title: 'Quality mode',
            description: 'This is the slowest plan, but you will get all the folder size and file size information.'
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
        setSelectedOption(option.title);
        setIsOpen(false);
        Cookies.set("mode", option.title);
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
        if (Cookies.get("mode")) {
            setSelectedOption(Cookies.get("mode") ?? "Optimized mode")
        } else {
            setSelectedOption(options[0]?.title ?? "Optimized mode")
            Cookies.set("mode", selectedOption)
        }
    }, [])

    return (
        <div className="relative inline-block text-left pl-5" ref={dropdownRef}>
            <button
                onClick={toggleDropdown}
                className="inline-flex justify-between items-center w-48 px-3 py-2 text-sm font-medium text-white bg-gray-700 rounded-md hover:bg-gray-600 focus:outline-none"
            >
                {selectedOption === "Quality mode" ?
                    <IoDiamond size={onButtonSize} /> :
                    <PiSpeedometerBold size={onButtonSize} />
                }
                {selectedOption}
                {isOpen ? <FaArrowDown size={onButtonSize - 5} /> : <FaArrowLeft size={onButtonSize - 5} />}
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
                                {option.mode === "optimized" ? (
                                    <PiSpeedometerBold className="p-2" size={iconSize} />
                                ) : option.mode === "quality" ? (
                                    <IoDiamond className="p-2" size={iconSize} />
                                ) : (
                                    <FaBalanceScale className="p-2" size={iconSize} />
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

export default ChangeMode;
