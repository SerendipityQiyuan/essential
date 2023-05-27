package vo

type CreateUpdateUserInfoRequest struct {
	Name      string `json:"name"`
	Sex       string `json:"sex"`
	Age       uint   `json:"age"`
	Introduce string `json:"introduce"`
}
