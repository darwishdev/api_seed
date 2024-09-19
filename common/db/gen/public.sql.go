// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: public.sql

package db

import (
	"context"
)

type IconsBulkCreateParams struct {
	IconName    string `json:"icon_name"`
	IconContent string `json:"icon_content"`
}

const settingTypeFindByType = `-- name: SettingTypeFindByType :one
SELECT
    setting_type_id
FROM
    setting_types
WHERE
    setting_type = $1
`

func (q *Queries) SettingTypeFindByType(ctx context.Context, settingType string) (int32, error) {
	row := q.db.QueryRow(ctx, settingTypeFindByType, settingType)
	var setting_type_id int32
	err := row.Scan(&setting_type_id)
	return setting_type_id, err
}

const settingTypesClear = `-- name: SettingTypesClear :exec
TRUNCATE TABLE setting_types CASCADE
`

func (q *Queries) SettingTypesClear(ctx context.Context) error {
	_, err := q.db.Exec(ctx, settingTypesClear)
	return err
}

type SettingsBulkCreateParams struct {
	SettingTypeID int32  `json:"setting_type_id"`
	SettingKey    string `json:"setting_key"`
	SettingValue  string `json:"setting_value"`
}

const settingsClear = `-- name: SettingsClear :exec
TRUNCATE TABLE settings CASCADE
`

func (q *Queries) SettingsClear(ctx context.Context) error {
	_, err := q.db.Exec(ctx, settingsClear)
	return err
}
