-- +goose Up
-- +goose StatementBegin
UPDATE memos
SET score = 0
WHERE score IS NULL;

ALTER TABLE memos
ALTER COLUMN score SET NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE memos
ALTER COLUMN score DROP NOT NULL;
-- +goose StatementEnd
