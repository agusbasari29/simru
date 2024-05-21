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
			SectionName: "Pemohon",
			Authority:   "P2AK3",
		}, {
			SectionName: "Balai K3",
			Authority:   "IPAK3",
		}, {
			SectionName: "PJK3",
			Authority:   "IPAK3",
		}, {
			SectionName: "Pemeriksa",
			Authority:   "Disnaker",
		}, {
			SectionName: "Validator",
			Authority:   "Disnaker",
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

// Summary V0
// A. PPAK3 (Customer)
// B. IPK3 ( vendor/pjk3)
// C. UNWASLIS_K3 (Pemeriksa dari Pemerintah)
// D. VALAK3 (Pemerintah sebagai Validator)
// E. BP (Badan Perijinan Provinsi)

// Summary V1
// A. P2AK3
// B. IPAK3 - Balai K3
// C. IPAK3 - PJK3
// D. Disnaker - Pemeriksa
// E. Disnaker - Validator