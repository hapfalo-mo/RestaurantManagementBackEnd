-- +goose Up
-- +goose StatementBegin
    CREATE TABLE booking (
        id INT(11)  NOT NULL PRIMARY KEY AUTO_INCREMENT ,
        user_id INT(11) NOT NULL,
        guest_count INT(20) NOT NULL,
        `time` DATETIME NOT NULL,
        created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
        status INT(11) NOT NULL DEFAULT 0,
        note VARCHAR(255) DEFAULT NULL,
        FOREIGN KEY (user_id) REFERENCES user(id)
    )
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
