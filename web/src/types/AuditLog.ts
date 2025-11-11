export interface AuditLog {
    id?: number,
    username: string,
    action: string,
    description: string,
    created_at: string
}