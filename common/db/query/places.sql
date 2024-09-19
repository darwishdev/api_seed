-- -- name: CitiesClear :exec
-- TRUNCATE TABLE places_schema.cities CASCADE;

-- -- name: CitiesBulkCreate :copyfrom
-- INSERT INTO places_schema.cities(city_name, city_code)
--     VALUES ($1, $2);

-- -- name: DistrictsClear :exec
-- TRUNCATE TABLE places_schema.districts CASCADE;

-- -- name: DistrictsBulkCreate :copyfrom
-- INSERT INTO places_schema.districts(city_id, district_name, district_code)
--     VALUES ($1, $2, $3);

-- -- name: NeighbourhoodsClear :exec
-- TRUNCATE TABLE places_schema.neighbourhoods CASCADE;

-- -- name: NeighbourhoodsBulkCreate :copyfrom
-- INSERT INTO places_schema.neighbourhoods(district_id, neighbourhood_name, neighbourhood_code)
--     VALUES ($1, $2, $3);

-- -- name: CityFindByName :one
-- SELECT
--     city_id
-- FROM
--     places_schema.cities
-- WHERE
--     city_name = $1;

-- -- name: DistrictFindByName :one
-- SELECT
--     district_id
-- FROM
--     places_schema.districts
-- WHERE
--     district_name = $1;

-- -- name: CitiesList :many
-- SELECT
--     city_id,
--     city_name
-- FROM
--     places_schema.cities;

-- -- name: DistrictsList :many
-- SELECT
--     district_id,
--     district_name
-- FROM
--     places_schema.districts;

-- -- name: AdressesClear :exec
-- TRUNCATE TABLE places_schema.addresses;

-- -- name: AdressesBulkCreate :copyfrom
-- INSERT INTO places_schema.addresses(street, building, floor, flat, remark, address_phone, neighbourhood_id, relation_id, address_relation_type_id)
--     VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);

-- -- name: AdressCreate :one
-- INSERT INTO places_schema.addresses(street, building, floor, flat, remark, address_phone, neighbourhood_id, relation_id, address_relation_type_id)
--     VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
-- RETURNING
--     *;

-- -- name: NeighbourhoodFindByName :one
-- SELECT
--     neighbourhood_id
-- FROM
--     places_schema.neighbourhoods
-- WHERE
--     neighbourhood_name = $1;

