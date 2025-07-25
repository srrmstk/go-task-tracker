-- +goose Up
-- +goose StatementBegin
ALTER TABLE memos
ADD COLUMN category_id UUID NOT NULL;

ALTER TABLE memos
ADD CONSTRAINT fk_memo_category
FOREIGN KEY (category_id)
REFERENCES categories(id)
ON DELETE RESTRICT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE memos
DROP CONSTRAINT IF EXISTS fk_memo_category;

ALTER TABLE memos
DROP COLUMN IF EXISTS category_id;
-- +goose StatementEnd
