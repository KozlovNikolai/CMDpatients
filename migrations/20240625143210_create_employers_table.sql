-- +goose Up
-- +goose StatementBegin
CREATE UNIQUE INDEX id_person_id_1719315068821_index ON employers (company) INCLUDE (person_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS id_person_id_1719315068821_index;
-- +goose StatementEnd



