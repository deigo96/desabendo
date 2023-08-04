package domain

type ErrorResponse struct {
	Message string `json:"message"`
}

type LoginResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	IsActive bool   `json:"is_active"`
	Token    string `json:"token"`
}

type AuthResponse struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	IsActive  string `json:"is_active"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Role      string `json:"role"`
}

type Pengajuan struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Job       string `json:"job"`
	Age       int    `json:"age"`
	CreatedAt string `json:"created_at"`
}

type ProfileAuth struct {
	Username string
	Role     int8
	UserId   int
}