-- +goose Up
-- +goose StatementBegin
-- Create table User
CREATE TABLE `users`(
    id INT(11) PRIMARY KEY NOT NULL AUTO_INCREMENT,
    phone_number VARCHAR(11) NOT NULL UNIQUE,
    `password` VARCHAR(64) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ,
    updated_at TIMESTAMP  NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,
    role INT(1) NOT NULL, 
    point INT(5) DEFAULT 0 NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
