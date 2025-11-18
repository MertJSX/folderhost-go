import { Link, useParams } from "react-router-dom"
import { FaUserFriends, FaUserEdit } from "react-icons/fa"
import { useEffect, useState } from "react"
import { type Account } from "../../types/Account"
import PermissionToggle from "../../components/minimal/PermissionToggle/PermissionToggle"
import { type AccountPermissions } from "../../types/AccountPermissions"
import axiosInstance from "../../utils/axiosInstance"
import MessageBox from "../../components/minimal/MessageBox/MessageBox"
import { useNavigate } from 'react-router-dom';
import { useCallback } from "react"

const EditUser = () => {
    const params = useParams();
    const navigate = useNavigate();
    const username = params.username;
    const [error, setError] = useState<string>("");
    const [memUser, setMemUser] = useState<Account | null>(null)
    const [user, setUser] = useState<Account>({
        id: 1,
        username: "",
        email: "",
        scope: "",
        password: "",
        permissions: {
            read_directories: false,
            read_files: false,
            create: false,
            change: false,
            delete: false,
            move: false,
            download_files: false,
            upload_files: false,
            rename: false,
            extract: false,
            copy: false,
            read_recovery: false,
            use_recovery: false,
            read_users: false,
            edit_users: false,
            read_logs: false
        }
    })

    useEffect(() => {
        document.title = "Edit User - folderhost"
        getUserData()
    }, [])

    const getUserData = useCallback(() => {
        axiosInstance.get(`/users/${username}`).then((data) => {
            setUser(data?.data?.user)
            setMemUser(data?.data?.user)
        }).catch((err) => {
            if (err?.response?.data?.err) {
                setError(err.response.data.err)
            }
        })
    }, [])

    const handleInputChange = useCallback((field: keyof Omit<Account, 'permissions'>, value: string) => {
        setUser(prev => ({
            ...prev,
            [field]: value
        }))
    }, [])

    const handlePermissionChange = useCallback((permission: keyof AccountPermissions, value: boolean) => {
        setUser(prev => ({
            ...prev,
            permissions: {
                ...prev.permissions,
                [permission]: value
            }
        }))
    }, [])

    const handleSubmit = useCallback(() => {
        axiosInstance.put(`/users/edit`, {
            user: user
        }).then(() => {
            navigate("/users")
        }).catch((err) => {
            if (err?.response?.data?.err) {
                setError(err.response.data.err)
            }
        })
    }, [user])

    const handleRemove = useCallback(() => {
        if (!window.confirm("If you remove a user, their logs will be deleted too. Are you sure you want to remove this user? This action cannot be undone.")) {
            return;
        }
        axiosInstance.delete(`/users/remove/${user.id}`).then(() => {
            navigate("/users")
        }).catch((err) => {
            if (err?.response?.data?.err) {
                setError(err.response.data.err)
            }
        })
    }, [user])

    return (
        <div>
            <MessageBox message={error !== "" ? error : ""} isErr={error !== ""} setMessage={setError} />
            <section className="my-10 flex flex-col bg-gray-800 gap-6 md:w-4/5 mx-auto p-6 max-w-[1000px] min-h-[600px] shadow-2xl rounded-lg">
                <div className="flex justify-between items-center">
                    <h1 className="flex text-3xl items-center gap-3 text-white font-semibold">
                        <FaUserEdit className="text-blue-400" /> Edit: {username}
                    </h1>
                    <Link
                        to="/users"
                        className="bg-sky-600 hover:bg-sky-500 text-white px-4 py-2 flex items-center gap-2 rounded-lg transition-colors duration-200 font-medium"
                    >
                        <FaUserFriends size={18} /> Users
                    </Link>
                </div>

                <hr className="border-gray-600" />

                <section className="flex flex-col gap-6 overflow-y-auto flex-1 p-2">
                    {/* Basic Information */}
                    <div className="space-y-4">
                        <h2 className="text-xl text-white font-semibold">Basic Information</h2>
                        <div className="grid grid-cols-1 gap-4">
                            <input
                                type="text"
                                placeholder="Username (Required)"
                                value={user.username}
                                onChange={(e) => handleInputChange('username', e.target.value)}
                                className="bg-gray-700 text-white px-4 py-3 rounded-lg border border-gray-500 focus:border-blue-400 focus:outline-none transition-colors placeholder-gray-400"
                            />
                            <input
                                type="email"
                                placeholder="Email"
                                value={user.email}
                                onChange={(e) => handleInputChange('email', e.target.value)}
                                className="bg-gray-700 text-white px-4 py-3 rounded-lg border border-gray-500 focus:border-blue-400 focus:outline-none transition-colors placeholder-gray-400"
                            />
                            <input
                                type="text"
                                placeholder="Scope ( Ex: /foldername | Not required )"
                                value={user.scope}
                                onChange={(e) => handleInputChange('scope', e.target.value)}
                                className="bg-gray-700 text-white px-4 py-3 rounded-lg border border-gray-500 focus:border-blue-400 focus:outline-none transition-colors placeholder-gray-400"
                            />
                        </div>
                    </div>

                    {/* Permissions */}
                    <div className="space-y-4">
                        <h2 className="text-xl text-white font-semibold">Permissions</h2>
                        <div className="grid grid-cols-1 md:grid-cols-2 gap-4 bg-gray-700 p-4 rounded-lg">
                            {/* File Operations */}
                            <div className="space-y-3">
                                <h3 className="text-lg text-gray-300 font-medium">File Operations</h3>
                                <PermissionToggle
                                    label="Read Files"
                                    checked={user.permissions.read_files}
                                    onChange={(checked) => handlePermissionChange('read_files', checked)}
                                />
                                <PermissionToggle
                                    label="Create Items"
                                    checked={user.permissions.create}
                                    onChange={(checked) => handlePermissionChange('create', checked)}
                                />
                                <PermissionToggle
                                    label="Change Files"
                                    checked={user.permissions.change}
                                    onChange={(checked) => handlePermissionChange('change', checked)}
                                />
                                <PermissionToggle
                                    label="Delete Items"
                                    checked={user.permissions.delete}
                                    onChange={(checked) => handlePermissionChange('delete', checked)}
                                />
                                <PermissionToggle
                                    label="Upload Files"
                                    checked={user.permissions.upload_files}
                                    onChange={(checked) => handlePermissionChange('upload_files', checked)}
                                />
                                <PermissionToggle
                                    label="Download Files"
                                    checked={user.permissions.download_files}
                                    onChange={(checked) => handlePermissionChange('download_files', checked)}
                                />
                            </div>

                            {/* Advanced Operations */}
                            <div className="space-y-3">
                                <h3 className="text-lg text-gray-300 font-medium">Advanced Operations</h3>
                                <PermissionToggle
                                    label="Move Items"
                                    checked={user.permissions.move}
                                    onChange={(checked) => handlePermissionChange('move', checked)}
                                />
                                <PermissionToggle
                                    label="Rename Items"
                                    checked={user.permissions.rename}
                                    onChange={(checked) => handlePermissionChange('rename', checked)}
                                />
                                <PermissionToggle
                                    label="Copy Items"
                                    checked={user.permissions.copy}
                                    onChange={(checked) => handlePermissionChange('copy', checked)}
                                />
                                <PermissionToggle
                                    label="Extract Archives"
                                    checked={user.permissions.extract}
                                    onChange={(checked) => handlePermissionChange('extract', checked)}
                                />
                                <PermissionToggle
                                    label="Read Directories"
                                    checked={user.permissions.read_directories}
                                    onChange={(checked) => handlePermissionChange('read_directories', checked)}
                                />
                            </div>

                            {/* System Operations */}
                            <div className="space-y-3">
                                <h3 className="text-lg text-gray-300 font-medium">System Operations</h3>
                                <PermissionToggle
                                    label="Read Recovery"
                                    checked={user.permissions.read_recovery}
                                    onChange={(checked) => handlePermissionChange('read_recovery', checked)}
                                />
                                <PermissionToggle
                                    label="Use Recovery"
                                    checked={user.permissions.use_recovery}
                                    onChange={(checked) => handlePermissionChange('use_recovery', checked)}
                                />
                                <PermissionToggle
                                    label="Read Users"
                                    checked={user.permissions.read_users}
                                    onChange={(checked) => handlePermissionChange('read_users', checked)}
                                />
                                <PermissionToggle
                                    label="Edit Users"
                                    checked={user.permissions.edit_users}
                                    onChange={(checked) => handlePermissionChange('edit_users', checked)}
                                />
                                <PermissionToggle
                                    label="Read Logs"
                                    checked={user.permissions.read_logs}
                                    onChange={(checked) => handlePermissionChange('read_logs', checked)}
                                />
                            </div>
                        </div>
                    </div>
                </section>

                {/* Submit Button */}
                <div className="flex justify-center gap-2 pt-4 border-t border-gray-600">
                    <button
                        onClick={handleSubmit}
                        disabled={(!user.username || !user.password) || (memUser == user)}
                        className="bg-green-700 hover:bg-green-600 disabled:bg-gray-500 disabled:cursor-not-allowed text-white px-6 py-3 rounded-lg font-semibold transition-colors duration-200 w-2/3"
                    >
                        Submit Changes
                    </button>
                    <button
                        title="Double click to remove"
                        onClick={handleRemove}
                        className="bg-red-500 hover:bg-red-700 disabled:bg-gray-500 disabled:cursor-not-allowed text-white px-6 py-3 rounded-lg font-semibold transition-colors duration-200 w-1/3"
                    >
                        Remove Account
                    </button>
                </div>
            </section>
        </div>
    )
}

export default EditUser