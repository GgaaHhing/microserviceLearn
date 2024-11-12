package req

type LoginByPassword struct {
	Mobile string `json:"mobile" binding:"required"`
	Passwd string `json:"passwd" binding:"required, min=6, max=16"`
}
