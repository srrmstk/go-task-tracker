-- +goose Up
-- +goose StatementBegin
ALTER TABLE categories
ADD CONSTRAINT unique_title UNIQUE (title);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE categories
DROP CONSTRAINT unique_title;
-- +goose StatementEnd
