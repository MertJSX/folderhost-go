import type { AccountPermissions } from "./AccountPermissions";

export interface Account {
    id?: number,
    username: string,
    email: string,
    scope: string,
    password: string,
    permissions: AccountPermissions
}