-- -- name: CategoriesClear :exec
-- TRUNCATE TABLE products_schema.categories CASCADE;

-- -- name: CategoriesBulkCreate :copyfrom
-- INSERT INTO products_schema.categories(category_name, category_image)
--     VALUES ($1, $2);

-- -- name: CategoryFindByName :one
-- SELECT
--     category_id
-- FROM
--     products_schema.categories
-- WHERE
--     category_name = $1;

-- -- name: UnitFindByName :one
-- SELECT
--     unit_id
-- FROM
--     products_schema.units
-- WHERE
--     CONCAT(unit_buy, ' / ', unit_sell, ' / ', unit_ratio) = $1;

-- -- name: ProductsClear :exec
-- TRUNCATE TABLE products_schema.products CASCADE;

-- -- name: ProductsBulkCreate :copyfrom
-- INSERT INTO products_schema.products(product_name, product_code, product_image, product_description, is_final, category_id, unit_id, product_cost, product_price)
--     VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);

-- -- name: UnitsClear :exec
-- TRUNCATE TABLE products_schema.units CASCADE;

-- -- name: UnitsBulkCreate :copyfrom
-- INSERT INTO products_schema.units(unit_buy, unit_sell, unit_ratio)
--     VALUES ($1, $2, $3);

-- -- name: ProductFindByName :one
-- SELECT
--     product_id
-- FROM
--     products_schema.products
-- WHERE
--     product_name = $1; 