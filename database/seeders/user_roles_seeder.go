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
			RoleName:  "Pimpinan Perusahaan/CEO",
			SectionID: 1,
		}, {
			Role:      "102",
			RoleName:  "PIC HR",
			SectionID: 1,
		}, {
			Role:      "103",
			RoleName:  "PIC HSE",
			SectionID: 1,
		}, {
			Role:      "201",
			RoleName:  "Pimpinan",
			SectionID: 2,
		}, {
			Role:      "202",
			RoleName:  "Staf TU",
			SectionID: 2,
		}, {
			Role:      "203",
			RoleName:  "Pemeriksa K3",
			SectionID: 2,
		}, {
			Role:      "204",
			RoleName:  "Staf Pemeriksa K3",
			SectionID: 2,
		}, {
			Role:      "301",
			RoleName:  "Pimpinan",
			SectionID: 3,
		}, {
			Role:      "302",
			RoleName:  "Marketing",
			SectionID: 3,
		}, {
			Role:      "303",
			RoleName:  "Admin",
			SectionID: 3,
		}, {
			Role:      "304",
			RoleName:  "Ahli K3",
			SectionID: 3,
		}, {
			Role:      "305",
			RoleName:  "Asisten Ahli K3",
			SectionID: 3,
		}, {
			Role:      "401",
			RoleName:  "Spesialis K3 Pemeriksa",
			SectionID: 4,
		}, {
			Role:      "402",
			RoleName:  "Waslis",
			SectionID: 4,
		}, {
			Role:      "403",
			RoleName:  "Asisten Waslis",
			SectionID: 4,
		}, {
			Role:      "404",
			RoleName:  "Staf TU",
			SectionID: 4,
		}, {
			Role:      "501",
			RoleName:  "Kabid",
			SectionID: 5,
		}, {
			Role:      "502",
			RoleName:  "Admin",
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
