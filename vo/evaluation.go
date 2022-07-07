package vo

type RequestEvaluation struct {
	PostId string `json:"post_id"`
	Like   uint   `json:"like"`
	IsLike string `json:"is_like"`
}
