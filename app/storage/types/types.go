package types

import (
	db "github.com/meloneg/mln_data_pool/common/db/gen"
	"github.com/meloneg/mln_data_pool/supabase"
	// supa "github.com/darwishdev/supabase-go"
)

type UsersBulkCreateResponse struct {
	User      db.UsersBulkCreateParams
	AuthUser  supabase.UserCredentials
	UserID    int
	RolesLen  int
	UserRoles []string
}
type RolesBulkCreateResponse struct {
	Role            db.RolesBulkCreateParams
	RolePermissions []string
	RoleID          int
	PermissionsLen  int
}
