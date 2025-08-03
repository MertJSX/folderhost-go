package types

type Account struct {
	Name        string               `yaml:"name"`
	Password    string               `yaml:"password"`
	Permissions []AccountPermissions `yaml:"permissions"`
}

type AccountPermissions struct {
	ReadDirectories bool `yaml:"read_directories"`
	ReadFiles       bool `yaml:"read_files"`
	Create          bool `yaml:"create"`
	Change          bool `yaml:"change"`
	Delete          bool `yaml:"delete"`
	Move            bool `yaml:"move"`
	DownloadFiles   bool `yaml:"download"`
	UploadFiles     bool `yaml:"upload"`
	Rename          bool `yaml:"rename"`
	Unzip           bool `yaml:"unzip"`
	Copy            bool `yaml:"copy"`
}
