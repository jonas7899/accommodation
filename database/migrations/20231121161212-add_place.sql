
-- +migrate Up
CREATE TABLE hotels.place (
	id uuid NOT NULL DEFAULT uuid_generate_v4(),
	place_name varchar(50) NOT NULL,
	postalcode varchar(50) NOT null,
	CONSTRAINT pk_place_id PRIMARY KEY (id)
);

-- +migrate Down
drop table if exists hotels.place;