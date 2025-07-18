-- +goose Up
-- +goose StatementBegin
ALTER TABlE tasks RENAME TO memos;
ALTER TABLE memos ADD COLUMN score INT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABlE memos DROP COLUMN score;
ALTER TABLE memos RENAME TO tasks;
-- +goose StatementEnd
