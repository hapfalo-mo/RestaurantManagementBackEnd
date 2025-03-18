-- +goose Up
-- +goose StatementBegin
ALTER TABLE booking
ADD COLUMN description VARCHAR(200),
ADD COLUMN handler_booking INT(11);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE booking
DROP COLUMN description,
DROP COLUMN handler_booking;
-- +goose StatementEnd
