package models

import (
	"time"

	"github.com/uptrace/bun"
)

type FieldMapping struct {
	ID                 int       `bun:"id,pk,autoincrement"`
	SourceConnectionID int       `bun:"source_connection_id"`
	TargetConnectionID int       `bun:"target_connection_id"`
	SourceField        string    `bun:"source_field"`
	TargetField        string    `bun:"target_field"`
	CreatedAt          time.Time `bun:"created_at,default:current_timestamp"`

	bun.BaseModel `bun:"table:field_mappings"`
}
