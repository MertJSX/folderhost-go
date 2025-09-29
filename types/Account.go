package types

type Account struct {
	Username    string             `yaml:"username"`
	Email       string             `yaml:"email"`
	Password    string             `yaml:"password"`
	Permissions AccountPermissions `yaml:"permissions"`
}

type AccountPermissions struct {
	ReadDirectories bool `yaml:"read_directories" json:"read_directories"`
	ReadFiles       bool `yaml:"read_files" json:"read_files"`
	Create          bool `yaml:"create" json:"create"`
	Change          bool `yaml:"change" json:"change"`
	Delete          bool `yaml:"delete" json:"delete"`
	Move            bool `yaml:"move" json:"move"`
	DownloadFiles   bool `yaml:"download" json:"download"`
	UploadFiles     bool `yaml:"upload" json:"upload"`
	Rename          bool `yaml:"rename" json:"rename"`
	Extract         bool `yaml:"extract" json:"unzip"`
	Copy            bool `yaml:"copy" json:"copy"`
	ReadRecovery    bool `yaml:"read_recovery" json:"read_recovery"`
	UseRecovery     bool `yaml:"use_recovery" json:"use_recovery"`
}
