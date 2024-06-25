-- +goose Up
-- +goose StatementBegin
INSERT INTO persons (id, name, email) VALUES
(1, 'Bob Smith', 'bob@example.com'),
(2, 'Dow Jones', 'dow@example.com'),
(3, 'Elon Mask', 'elon@example.com'),
(4, 'John Dow', 'john@example.com');
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DELETE FROM persons WHERE id in (1, 2, 3, 4);
-- +goose StatementEnd