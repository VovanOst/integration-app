package models

import (
	"database/sql"
	"time"

	"github.com/uptrace/bun"
)

type Webhook struct {
	ID           int            `bun:"id,pk,autoincrement"`
	ConnectionID int            `bun:"connection_id"`
	EventType    string         `bun:"event_type"`
	CallbackURL  string         `bun:"callback_url"`
	SecretKey    sql.NullString `bun:"secret_key"`
	IsActive     bool           `bun:"is_active,default:true"`
	CreatedAt    time.Time      `bun:"created_at,default:current_timestamp"`

	bun.BaseModel `bun:"table:webhooks"`
}
