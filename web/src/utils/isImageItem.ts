import type { DirectoryItem } from "../types/DirectoryItem";

export const isImageItem = (itemInfo: DirectoryItem | null): boolean => {
    if (itemInfo == null) {
        return false
    }
    const imageExtensions = [
        'png', 'jpg', 'jpeg', 'svg', 'webp', 'ico',
        'gif', 'mp', 'tiff', 'tif', 'bmp', 'avif'
    ];

    if (!itemInfo || itemInfo.isDirectory) {
        return false;
    }

    const extension = itemInfo.name?.split('.').pop()?.toLowerCase();
    return extension ? imageExtensions.includes(extension) : false;
};