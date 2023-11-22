
-- +migrate Up
CREATE TABLE hotels.hotel_type (
id uuid NOT NULL DEFAULT uuid_generate_v4(),
	type_name varchar(50) NOT NULL,
	CONSTRAINT pk_hotel_type_id PRIMARY KEY (id)
);

CREATE TABLE hotels.hotel (
	id uuid NOT NULL DEFAULT uuid_generate_v4(),
	hotel_name varchar(50) NOT NULL,
	type_id uuid NOT NULL,
	stars int2  NOT NULL,
	wellness bool NULL,
	carpark bool NULL,
	CONSTRAINT pk_hotel_id PRIMARY KEY (id),
	CONSTRAINT fk_hotel_type_id FOREIGN KEY (type_id) REFERENCES hotels.hotel_type(id)	
);

CREATE TABLE hotels.hotel_address (
	id uuid NOT NULL DEFAULT uuid_generate_v4(),
	hotel_id uuid NOT null,
	address_id uuid NOT null,
	description varchar(50) NOT NULL,
	CONSTRAINT pk_hotel_address_id PRIMARY KEY (id),
	CONSTRAINT fk_hotel_address_hotel_id FOREIGN KEY (hotel_id) REFERENCES hotels.hotel(id),
	CONSTRAINT fk_hotel_address_address_id FOREIGN KEY (address_id) REFERENCES hotels.address(id)
);

CREATE TABLE hotels.hotel_room (
	id uuid NOT NULL DEFAULT uuid_generate_v4(),
	hotel_id uuid NOT NULL,
	room_name varchar(50) NOT NULL,
	room_size int2,
 	air_conditioner bool,
	wifi bool,
	bathroom bool,
	iron bool,
	CONSTRAINT pk_hotel_room_id PRIMARY KEY (id),
	CONSTRAINT fk_hotel_room_hotel_id FOREIGN KEY (hotel_id) REFERENCES hotels.hotel(id)
);

CREATE TABLE hotels.hotel_room_space (
	id uuid NOT NULL DEFAULT uuid_generate_v4(),
	hotel_room_id uuid NOT NULL,
	space_name varchar(50) NOT NULL,
	CONSTRAINT pk_hotel_room_space_id PRIMARY KEY (id),
	CONSTRAINT fk_hotel_room_space_hotel_room_id FOREIGN KEY (hotel_room_id) REFERENCES hotels.hotel_room(id)
);

CREATE TABLE hotels.bad_type (
	id uuid NOT NULL DEFAULT uuid_generate_v4(),
	type_name varchar(50) NOT NULL,
	sleeps int2 not null,
	CONSTRAINT pk_bad_type_id PRIMARY KEY (id)
);

CREATE TABLE hotels.hotel_room_bad (
	id uuid NOT NULL DEFAULT uuid_generate_v4(),
	hotel_room_space_id uuid NOT NULL,
	bad_num int2,
 	bad_type_id uuid NOT NULL,
	CONSTRAINT pk_hotel_room_bad_id PRIMARY KEY (id),
	CONSTRAINT fk_hotel_room_bad_hotel_room_space_id FOREIGN KEY (hotel_room_space_id) REFERENCES hotels.hotel_room_space(id),
	CONSTRAINT fk_hotel_room_bad_ FOREIGN KEY (bad_type_id) REFERENCES hotels.bad_type(id)
);

-- +migrate Down
drop table if exists hotels.hotel_room_bad;
drop table if exists hotels.bad_type;
drop table if exists hotels.hotel_room_space;
drop table if exists hotels.hotel_room;
drop table if exists hotels.hotel_address;
drop table if exists hotels.hotel;
drop table if exists hotels.hotel_type;