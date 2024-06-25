-- +goose Up
-- +goose StatementBegin
CREATE TABLE persons (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL
);
CREATE TABLE employers (
    id SERIAL PRIMARY KEY,
    company VARCHAR(100) NOT NULL,
    person_id INT REFERENCES persons(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table employers, persons;
-- +goose StatementEnd
