package model

type User struct {
    ID           int64     `json:"id" db:"id"`
    Callsign     string    `json:"callsign" db:"callsign"`           // уникальный, используется как логин
    PasswordHash string    `json:"-" db:"password_hash"`            // хеш пароля, не выводим в JSON
}
