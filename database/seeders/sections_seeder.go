package seeders

import (
	"time"

	"github.com/sip/simru/entity"
)

type sectionSeeders struct {
	SectionName string
	Authority   string
}

func SectionSeedersUp() {
	seeders := []sectionSeeders{
		{
			SectionName: "PPAK3",
			Authority:   "Customer",
		}, {
			SectionName: "IPK3",
			Authority:   "Vendor/PJK3",
		}, {
			SectionName: "UNWASLIS_K3",
			Authority:   "Pemeriksa dari Pemerintah",
		}, {
			SectionName: "VALAK3",
			Authority:   "Pemerintah sebagai Validator",
		}, {
			SectionName: "BP",
			Authority:   "Badan Perijinan Provinsi",
		},
	}
	section := entity.Sections{}
	for _, seed := range seeders {
		section.SectionName = seed.SectionName
		section.Authority = seed.Authority
		section.CreatedAt = time.Now()
		sectionRepo.CreateSection(section)
	}
}

// Summary
// A. PPAK3 (Customer)
// B. IPK3 ( vendor/pjk3)
// C. UNWASLIS_K3 (Pemeriksa dari Pemerintah)
// D. VALAK3 (Pemerintah sebagai Validator)
// E. BP (Badan Perijinan Provinsi)
