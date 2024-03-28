package commonModel

type SignTaskActor struct {
	Actor          *Actor          `json:"actor"`
	FillFields     []FillField     `json:"fillFields"`
	SignFields     []SignField     `json:"signFields"`
	SignConfigInfo *SignConfigInfo `json:"signConfigInfo"`
}
