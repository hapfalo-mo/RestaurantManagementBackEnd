-- +goose Up
-- +goose StatementBegin
CREATE TABLE user_token (
    id INT(11) NOT NULL  AUTO_INCREMENT,
    user_id int(11) NOT NULL,
    refresh_token VARCHAR(64) NOT NULL,
    PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
