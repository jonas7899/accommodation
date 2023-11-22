
-- +migrate Up
insert into hotels.place(id, place_name, postalcode) values 
('f2c79b47-f1f8-43ea-86da-e80f2c5b5112', 'Budapest 1. ker.', '1014');

insert into hotels.address(id, place_id, street, street_type, address) values 
('00efd485-5d99-486d-ac92-bab8bcc1f74e', 'f2c79b47-f1f8-43ea-86da-e80f2c5b5112', 'Hess András', 'tér', '1-3');

insert into hotels.hotel_type(id, type_name) values 
('f64ba425-0fe2-447a-9e28-e76ed241ed77', 'apartman'),
('b89b7aef-8993-49a5-a3b4-a9fb1ce771ad', 'pension'),
('4d0d752c-7d6f-4a84-9b02-0b31c1a0d706', 'hotel'),
('d14b1681-5e9b-4305-8539-ae960b6eb840', 'inn'),
('0f0b75a4-bdca-49b2-b9d1-cb1bfe76b290', 'motel'),
('b5e1f029-bd0c-4b50-8c8c-cdbe0fe688e1', 'hostel');

insert into hotels.hotel(id, hotel_name, type_id, stars, wellness, carpark) values 
('4c1bec13-f7e5-4133-b72f-80922044eabc', 'Hilton Budapest', '4d0d752c-7d6f-4a84-9b02-0b31c1a0d706', 5, true, true);

insert into hotels.hotel_address(id, hotel_id, address_id, description) values 
('05d683bb-7f2e-40b9-91cb-c7a4016816d7', '4c1bec13-f7e5-4133-b72f-80922044eabc', '00efd485-5d99-486d-ac92-bab8bcc1f74e', 'Budai vár');

insert into hotels.hotel_room(id, hotel_id, room_name, room_size, air_conditioner, wifi, bathroom, iron) values 
('1b8b6c84-0c71-4339-8b0b-259f0583da63', '4c1bec13-f7e5-4133-b72f-80922044eabc', 'Superior Panoramic Twin Room', 30, true, true, true, true),
('b03a663c-f594-4710-b0f4-63926941cbd5', '4c1bec13-f7e5-4133-b72f-80922044eabc', 'Panoramic Family Double Room', 40, true, true, true, flse),
('4eebd534-6205-4cb9-a4f9-caa243383aaf', '4c1bec13-f7e5-4133-b72f-80922044eabc', 'Superior Panoramic Double Room', 35, true, false, true, true);

insert into hotels.hotel_room_space(id, hotel_room_id, space_name) values 
('d911bf67-af9d-4cdf-899d-ebf694be2f0c', '1b8b6c84-0c71-4339-8b0b-259f0583da63', 'badroom'),
('732023d5-868a-474f-9b50-84dd965b0ddc', '1b8b6c84-0c71-4339-8b0b-259f0583da63', 'bathroom'),
('efa9b867-3031-47ce-b1a7-ce4a6535cba3', '1b8b6c84-0c71-4339-8b0b-259f0583da63', 'hall'),

('dba7d07c-819a-44d1-8234-7cb0c210eace', 'b03a663c-f594-4710-b0f4-63926941cbd5', 'badroom'),
('8dd6825f-4a25-4ccc-a08a-e6b05f1746e1', 'b03a663c-f594-4710-b0f4-63926941cbd5', 'bathroom'),
('8ed9a47b-4dcd-4bd6-9d34-546e77219b55', 'b03a663c-f594-4710-b0f4-63926941cbd5', 'hall'),

('eb623109-e523-4fe9-909e-e04e815dee5b', '4eebd534-6205-4cb9-a4f9-caa243383aaf', 'badroom'),
('e75330f8-f53c-4fcf-86b7-c7286bd57996', '4eebd534-6205-4cb9-a4f9-caa243383aaf', 'bathroom'),
('4af762b8-d870-4134-95f6-e65ab1c2b0b8', '4eebd534-6205-4cb9-a4f9-caa243383aaf', 'hall');

insert into hotels.bad_type(id, type_name, sleeps) values 
('d971fb2e-7b55-4674-9e00-aaa4f9c0b4fc', 'extra large bed', 2),
('1b8eedc0-f1e2-4108-93ad-c4d91b8814bb', 'double', 2),
('78ebeafc-ebf6-4d43-90eb-637aeb793e24', 'single', 1),
('be95d198-1f5d-4ed0-aa11-ba694be5b0c5', 'extra bed', 1);

insert into hotels.hotel_room_bad(id, hotel_room_space_id, bad_num, bad_type_id) values 
('5a52abfd-48ee-435b-97c0-8c180d868d67', 'd911bf67-af9d-4cdf-899d-ebf694be2f0c', 1, 'd971fb2e-7b55-4674-9e00-aaa4f9c0b4fc'),
('8461df0a-d3f8-4ba6-9a2f-9b86957bd52f', 'dba7d07c-819a-44d1-8234-7cb0c210eace', 2, '78ebeafc-ebf6-4d43-90eb-637aeb793e24'),
('9f6a4831-fec6-4586-8839-f95c58592e85', 'eb623109-e523-4fe9-909e-e04e815dee5b', 1, '1b8eedc0-f1e2-4108-93ad-c4d91b8814bb'),
('da770cdf-c43c-439c-a260-110ee719ae6d', 'eb623109-e523-4fe9-909e-e04e815dee5b', 1, 'be95d198-1f5d-4ed0-aa11-ba694be5b0c5');

INSERT INTO hotels.reservation (id,hotel_room_id,reservation_begin,reservation_end) VALUES
('42f0d637-658a-47b5-954d-be6369d60b7f','1b8b6c84-0c71-4339-8b0b-259f0583da63','2023-12-01','2023-12-12'),
('029cdc4f-0893-4af3-9278-bf45e8091359','1b8b6c84-0c71-4339-8b0b-259f0583da63','2023-12-16','2023-12-19');

-- +migrate Down
delete from hotels.reservation;
delete from hotels.hotel_room_bad;
delete from hotels.bad_type;
delete from hotels.hotel_room_space;
delete from hotels.hotel_room;
delete from hotels.hotel_address;
delete from hotels.hotel;
delete from hotels.hotel_type;
delete from hotels.address;
delete from hotels.place;