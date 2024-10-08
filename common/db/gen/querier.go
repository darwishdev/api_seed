// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1

package db

import (
	"context"
)

type Querier interface {
	CustomersBulkCreate(ctx context.Context, arg []CustomersBulkCreateParams) (int64, error)
	IconFindByName(ctx context.Context, iconName string) (int32, error)
	IconsBulkCreate(ctx context.Context, arg []IconsBulkCreateParams) (int64, error)
	NavigationBarCreate(ctx context.Context, arg NavigationBarCreateParams) error
	NavigationBarsClear(ctx context.Context) error
	OwnersBulkCreate(ctx context.Context, arg []OwnersBulkCreateParams) (int64, error)
	PermissionIdsByFunctions(ctx context.Context, dollar_1 []string) ([]int32, error)
	PermissionsBulkCreate(ctx context.Context, arg []PermissionsBulkCreateParams) (int64, error)
	PermissionsClear(ctx context.Context) error
	RoleIdsByNames(ctx context.Context, dollar_1 []string) ([]int32, error)
	RolePermissionsBulkCreate(ctx context.Context, arg []RolePermissionsBulkCreateParams) (int64, error)
	RolesBulkCreate(ctx context.Context, arg []RolesBulkCreateParams) (int64, error)
	RolesClear(ctx context.Context) error
	RolesCount(ctx context.Context) (int64, error)
	SettingTypeFindByType(ctx context.Context, settingType string) (int32, error)
	SettingTypesBulkCreate(ctx context.Context, settingType []string) (int64, error)
	SettingTypesClear(ctx context.Context) error
	SettingsBulkCreate(ctx context.Context, arg []SettingsBulkCreateParams) (int64, error)
	SettingsClear(ctx context.Context) error
	UserPermissionsBulkCreate(ctx context.Context, arg []UserPermissionsBulkCreateParams) (int64, error)
	UserRolesBulkCreate(ctx context.Context, arg []UserRolesBulkCreateParams) (int64, error)
	UsersBulkCreate(ctx context.Context, arg []UsersBulkCreateParams) (int64, error)
	UsersClear(ctx context.Context) error
}

var _ Querier = (*Queries)(nil)
