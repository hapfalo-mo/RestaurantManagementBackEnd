-- +goose Up
-- +goose StatementBegin
ALTER TABLE BOOKING 
ADD COLUMN `customer_name` VARCHAR(100) NOT NULL,
ADD COLUMN `customer_phone` VARCHAR(15) NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE BOOKING 
DROP COLUMN `customer_name`,
DROP COLUMN `customer_phone`;
-- +goose StatementEnd
