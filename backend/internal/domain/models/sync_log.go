package models

import (
	"encoding/json"
	"time"

	"github.com/uptrace/bun"
)

type SyncLog struct {
	ID                 int             `bun:"id,pk,autoincrement"`
	SourceConnectionID int             `bun:"source_connection_id"`
	TargetConnectionID int             `bun:"target_connection_id"`
	EventType          string          `bun:"event_type"`
	Status             string          `bun:"status"` // success, error, pending
	SourceData         json.RawMessage `bun:"source_data,type:jsonb"`
	TargetData         json.RawMessage `bun:"target_data,type:jsonb"`
	ErrorMessage       string          `bun:"error_message"`
	CreatedAt          time.Time       `bun:"created_at,default:current_timestamp"`

	bun.BaseModel `bun:"table:sync_logs"`
}
