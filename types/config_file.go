package types

type ConfigFile struct {
	Port                 int     `yaml:"port"`
	Folder               string  `yaml:"folder"`
	StorageLimit         string  `yaml:"storage_limit"`
	SecretJwtKey         string  `yaml:"secret_jwt_key"`
	BirthDate            string  `yaml:"birthDate"`
	DateModified         string  `yaml:"dateModified"`
	Size                 string  `yaml:"size"`
	AdminAccount         Account `yaml:"admin"`
	RecoveryBin          bool    `yaml:"recovery_bin"`
	BinStorageLimit      string  `yaml:"bin_storage_limit"`
	LogActivities        bool    `yaml:"log_activities"`
	GetFoldersizeOnStart bool    `yaml:"get_foldersize_on_start"`
	ClearLogsAfter       int     `yaml:"clear_logs_after"`
}
