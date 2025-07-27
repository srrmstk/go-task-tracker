-- +goose Up
-- +goose StatementBegin
ALTER TABLE users
ADD COLUMN email VARCHAR(255) NOT NULL UNIQUE;
ALTER TABLE users
ADD COLUMN confirmed BOOLEAN NOT NULL DEFAULT false;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users
DROP COLUMN IF EXISTS email;
ALTER TABLE users
DROP COLUMN IF EXISTS confirmed;
-- +goose StatementEnd
