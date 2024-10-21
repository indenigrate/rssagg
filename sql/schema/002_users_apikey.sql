-- +goose Up
ALTER TABLE "user" ADD COLUMN api_key VARCHAR(64) UNIQUE NOT NULL DEFAULT (
    encode(sha256(random()::text::bytea), 'hex')
);
-- +goose Down
ALTER TABLE "user" DROP COLUMN api_key;