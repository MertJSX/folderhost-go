package types

type AuditLog struct {
	ID          *int   `yaml:"id,omitempty" json:"id,omitempty"`
	Username    string `yaml:"username" json:"username"`
	Action      string `yaml:"action" json:"action"`
	Description string `yaml:"description" json:"description"`
	CreatedAt   string `yaml:"created_at" json:"created_at"`
}
