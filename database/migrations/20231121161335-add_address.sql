
-- +migrate Up
CREATE TABLE hotels.address (
	id uuid NOT NULL DEFAULT uuid_generate_v4(),
	place_id uuid not null,
	street varchar(50) NOT null,
	street_type varchar(50) NOT null,
	address varchar(50) NOT null,
	CONSTRAINT pk_address_id PRIMARY KEY (id),
	CONSTRAINT fk_address_place_id FOREIGN KEY (place_id) REFERENCES hotels.place(id)
);

-- +migrate Down
drop table if exists hotels.address;