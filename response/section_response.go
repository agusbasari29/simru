package response

import "github.com/sip/simru/entity"

type SectionsResponse struct {
	ID          uint64 `json:"id"`
	SectionName string `json:"section_name"`
	Authority   string `json:"authority"`
}

func SectionsResponseFormatter(sections entity.Sections) SectionsResponse {
	formatter := SectionsResponse{}
	formatter.ID = sections.ID
	formatter.SectionName = sections.SectionName
	formatter.Authority = sections.Authority
	return formatter
}