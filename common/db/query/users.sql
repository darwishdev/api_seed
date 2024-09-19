-- name: PermissionsClear :exec
TRUNCATE TABLE accounts_schema."permissions" CASCADE;

-- name: UsersClear :exec
TRUNCATE TABLE accounts_schema.users CASCADE;

-- name: RolesClear :exec
TRUNCATE TABLE accounts_schema.roles CASCADE;

-- name: PermissionsBulkCreate :copyfrom
INSERT INTO accounts_schema.permissions(permission_name, permission_function, permission_description, permission_group)
    VALUES ($1, $2, $3, $4);

-- name: RolesBulkCreate :copyfrom
INSERT INTO accounts_schema.roles(role_name, role_description)
    VALUES ($1, $2);

-- name: UsersBulkCreate :copyfrom
INSERT INTO accounts_schema.users(user_name, user_code ,  user_image, user_email,   user_phone, user_password)
    VALUES ($1, $2, $3, $4, $5, $6);

-- name: OwnersBulkCreate :copyfrom
INSERT INTO accounts_schema.owners(owner_name, owner_image, owner_email,   owner_phone, owner_password , owner_national_id)
    VALUES ($1, $2, $3, $4, $5, $6);

-- name: RolePermissionsBulkCreate :copyfrom
INSERT INTO accounts_schema.role_permissions(role_id, permission_id)
    VALUES ($1, $2);

-- name: UserRolesBulkCreate :copyfrom
INSERT INTO accounts_schema.user_roles(user_id, role_id)
    VALUES ($1, $2);

-- name: UserPermissionsBulkCreate :copyfrom
INSERT INTO accounts_schema.user_permissions(user_id, permission_id)
    VALUES ($1, $2);

-- name: NavigationBarsClear :exec
TRUNCATE TABLE accounts_schema.navigation_bars CASCADE;




-- name: NavigationBarCreate :exec
INSERT INTO accounts_schema.navigation_bars(menu_key, label,label_ar ,icon_id, "route", parent_id, permission_id)
    VALUES ($1, $2, $3, $4,$5,(
            SELECT
                navigation_bar_id
            FROM
                accounts_schema.navigation_bars parent
            WHERE
                parent.menu_key = sqlc.narg('parent_key')),
(
                SELECT
                    permission_id
                FROM
                    accounts_schema.permissions
                WHERE
                    permission_function = sqlc.arg('permission_name')));



-- name: IconFindByName :one
SELECT
    icon_id
FROM
    icons
WHERE
    icon_name = $1;
-- name: RoleIdsByNames :many
SELECT
    role_id
FROM
    accounts_schema.roles
WHERE
    role_name = ANY ($1::text[]);

-- name: PermissionIdsByFunctions :many
SELECT
    permission_id
FROM
    accounts_schema.permissions
WHERE
    permission_function = ANY ($1::text[]);

-- name: RolesCount :one
SELECT
    COUNT(*)
FROM
    accounts_schema.roles;




 
-- name: CustomersBulkCreate :copyfrom
INSERT INTO accounts_schema.customers(customer_name, customer_code ,  customer_image, customer_email,   customer_phone, customer_password , customer_national_id)
    VALUES ($1, $2, $3, $4, $5, $6 , $7);
