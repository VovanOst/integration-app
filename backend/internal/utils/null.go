package utils

import (
	"database/sql"
	"time"
)

// ToNullString конвертирует string в sql.NullString
func ToNullString(s string) sql.NullString {
	return sql.NullString{
		String: s,
		Valid:  s != "",
	}
}

// FromNullString безопасно получает значение из sql.NullString
func FromNullString(ns sql.NullString) string {
	if ns.Valid {
		return ns.String
	}
	return ""
}

// IsNullString проверяет является ли значение null
func IsNullString(ns sql.NullString) bool {
	return !ns.Valid || ns.String == ""
}

// ToNullTime конвертирует time.Time в sql.NullTime
func ToNullTime(t time.Time) sql.NullTime {
	return sql.NullTime{
		Time:  t,
		Valid: !t.IsZero(),
	}
}

// FromNullTime безопасно получает значение из sql.NullTime
func FromNullTime(nt sql.NullTime) time.Time {
	if nt.Valid {
		return nt.Time
	}
	return time.Time{}
}

// IsNullTime проверяет является ли значение null
func IsNullTime(nt sql.NullTime) bool {
	return !nt.Valid || nt.Time.IsZero()
}
