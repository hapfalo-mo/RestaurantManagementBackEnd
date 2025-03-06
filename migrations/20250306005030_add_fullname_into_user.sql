-- +goose Up
-- +goose StatementBegin
ALTER TABLE `user` ADD COLUMN full_name VARCHAR(250) NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
