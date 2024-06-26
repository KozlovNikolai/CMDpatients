-- +goose Up
-- +goose StatementBegin
SELECT setval(pg_get_serial_sequence('patients', 'id'), max(id)) FROM patients;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
