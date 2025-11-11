import type { AccountPermissions } from "./AccountPermissions";

export interface Account {
    id?: number,
    username: string,
    email: string,
    password: string,
    permissions: AccountPermissions
}