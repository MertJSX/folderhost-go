import React, { type ReactNode } from 'react'

interface ExplorerRMItemProps {
    children?: ReactNode,
    className?: string | undefined,
    title?: string | undefined,
    isDisabled?: boolean,
    onClick?: () => void
}

const ExplorerRMItem: React.FC<ExplorerRMItemProps> = ({
        children,
        className,
        title,
        isDisabled = false,
        onClick
    }) => {
    return (
        <button
            title={title}
            onClick={onClick}
            disabled={isDisabled}
            className={"flex items-center gap-1 w-[80%] p-2 text-left text-base transition-all hover:bg-slate-900 hover:translate-x-1 relative " + className}
        >
            {children}
        </button>
    )
}

export default ExplorerRMItem