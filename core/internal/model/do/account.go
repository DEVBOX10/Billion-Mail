package do

// AccountCreateInput defines the input for creating an account
type AccountCreateInput struct {
	Username string `json:"username"   v:"required#Username cannot be empty"`
	Password string `json:"password"   v:"required#Password cannot be empty"`
	Email    string `json:"email"      v:"required|email#Email cannot be empty|Invalid email format"`
	Status   int    `json:"status"     d:"1"`
	Language string `json:"language"   d:"en"`
	RoleIds  []int  `json:"role_ids"`
}

// AccountUpdateInput defines the input for updating an account
type AccountUpdateInput struct {
	AccountId int     `json:"account_id"  v:"required|min:1#Account ID cannot be empty|Account ID must be greater than 0"`
	Password  *string `json:"password"`
	Email     *string `json:"email"       v:"email#Invalid email format"`
	Status    *int    `json:"status"`
	Language  *string `json:"language"`
	RoleIds   []int   `json:"role_ids"`
}

// AccountGetListInput defines the input for getting account list
type AccountGetListInput struct {
	Page     int    `json:"page"       d:"1"`
	PageSize int    `json:"page_size"  d:"20"`
	Keyword  string `json:"keyword"`
}

// AccountGetOneInput defines the input for getting a single account
type AccountGetOneInput struct {
	AccountId int `json:"account_id"  v:"required|min:1#Account ID cannot be empty|Account ID must be greater than 0"`
}

// AccountDeleteInput defines the input for deleting an account
type AccountDeleteInput struct {
	AccountId int `json:"account_id"  v:"required|min:1#Account ID cannot be empty|Account ID must be greater than 0"`
}

// AccountLoginInput defines the input for account login
type AccountLoginInput struct {
	Username string `json:"username"   v:"required#Username cannot be empty"`
	Password string `json:"password"   v:"required#Password cannot be empty"`
	Language string `json:"language"   d:"en"`
}

// AccountWithRoles defines account information with roles
type AccountWithRoles struct {
	AccountId     int      `json:"account_id"`
	Username      string   `json:"username"`
	Email         string   `json:"email"`
	Status        int      `json:"status"`
	Language      string   `json:"language"`
	LastLoginTime int      `json:"last_login_time"`
	CreateTime    int      `json:"create_time"`
	UpdateTime    int      `json:"update_time"`
	Roles         []int    `json:"roles"`      // Role IDs
	RoleNames     []string `json:"role_names"` // Role names
}
