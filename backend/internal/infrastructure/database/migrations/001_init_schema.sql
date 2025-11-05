-- +migrate Up
CREATE TABLE IF NOT EXISTS connections (
    id SERIAL PRIMARY KEY,
    system_type VARCHAR(50) NOT NULL,
    name VARCHAR(255) NOT NULL,
    access_token TEXT NOT NULL,
    refresh_token TEXT,
    expires_at TIMESTAMP,
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );

CREATE TABLE IF NOT EXISTS field_mappings (
    id SERIAL PRIMARY KEY,
    source_connection_id INT REFERENCES connections(id) ON DELETE CASCADE,
    target_connection_id INT REFERENCES connections(id) ON DELETE CASCADE,
    source_field VARCHAR(255) NOT NULL,
    target_field VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(source_connection_id, target_connection_id, source_field)
    );

CREATE TABLE IF NOT EXISTS webhooks (
    id SERIAL PRIMARY KEY,
    connection_id INT REFERENCES connections(id) ON DELETE CASCADE,
    event_type VARCHAR(100) NOT NULL,
    callback_url TEXT NOT NULL,
    secret_key VARCHAR(255),
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );

CREATE TABLE IF NOT EXISTS sync_logs (
    id SERIAL PRIMARY KEY,
    source_connection_id INT REFERENCES connections(id) ON DELETE CASCADE,
    target_connection_id INT REFERENCES connections(id) ON DELETE CASCADE,
    event_type VARCHAR(100),
    status VARCHAR(50) NOT NULL,
    source_data JSONB,
    target_data JSONB,
    error_message TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );

CREATE INDEX idx_connections_is_active ON connections(is_active);
CREATE INDEX idx_mappings_source ON field_mappings(source_connection_id);
CREATE INDEX idx_sync_logs_created_at ON sync_logs(created_at DESC);
CREATE INDEX idx_sync_logs_status ON sync_logs(status);
CREATE INDEX idx_webhooks_connection_id ON webhooks(connection_id);

-- +migrate Down
DROP INDEX IF EXISTS idx_webhooks_connection_id;
DROP INDEX IF EXISTS idx_sync_logs_status;
DROP INDEX IF EXISTS idx_sync_logs_created_at;
DROP INDEX IF EXISTS idx_mappings_source;
DROP INDEX IF EXISTS idx_connections_is_active;
DROP TABLE IF EXISTS sync_logs;
DROP TABLE IF EXISTS webhooks;
DROP TABLE IF EXISTS field_mappings;
DROP TABLE IF EXISTS connections;
