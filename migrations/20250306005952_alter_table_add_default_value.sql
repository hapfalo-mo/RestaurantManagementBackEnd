-- +goose Up
-- +goose StatementBegin
ALTER TABLE `user` MODIFY COLUMN role VARCHAR(1) NOT NULL DEFAULT 0;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
