DROP TABLE if exists vendor_locations;




CREATE TABLE vendor_locations (
	vendor_id int4 NOT NULL,
	indexed_date date  NULL,
	declared_location text  NULL,
	parsed_location text NULL,
	curated_countries_ids text NULL,
	curated_date date NULL,
	curated_status_id int4 NULL,
	curator_id int4 NULL,
	id serial4 NOT NULL,
	    need_human_curation boolean NULL,
    ai_curated_country_name text NULL,
    ai_curated_country_ids integer NULL,
    timezone_based_country text NULL,
    timezone_based_country_certainty integer NULL,
	CONSTRAINT vendor_locations_pkey PRIMARY KEY (id),
	CONSTRAINT vendor_locations_unique UNIQUE (vendor_id, declared_location)
);


-- Insertar datos en la tabla vendor_locations
INSERT INTO vendor_locations (vendor_id, indexed_date, declared_location, parsed_location, curated_countries_ids, curated_date, curated_status_id, curator_id, id, need_human_curation, ai_curated_country_name, ai_curated_country_ids, timezone_based_country, timezone_based_country_certainty) VALUES (3555999, '2024-10-27', 'Argentina', 'Argentina', '{7}', NULL, NULL, NULL, 3318921, 0, 'Argentina', NULL, '?', 0);
INSERT INTO vendor_locations (vendor_id, indexed_date, declared_location, parsed_location, curated_countries_ids, curated_date, curated_status_id, curator_id, id, need_human_curation, ai_curated_country_name, ai_curated_country_ids, timezone_based_country, timezone_based_country_certainty) VALUES (1405397, '2024-04-09', NULL, NULL, '{80}', NULL, NULL, NULL, 2801157, NULL, NULL, NULL, '?', 0);
INSERT INTO vendor_locations (vendor_id, indexed_date, declared_location, parsed_location, curated_countries_ids, curated_date, curated_status_id, curator_id, id, need_human_curation, ai_curated_country_name, ai_curated_country_ids, timezone_based_country, timezone_based_country_certainty) VALUES (1406026, '2024-04-09', NULL, NULL, '{36}', NULL, NULL, NULL, 2801795, NULL, NULL, NULL, 'AU', 7);
INSERT INTO vendor_locations (vendor_id, indexed_date, declared_location, parsed_location, curated_countries_ids, curated_date, curated_status_id, curator_id, id, need_human_curation, ai_curated_country_name, ai_curated_country_ids, timezone_based_country, timezone_based_country_certainty) VALUES (1406025, '2024-04-09', 'Berlin, Germany', 'Germany', '{64}', NULL, NULL, NULL, 2801814, 0, 'Germany', NULL, 'DE', 8);
INSERT INTO vendor_locations (vendor_id, indexed_date, declared_location, parsed_location, curated_countries_ids, curated_date, curated_status_id, curator_id, id, need_human_curation, ai_curated_country_name, ai_curated_country_ids, timezone_based_country, timezone_based_country_certainty) VALUES (283816, '2024-04-09', 'Spain', 'Spain', '{163}', NULL, NULL, NULL, 1460819, NULL, NULL, NULL, NULL, NULL);
INSERT INTO vendor_locations (vendor_id, indexed_date, declared_location, parsed_location, curated_countries_ids, curated_date, curated_status_id, curator_id, id, need_human_curation, ai_curated_country_name, ai_curated_country_ids, timezone_based_country, timezone_based_country_certainty) VALUES (1405464, '2024-04-09', NULL, NULL, NULL, NULL, NULL, NULL, 2801268, NULL, NULL, NULL, 'CO', 7);
INSERT INTO vendor_locations (vendor_id, indexed_date, declared_location, parsed_location, curated_countries_ids, curated_date, curated_status_id, curator_id, id, need_human_curation, ai_curated_country_name, ai_curated_country_ids, timezone_based_country, timezone_based_country_certainty) VALUES (1405399, '2024-04-09', NULL, NULL, '{{7}}', NULL, 1, NULL, 2801163, NULL, NULL, NULL, '?', 0);
INSERT INTO vendor_locations (vendor_id, indexed_date, declared_location, parsed_location, curated_countries_ids, curated_date, curated_status_id, curator_id, id, need_human_curation, ai_curated_country_name, ai_curated_country_ids, timezone_based_country, timezone_based_country_certainty) VALUES (1405400, '2024-04-09', NULL, NULL, '{163}', NULL, NULL, NULL, 2801161, NULL, NULL, NULL, '?', 0);
INSERT INTO vendor_locations (vendor_id, indexed_date, declared_location, parsed_location, curated_countries_ids, curated_date, curated_status_id, curator_id, id, need_human_curation, ai_curated_country_name, ai_curated_country_ids, timezone_based_country, timezone_based_country_certainty) VALUES (1405465, '2024-04-09', NULL, NULL, '{{7}}', NULL, 1, NULL, 2801261, NULL, NULL, NULL, 'BR', 8);
INSERT INTO vendor_locations (vendor_id, indexed_date, declared_location, parsed_location, curated_countries_ids, curated_date, curated_status_id, curator_id, id, need_human_curation, ai_curated_country_name, ai_curated_country_ids, timezone_based_country, timezone_based_country_certainty) VALUES (1406024, '2024-04-09', 'Tandil', 'Argentina', '{{7}}', NULL, 1, NULL, 2257572, NULL, NULL, NULL, 'AR', 6);
INSERT INTO vendor_locations (vendor_id, indexed_date, declared_location, parsed_location, curated_countries_ids, curated_date, curated_status_id, curator_id, id, need_human_curation, ai_curated_country_name, ai_curated_country_ids, timezone_based_country, timezone_based_country_certainty) VALUES (513134, '2024-11-28', 'Nanning, China', 'China', NULL, NULL, NULL, NULL, 3402065, 0, 'China', NULL, 'CN', 9);
INSERT INTO vendor_locations (vendor_id, indexed_date, declared_location, parsed_location, curated_countries_ids, curated_date, curated_status_id, curator_id, id, need_human_curation, ai_curated_country_name, ai_curated_country_ids, timezone_based_country, timezone_based_country_certainty) VALUES (377779, '2024-04-09', 'Germany', 'Germany', '{64}', NULL, NULL, NULL, 1792371, 0, 'Germany', NULL, '?', 0);
INSERT INTO vendor_locations (vendor_id, indexed_date, declared_location, parsed_location, curated_countries_ids, curated_date, curated_status_id, curator_id, id, need_human_curation, ai_curated_country_name, ai_curated_country_ids, timezone_based_country, timezone_based_country_certainty) VALUES (3277960, '2024-09-19', 'Argentina', 'Argentina', '{{7}}', NULL, 1, 5, 2987448, NULL, NULL, NULL, NULL, NULL);
INSERT INTO vendor_locations (vendor_id, indexed_date, declared_location, parsed_location, curated_countries_ids, curated_date, curated_status_id, curator_id, id, need_human_curation, ai_curated_country_name, ai_curated_country_ids, timezone_based_country, timezone_based_country_certainty) VALUES (1408555, '2024-04-09', NULL, NULL, NULL, NULL, NULL, NULL, 2804316, NULL, NULL, NULL, '?', 0);
INSERT INTO vendor_locations (vendor_id, indexed_date, declared_location, parsed_location, curated_countries_ids, curated_date, curated_status_id, curator_id, id, need_human_curation, ai_curated_country_name, ai_curated_country_ids, timezone_based_country, timezone_based_country_certainty) VALUES (1406028, '2024-04-09', NULL, NULL, '{36}', NULL, NULL, NULL, 2801793, NULL, NULL, NULL, 'AU', 7);
INSERT INTO vendor_locations (vendor_id, indexed_date, declared_location, parsed_location, curated_countries_ids, curated_date, curated_status_id, curator_id, id, need_human_curation, ai_curated_country_name, ai_curated_country_ids, timezone_based_country, timezone_based_country_certainty) VALUES (1405398, '2024-04-09', 'San Francisco, C.A., USA', 'United States', '{210}', NULL, NULL, NULL, 2801155, NULL, NULL, NULL, '?', 0);
INSERT INTO vendor_locations (vendor_id, indexed_date, declared_location, parsed_location, curated_countries_ids, curated_date, curated_status_id, curator_id, id, need_human_curation, ai_curated_country_name, ai_curated_country_ids, timezone_based_country, timezone_based_country_certainty) VALUES (1405490, '2024-04-09', 'New York and Singapore', 'Singapore', '{168}', NULL, NULL, NULL, 2801222, NULL, NULL, NULL, '?', 0);
