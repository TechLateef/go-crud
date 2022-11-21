package dto

type CreateToDoDto struct {
	Title       string `json:"title" form:"title" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	UserID      uint64 `json:"userid,omitempty" form:"user_id,omitempty"`
}

type UpdateToDoDto struct {
	Title       string `json:"title" form:"title" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	UserID      uint64 `json:"user_id,omitempty" form:"user_id,omitempty"`
}
