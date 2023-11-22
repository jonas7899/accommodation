
-- +migrate Up
CREATE SCHEMA hotels;

-- +migrate Down
DROP SCHEMA IF EXISTS hotels;
