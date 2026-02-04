-- +goose Up
-- +goose StatementBegin
CREATE TABLE runs (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    repo TEXT,
    commit_sha TEXT,
    status TEXT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE runs;
-- +goose StatementEnd
