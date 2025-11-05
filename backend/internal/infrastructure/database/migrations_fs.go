package database

import (
	"embed"
)

// ✅ Правильный путь - миграции в папке migrations рядом с этим файлом
//
//go:embed migrations/*.sql
var sqlFS embed.FS
