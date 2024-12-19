-- +goose Up
-- +goose StatementBegin
CREATE TABLE users
(
    id       INTEGER PRIMARY KEY AUTOINCREMENT,
    login    TEXT NOT NULL,
    email    TEXT NOT NULL,
    password TEXT NOT NULL,

    UNIQUE (login, email)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
