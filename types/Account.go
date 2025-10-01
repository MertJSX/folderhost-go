package types

type Account struct {
	Username    string             `yaml:"username" json:"username"`
	Email       string             `yaml:"email" json:"email"`
	Password    string             `yaml:"password" json:"password"`
	Permissions AccountPermissions `yaml:"permissions" json:"permissions"`
}

type AccountPermissions struct {
	ReadDirectories bool `yaml:"read_directories" json:"read_directories"`
	ReadFiles       bool `yaml:"read_files" json:"read_files"`
	Create          bool `yaml:"create" json:"create"`
	Change          bool `yaml:"change" json:"change"`
	Delete          bool `yaml:"delete" json:"delete"`
	Move            bool `yaml:"move" json:"move"`
	DownloadFiles   bool `yaml:"download" json:"download_files"`
	UploadFiles     bool `yaml:"upload" json:"upload_files"`
	Rename          bool `yaml:"rename" json:"rename"`
	Extract         bool `yaml:"extract" json:"extract"`
	Copy            bool `yaml:"copy" json:"copy"`
	ReadRecovery    bool `yaml:"read_recovery" json:"read_recovery"`
	UseRecovery     bool `yaml:"use_recovery" json:"use_recovery"`
	ReadUsers       bool `yaml:"read_users" json:"read_users"`
	EditUsers       bool `yaml:"edit_users" json:"edit_users"`
	ReadLogs        bool `yaml:"read_logs" json:"read_logs"`
}
