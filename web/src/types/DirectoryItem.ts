export interface DirectoryItem {
    id: number,
    name: string,
    parentPath: string,
    path: string,
    isDirectory: boolean,
    dateModified: string,
    size: string,
    sizeBytes: number,
    storage_limit: string
}