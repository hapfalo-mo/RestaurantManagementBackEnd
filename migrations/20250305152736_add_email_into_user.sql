-- +goose Up
-- +goose StatementBegin
ALTER TABLE `user` ADD COLUMN email VARCHAR(50) NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
