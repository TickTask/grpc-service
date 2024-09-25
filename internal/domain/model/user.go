package model

type User struct {
	ID       int64  `db:"id"`
	Login    string `db:"login"`
	Name     string `db:"name"`
	PassHash []byte `db:"hash_password"`
}
