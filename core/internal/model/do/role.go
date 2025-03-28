package do

// RoleCreateInput defines the input for creating a role
type RoleCreateInput struct {
	RoleName      string `json:"role_name"   v:"required#Role name cannot be empty"`
	Description   string `json:"description"`
	Status        int    `json:"status"      d:"1"`
	PermissionIds []int  `json:"permission_ids"`
}

// RoleUpdateInput defines the input for updating a role
type RoleUpdateInput struct {
	RoleId        int     `json:"role_id"    v:"required|min:1#Role ID cannot be empty|Role ID must be greater than 0"`
	RoleName      *string `json:"role_name"`
	Description   *string `json:"description"`
	Status        *int    `json:"status"`
	PermissionIds []int   `json:"permission_ids"`
}

// RoleGetListInput defines the input for getting role list
type RoleGetListInput struct {
	Page     int    `json:"page"       d:"1"`
	PageSize int    `json:"page_size"  d:"20"`
	Keyword  string `json:"keyword"`
}

// RoleGetOneInput defines the input for getting a single role
type RoleGetOneInput struct {
	RoleId int `json:"role_id"  v:"required|min:1#Role ID cannot be empty|Role ID must be greater than 0"`
}

// RoleDeleteInput defines the input for deleting a role
type RoleDeleteInput struct {
	RoleId int `json:"role_id"  v:"required|min:1#Role ID cannot be empty|Role ID must be greater than 0"`
}

// RoleWithPermissions defines role information with permissions
type RoleWithPermissions struct {
	RoleId          int      `json:"role_id"`
	RoleName        string   `json:"role_name"`
	Description     string   `json:"description"`
	Status          int      `json:"status"`
	CreateTime      int      `json:"create_time"`
	UpdateTime      int      `json:"update_time"`
	Permissions     []int    `json:"permissions"`      // Permission IDs
	PermissionNames []string `json:"permission_names"` // Permission names
}
