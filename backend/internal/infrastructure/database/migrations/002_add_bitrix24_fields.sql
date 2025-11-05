-- +migrate Up
ALTER TABLE connections ADD COLUMN IF NOT EXISTS bitrix24_webhook_secret VARCHAR(255);
ALTER TABLE connections ADD COLUMN IF NOT EXISTS metadata JSONB;

-- +migrate Down
ALTER TABLE connections DROP COLUMN IF EXISTS metadata;
ALTER TABLE connections DROP COLUMN IF EXISTS bitrix24_webhook_secret;