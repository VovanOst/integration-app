CREATE TABLE IF NOT EXISTS mappings (
                                        id SERIAL PRIMARY KEY,
                                        bitrix_field VARCHAR(255) NOT NULL,
    facebook_field VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );

CREATE INDEX idx_bitrix_field ON mappings(bitrix_field);