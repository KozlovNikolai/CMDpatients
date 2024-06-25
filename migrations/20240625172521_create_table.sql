-- +goose Up
-- +goose StatementBegin
CREATE TABLE patients
(
    id serial,
    created_at timestamp(0) default NULL::timestamp without time zone,
    surname    varchar(100) not null,
    name       varchar(100) not null,
    patronymic varchar(100),
    gender     smallint     not null,
    birthday   date         not null,
    PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table patients;
-- +goose StatementEnd
