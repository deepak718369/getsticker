-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE if not exists sticker(
id UUID,
 sticker varchar,
 timing TIMESTAMP WITHOUT TIME ZONE,
 priorities varchar
);


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back