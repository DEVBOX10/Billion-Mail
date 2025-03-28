package do

// PermissionCreateInput defines the input for creating a permission
type PermissionCreateInput struct {
	PermissionName string `json:"permission_name" v:"required#Permission name cannot be empty"`
	Description    string `json:"description"`
	Module         string `json:"module"          v:"required#Module name cannot be empty"`
	Action         string `json:"action"          v:"required#Action type cannot be empty"`
	Resource       string `json:"resource"        v:"required#Resource type cannot be empty"`
	Status         int    `json:"status"          d:"1"`
}

// PermissionUpdateInput defines the input for updating a permission
type PermissionUpdateInput struct {
	PermissionId   int     `json:"permission_id"   v:"required|min:1#Permission ID cannot be empty|Permission ID must be greater than 0"`
	PermissionName *string `json:"permission_name"`
	Description    *string `json:"description"`
	Module         *string `json:"module"`
	Action         *string `json:"action"`
	Resource       *string `json:"resource"`
	Status         *int    `json:"status"`
}

// PermissionGetListInput defines the input for getting permission list
type PermissionGetListInput struct {
	Page     int    `json:"page"       d:"1"`
	PageSize int    `json:"page_size"  d:"20"`
	Keyword  string `json:"keyword"`
	Module   string `json:"module"`
	Action   string `json:"action"`
	Resource string `json:"resource"`
}

// PermissionGetOneInput defines the input for getting a single permission
type PermissionGetOneInput struct {
	PermissionId int `json:"permission_id"  v:"required|min:1#Permission ID cannot be empty|Permission ID must be greater than 0"`
}

// PermissionDeleteInput defines the input for deleting a permission
type PermissionDeleteInput struct {
	PermissionId int `json:"permission_id"  v:"required|min:1#Permission ID cannot be empty|Permission ID must be greater than 0"`
}

// PermissionCheckInput defines the input for checking permission
type PermissionCheckInput struct {
	AccountId int    `json:"account_id"  v:"required|min:1#Account ID cannot be empty|Account ID must be greater than 0"`
	Module    string `json:"module"      v:"required#Module name cannot be empty"`
	Action    string `json:"action"      v:"required#Action type cannot be empty"`
	Resource  string `json:"resource"    v:"required#Resource type cannot be empty"`
}
