export interface AccountPermissions {
    read_directories: boolean,
    read_files: boolean,
    create: boolean,
    change: boolean,
    delete: boolean,
    move: boolean,
    download_files: boolean,
    upload_files: boolean,
    rename: boolean,
    extract: boolean,
    copy: boolean,
    read_recovery: boolean,
    use_recovery: boolean,
    read_users: boolean,
    edit_users: boolean,
    read_logs: boolean
}