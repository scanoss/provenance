DROP TABLE if exists vendor_locations;

CREATE TABLE vendor_locations (
	vendor_id int4 NOT NULL,
	indexed_date date NOT NULL,
	declared_location text NOT NULL,
	parsed_location text NULL,
	curated_countries_ids text NULL,
	curated_date date NULL,
	curated_status_id int4 NULL,
	curator_id int4 NULL,
	id serial4 NOT NULL,
	CONSTRAINT vendor_locations_pkey PRIMARY KEY (id),
	CONSTRAINT vendor_locations_unique UNIQUE (vendor_id, declared_location)
);
INSERT INTO vendor_locations
(vendor_id, indexed_date, declared_location, parsed_location, curated_countries_ids, curated_date, curated_status_id, curator_id, id)
VALUES(283816, '2024-04-09', 'Spain', 'Spain', '{163}', NULL, NULL, NULL, 1460819);
INSERT INTO vendor_locations
(vendor_id, indexed_date, declared_location, parsed_location, curated_countries_ids, curated_date, curated_status_id, curator_id, id)
VALUES(288334, '2024-04-09', 'Germany', 'Germany', NULL, NULL, NULL, NULL, 1563326);
INSERT INTO vendor_locations
(vendor_id, indexed_date, declared_location, parsed_location, curated_countries_ids, curated_date, curated_status_id, curator_id, id)
VALUES(351225, '2024-04-09', 'Portland, OR', 'United States', NULL, NULL, NULL, NULL, 855571);
INSERT INTO vendor_locations
(vendor_id, indexed_date, declared_location, parsed_location, curated_countries_ids, curated_date, curated_status_id, curator_id, id)
VALUES(376679, '2024-04-09', 'San Francisco, CA', 'United States', NULL, NULL, NULL, NULL, 2657104);
INSERT INTO vendor_locations
(vendor_id, indexed_date, declared_location, parsed_location, curated_countries_ids, curated_date, curated_status_id, curator_id, id)
VALUES(501929, '2024-04-09', 'Finland', 'Finland', NULL, NULL, NULL, NULL, 858374);
INSERT INTO vendor_locations
(vendor_id, indexed_date, declared_location, parsed_location, curated_countries_ids, curated_date, curated_status_id, curator_id, id)
VALUES(513134, '2024-04-09', 'Shenzhen, China', 'China', NULL, NULL, NULL, NULL, 1969637);
INSERT INTO vendor_locations
(vendor_id, indexed_date, declared_location, parsed_location, curated_countries_ids, curated_date, curated_status_id, curator_id, id)
VALUES(529783, '2024-04-09', 'Stockholm', 'Sweden', NULL, NULL, NULL, NULL, 1927743);
INSERT INTO vendor_locations
(vendor_id, indexed_date, declared_location, parsed_location, curated_countries_ids, curated_date, curated_status_id, curator_id, id)
VALUES(721581, '2024-04-09', '', '', NULL, NULL, NULL, NULL, 1788060);
INSERT INTO vendor_locations
(vendor_id, indexed_date, declared_location, parsed_location, curated_countries_ids, curated_date, curated_status_id, curator_id, id)
VALUES(793636, '2024-04-09', 'Tokyo', 'Japan', NULL, NULL, NULL, NULL, 1800952);
INSERT INTO vendor_locations
(vendor_id, indexed_date, declared_location, parsed_location, curated_countries_ids, curated_date, curated_status_id, curator_id, id)
VALUES(938258, '2024-04-09', '', '', NULL, NULL, NULL, NULL, 1644861);
INSERT INTO vendor_locations
(vendor_id, indexed_date, declared_location, parsed_location, curated_countries_ids, curated_date, curated_status_id, curator_id, id)
VALUES(1197052, '2024-04-09', 'Germany', 'Germany', NULL, NULL, NULL, NULL, 2046230);
INSERT INTO vendor_locations
(vendor_id, indexed_date, declared_location, parsed_location, curated_countries_ids, curated_date, curated_status_id, curator_id, id)
VALUES(1197561, '2024-04-09', 'Valparaíso, Chile', 'Chile', NULL, NULL, NULL, NULL, 2204002);
INSERT INTO vendor_locations
(vendor_id, indexed_date, declared_location, parsed_location, curated_countries_ids, curated_date, curated_status_id, curator_id, id)
VALUES(1201092, '2024-04-09', '', '', NULL, NULL, NULL, NULL, 2000149);
INSERT INTO vendor_locations
(vendor_id, indexed_date, declared_location, parsed_location, curated_countries_ids, curated_date, curated_status_id, curator_id, id)
VALUES(1202134, '2024-04-09', '', '', NULL, NULL, NULL, NULL, 2001469);
INSERT INTO vendor_locations
(vendor_id, indexed_date, declared_location, parsed_location, curated_countries_ids, curated_date, curated_status_id, curator_id, id)
VALUES(1320559, '2024-04-09', 'Shenzhen, China', 'China', NULL, NULL, NULL, NULL, 2726982);
INSERT INTO vendor_locations
(vendor_id, indexed_date, declared_location, parsed_location, curated_countries_ids, curated_date, curated_status_id, curator_id, id)
VALUES(1405397, '2024-04-09', '', '', NULL, NULL, NULL, NULL, 2801157);
INSERT INTO vendor_locations
(vendor_id, indexed_date, declared_location, parsed_location, curated_countries_ids, curated_date, curated_status_id, curator_id, id)
VALUES(1405399, '2024-04-09', '', '', '{{7}}', NULL, 1, NULL, 2801163);
INSERT INTO vendor_locations
(vendor_id, indexed_date, declared_location, parsed_location, curated_countries_ids, curated_date, curated_status_id, curator_id, id)
VALUES(1405400, '2024-04-09', '', '', NULL, NULL, NULL, NULL, 2801161);
INSERT INTO vendor_locations
(vendor_id, indexed_date, declared_location, parsed_location, curated_countries_ids, curated_date, curated_status_id, curator_id, id)
VALUES(1405464, '2024-04-09', '', '', '{{163}}', NULL, NULL, NULL, 2801268);
INSERT INTO vendor_locations
(vendor_id, indexed_date, declared_location, parsed_location, curated_countries_ids, curated_date, curated_status_id, curator_id, id)
VALUES(1405465, '2024-04-09', '', '', '{{7}}', NULL, 1, NULL, 2801261);
INSERT INTO vendor_locations
(vendor_id, indexed_date, declared_location, parsed_location, curated_countries_ids, curated_date, curated_status_id, curator_id, id)
VALUES(1406024, '2024-04-09', 'Tandil', 'Argentina', '{{7}}', NULL, 1, NULL, 2257572);
INSERT INTO vendor_locations
(vendor_id, indexed_date, declared_location, parsed_location, curated_countries_ids, curated_date, curated_status_id, curator_id, id)
VALUES(1406025, '2024-04-09', 'Berlin, Germany', 'Germany', NULL, NULL, NULL, NULL, 2801814);
INSERT INTO vendor_locations
(vendor_id, indexed_date, declared_location, parsed_location, curated_countries_ids, curated_date, curated_status_id, curator_id, id)
VALUES(1406026, '2024-04-09', '', '', NULL, NULL, NULL, NULL, 2801795);
INSERT INTO vendor_locations
(vendor_id, indexed_date, declared_location, parsed_location, curated_countries_ids, curated_date, curated_status_id, curator_id, id)
VALUES(1485164, '2024-04-09', 'Lima, PE', 'Peru', NULL, NULL, NULL, NULL, 2426204);
INSERT INTO vendor_locations
(vendor_id, indexed_date, declared_location, parsed_location, curated_countries_ids, curated_date, curated_status_id, curator_id, id)
VALUES(1620783, '2024-04-09', '', '', NULL, NULL, NULL, NULL, 957429);
INSERT INTO vendor_locations
(vendor_id, indexed_date, declared_location, parsed_location, curated_countries_ids, curated_date, curated_status_id, curator_id, id)
VALUES(2108508, '2024-04-09', 'Bangalore', 'India', NULL, NULL, NULL, NULL, 357930);
