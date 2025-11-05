package models

import (
	"database/sql"
	"time"

	"github.com/uptrace/bun"
)

type Connection struct {
	ID           int            `bun:"id,pk,autoincrement"`
	SystemType   string         `bun:"system_type"` // bitrix24, facebook, etc
	Name         string         `bun:"name"`
	AccessToken  string         `bun:"access_token"`
	RefreshToken sql.NullString `bun:"refresh_token"`
	ExpiresAt    sql.NullTime   `bun:"expires_at"`
	IsActive     bool           `bun:"is_active,default:true"`
	CreatedAt    time.Time      `bun:"created_at,default:current_timestamp"`
	UpdatedAt    time.Time      `bun:"updated_at,default:current_timestamp"`

	bun.BaseModel `bun:"table:connections"`
}
