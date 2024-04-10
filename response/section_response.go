package response

type SectionResponse struct {
	ID          uint64 `json:"id"`
	SectionName string `json:"section_name"`
	Authority   string `json:"authority"`
}
