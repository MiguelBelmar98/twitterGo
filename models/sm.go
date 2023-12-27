package models

type Secret struct {
	Host     string `json:"host"`
	Username string `json:"user"`
	Password string `json:"password"`
	JWTSign  string `json:"jwtsign"`
	Database string `json:"databse"`
}
