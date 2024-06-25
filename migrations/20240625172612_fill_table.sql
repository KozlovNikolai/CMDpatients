-- +goose Up
-- +goose StatementBegin
INSERT INTO patients (id, created_at, surname, name, patronymic, gender, birthday) VALUES
(1,'2002-04-20 07:37:02.66','Smith', 'Bob', '',1,'1976-04-20'),
(2,'2004-08-31 13:41:54.66','Jones','Dow', '',1,'1980-05-22'),
(3, '2000-11-07 17:01:23.66','Mask','Elon', '',1,'1977-06-24'),
(4, '2002-05-02 10:12:45.66','Dow','John', '',1,'1947-07-26'),
(5, '2002-04-20 09:33:12.66','Wuchacki','Jane', '',0,'2001-08-28'),
(6, '2003-10-20 14:31:12.66','Kidman','Nikol', '',0,'2003-09-30');
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DELETE FROM patients WHERE id in (1, 2, 3, 4,5,6);
-- +goose StatementEnd
