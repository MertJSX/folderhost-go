package types

type RecoveryRecord struct {
	Id          int    `json:"id"`
	Username    string `json:"username"`
	OldLocation string `json:"oldLocation"`
	BinLocation string `json:"binLocation"`
	IsDirectory bool   `json:"isDirectory"`
	SizeDisplay string `json:"sizeDisplay"`
	SizeBytes   int64  `json:"sizeBytes"`
	CreatedAt   string `json:"created_at"`
}
