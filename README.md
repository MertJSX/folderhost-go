<p align="center">
  <img width="500" height="160" alt="image" src="https://github.com/user-attachments/assets/85e0398c-c0f6-4f44-8c67-13d141cf1335" />
</p>
<p align="center">
  <img src="https://img.shields.io/badge/language-Go-blue?style=flat-square" alt="Go Language">
  <a href="https://github.com/kullaniciAdi/projeAdi/releases">
    <img src="https://img.shields.io/github/v/release/MertJSX/folderhost?style=flat-square" alt="Latest Release">
  </a>
  <a href="LICENSE">
    <img alt="License" src="https://img.shields.io/github/license/MertJSX/folderhost">
  </a>
  <a href="https://github.com/MertJSX/folderhost/releases">
    <img alt="Downloads" src="https://img.shields.io/github/downloads/MertJSX/folderhost/total?style=flat-square">
  </a>
</p>

<p align="center">
  <strong>Self-hosted cloud platform in a single binary</strong> - Share files, collaborate in real-time, and manage users with zero dependencies.
</p>


> 🚀 **No dependency required** - Download and run on Windows or Linux

---

## 🖥️ Screenshots

<img width="700px" alt="image" src="https://github.com/user-attachments/assets/cc85a58c-3a25-4341-86d8-7dfac9bca41d" />

<details>
  <summary>More Screenshots</summary>
    <img width="700px" alt="image" src="https://github.com/user-attachments/assets/04286979-6bd9-4c02-92a1-b994242fc576" />
    <img width="700px" alt="image" src="https://github.com/user-attachments/assets/ca462852-6fae-4ec2-9479-af7f3e3a09a1" />
</details>


---

## 📥 Quick Start

> ⚠️ **Beta Notice:** FolderHost is currently in **beta**. Some features may be incomplete or subject to change.
 
<br>

[![Download Latest Release](https://img.shields.io/github/v/release/MertJSX/folderhost?style=for-the-badge&logo=github&label=Download&color=2ea44f)](https://github.com/MertJSX/folderhost/releases/latest)

### 🪟 Windows
```powershell
# Download the .exe, then:
folderhost.exe
```

### 🐧 Linux
> Just copy and paste this and folderhost will start working. It's around 23 mb for linux.
```bash
# Download and extract
wget https://github.com/MertJSX/folderhost/releases/latest/download/folderhost-linux-amd64.zip
unzip folderhost-linux-amd64.zip
chmod +x folderhost

# Run
./folderhost
```

---

## 📊 Why FolderHost?
**Lightweight alternative to Nextcloud** - Get the same file sharing features without the complexity.

| Feature | FolderHost | Nextcloud |
|---------|-----------|-----------|
| **Binary Size** | 23MB | 200MB+ |
| **Dependencies** | None | PHP+Database |
| **Setup Time** | 30 seconds | 10+ minutes |
| Single Binary | ✅ | ❌ |
| Zero Dependencies | ✅ | ❌ |
| Real-time Editing | ✅ | ❌ |
| Easy Setup | ✅ | ❌ |

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

## ⚙️ Configuration
> If you encounter any issues or have questions, please feel free to open an [issue](https://github.com/MertJSX/folderhost/issues) or provide feedback. Your input is highly appreciated!

On first run, a `config.yml` file will be created. Edit it to customize:

⚠️ **Important:** Change `secret_jwt_key` and admin password before starting to use!

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
# Here is the GitHub page of Folderhost: https://github.com/MertJSX/folderhost
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
  password: "123"
  email: "example@email.com"
  scope: "" # for example "/yourfolder", this attribute will set a specific location for user and user can't escape it
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
</details>

**🎯 Default Access**

Once running, open your browser to:
```
http://localhost:5000
```

Default credentials: `admin` / `123` (⚠️ **Change immediately!**)

[📖 View All Releases](https://github.com/MertJSX/folderhost/releases) • [🐛 Report Issues](https://github.com/MertJSX/folderhost/issues)


## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

---

## 📝 License

[MIT License](LICENSE)

---

## 🙏 Credits

Built with ❤️ by [MertJSX](https://github.com/MertJSX)

**Tech Stack:**
- Backend: Go
- Frontend: React + TypeScript + Vite
- Editor: Monaco Editor
- Database: SQLite
