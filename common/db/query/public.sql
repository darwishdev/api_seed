-- name: SettingTypesClear :exec
TRUNCATE TABLE setting_types CASCADE;

-- name: SettingTypesBulkCreate :copyfrom
INSERT INTO setting_types(setting_type)
    VALUES ($1);

-- name: SettingTypeFindByType :one
SELECT
    setting_type_id
FROM
    setting_types
WHERE
    setting_type = $1;

-- name: SettingsClear :exec
TRUNCATE TABLE settings CASCADE;

-- name: SettingsBulkCreate :copyfrom
INSERT INTO settings(setting_type_id, setting_key, setting_value)
    VALUES ($1, $2, $3);


-- name: IconsBulkCreate :copyfrom
INSERT INTO icons( icon_name,icon_content)
    VALUES ($1, $2);
