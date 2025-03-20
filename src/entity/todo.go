package entity

type (
	ListTodoParams struct {
		PaginationParams
		IsDeleted int    `json:"is_deleted" query:"is_deleted"`
		Search    string `json:"search" query:"search"`
	}
	CreateTodoParams struct {
		Title       string `json:"title" validate:"required,min=1,max=255"`
		Description string `json:"description" validate:""`
	}
	CountTodoParams struct {
		IsDeleted int8   `json:"is_deleted" query:"is_deleted" validate:"oneof=0 1"`
		Search    string `json:"search" query:"search"`
	}
)
