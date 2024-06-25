-- +goose Up
-- +goose StatementBegin
ALTER TABLE persons ADD CONSTRAINT email_unique UNIQUE (email);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE persons DROP email_unique;
-- +goose StatementEnd
