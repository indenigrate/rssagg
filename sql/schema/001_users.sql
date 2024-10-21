-- +goose Up
CREATE TABLE
    "user" (
        id UUID NOT NULL,
        created_at TIMESTAMP NOT NULL,
        updated_at TIMESTAMP NOT NULL,
        name TEXT NOT NULL
    );

-- +goose Down
DROP TABLE "user";