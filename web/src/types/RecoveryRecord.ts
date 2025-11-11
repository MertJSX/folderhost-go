export interface RecoveryRecord {
    id: number,
    username: string,
    oldLocation: string,
    binLocation: string,
    isDirectory: boolean,
    sizeDisplay: string,
    sizeBytes: number,
    created_at: string
}