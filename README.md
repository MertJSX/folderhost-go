# 📁 FolderHost

**Self-hosted cloud platform in a single binary** - Share files, collaborate in real-time, and manage users with zero dependencies.

---

## ✨ Features

### 🚀 Core
- **Single Binary Deployment** - No dependencies, just run
- **High Performance** - Built with Go backend + React frontend
- **Real-time Collaboration** - Live code editing with Monaco Editor
- **Multi-user Support** - Permissions system

### 🔧 File Management
- Full file operations (upload, download, move, copy, rename)
- Chunked file uploads for large files
- Recovery bin with configurable limits
- Storage quota management per folder

### 🔒 Security & Monitoring
- JWT-based authentication
- Granular user permissions
- Audit logs for all activities
- Configurable storage limits

---

## 🖥️ Web Panel

### Explorer
<img width="600px" src="https://github.com/user-attachments/assets/9c2825fa-08ac-4eb8-9767-0a0ba3029046" width="700px">

### Collaborative Code Editor
<img width="600px" alt="image" src="https://github.com/user-attachments/assets/04286979-6bd9-4c02-92a1-b994242fc576" />

---

## Default config.yml

<details>
  <summary>Show config</summary>

  ```yml
#      _______   __   __
#     / _____/  / /  / /
#    / /__     / /__/ /
#   / ___/    / ___  /
#  / /       / /  / /
# /_/       /_/  /_/  By MertJSX
#
# Thanks for using my application!!! Please report if you catch any bugs!
# Here is the GitHub page of Folderhost: https://github.com/MertJSX/folderhost-go
#

# Port is required. Don't delete it!
port: 5000

# This is folder path. You can change it, but don't delete.
folder: "./host"

# Limit of the folder. Examples: 10 GB, 300 MB, 5.5 GB, 1 TB...
# You can remove it if you trust users.
storage_limit: "10 GB"

# This is secret json web token key to create tokens.
secret_jwt_key: "you must change it" # Example: 5asdasd1asd

# Admin account properties
admin:
  username: "admin"
  email: "example@email.com"
  password: "123"
  permissions:
    read_directories: true
    read_files: true
    create: true
    change: true
    delete: true
    move: true
    download: true
    upload: true
    rename: true
    extract: true
    copy: true
    read_recovery: true
    use_recovery: true
    read_users: true
    edit_users: true
    read_logs: true

# Holds deleted files. Accidentally, you might delete files that you don't want to delete.
recovery_bin: true

# Optionally you can limit recovery_bin storage. You can remove it if you want.
bin_storage_limit: "5 GB"

# Enable/Disable logging activities
log_activities: true

# Clears logs automatically after some days. If you want to disable it set the value to 0.
clear_logs_after: 7 # Days
```
