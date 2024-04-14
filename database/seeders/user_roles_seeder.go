package seeders

import (
	"time"

	"github.com/sip/simru/entity"
)

type userRolesSeeder struct {
	Role      string
	RoleName  string
	SectionID uint64
}

func UserRoleSeederUp() {
	seeders := []userRolesSeeder{
		{
			Role:      "101",
			RoleName:  "Pimpinan P2AK3",
			SectionID: 1,
		}, {
			Role:      "102",
			RoleName:  "Admin P2AK3",
			SectionID: 1,
		}, {
			Role:      "103",
			RoleName:  "PIC P2AK3",
			SectionID: 1,
		}, {
			Role:      "201",
			RoleName:  "Pimpinan Inspector K3 (PIPK3)",
			SectionID: 2,
		}, {
			Role:      "202",
			RoleName:  "TU IPK3",
			SectionID: 2,
		}, {
			Role:      "203",
			RoleName:  "Inspector K3 (IK3)",
			SectionID: 2,
		}, {
			Role:      "204",
			RoleName:  "Asisten IK3",
			SectionID: 2,
		}, {
			Role:      "301",
			RoleName:  "Waslis K3 (Atwaslis K3)",
			SectionID: 3,
		}, {
			Role:      "302",
			RoleName:  "Pengawas Spesialis K3 (Waslis K3)",
			SectionID: 3,
		}, {
			Role:      "303",
			RoleName:  "Admin Waslis K3",
			SectionID: 3,
		}, {
			Role:      "304",
			RoleName:  "Staf Waslis K3",
			SectionID: 3,
		}, {
			Role:      "401",
			RoleName:  "Kadis",
			SectionID: 4,
		}, {
			Role:      "402",
			RoleName:  "Subkord K3",
			SectionID: 4,
		}, {
			Role:      "403",
			RoleName:  "TU Valak3",
			SectionID: 4,
		}, {
			Role:      "404",
			RoleName:  "Valak3",
			SectionID: 4,
		}, {
			Role:      "405",
			RoleName:  "Asisten Valak3",
			SectionID: 4,
		}, {
			Role:      "501",
			RoleName:  "Pimpinan BP",
			SectionID: 5,
		}, {
			Role:      "502",
			RoleName:  "TU BP",
			SectionID: 5,
		},
	}

	userRole := entity.UserRoles{}
	for _, seed := range seeders {
		userRole.Role = seed.Role
		userRole.RoleName = seed.RoleName
		userRole.SectionID = seed.SectionID
		userRole.CreatedAt = time.Now()
		userRoleRepo.CreateUserRole(userRole)
	}
}
