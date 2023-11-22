
-- +migrate Up
CREATE TABLE hotels.reservation (
	id uuid NOT NULL DEFAULT uuid_generate_v4(),
	hotel_room_id uuid not null,
	reservation_begin date NOT NULL,
	reservation_end date NOT null,
	CONSTRAINT fk_reservation_hotel_room_id FOREIGN KEY (hotel_room_id) REFERENCES hotels.hotel(id)
);

-- +migrate Down
drop table if exists hotels.reservation;