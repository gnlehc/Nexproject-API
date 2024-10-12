package migrate

import (
	"help/helper"
	"help/model"
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func DBMigrate(db *gorm.DB) error {
	migrator := db.Migrator()

	if !migrator.HasTable(&model.Portofolio{}) {
		if err := db.AutoMigrate(&model.Portofolio{}); err != nil {
			return err
		}
	}
	if !migrator.HasTable(&model.Skill{}) {
		if err := db.AutoMigrate(&model.Skill{}); err != nil {
			return err
		}
	}
	if !migrator.HasTable(&model.SMEType{}) {
		if err := db.AutoMigrate(&model.SMEType{}); err != nil {
			return err
		}
	}
	if !migrator.HasTable(&model.Talent{}) {
		if err := db.AutoMigrate(&model.Talent{}); err != nil {
			return err
		}
	}
	if !migrator.HasTable(&model.SME{}) {
		if err := db.AutoMigrate(&model.SME{}); err != nil {
			return err
		}
	}
	if !migrator.HasTable(&model.Job{}) {
		if err := db.AutoMigrate(&model.Job{}); err != nil {
			return err
		}
	}

	// seed dummy data
	if err := seedTalentData(db); err != nil {
		return err
	}
	if err := seedSMEData(db); err != nil {
		return err
	}
	return nil
}

func seedTalentData(db *gorm.DB) error {
	var count int64
	if err := db.Model(&model.Talent{}).Count(&count).Error; err != nil {
		return err
	}

	if count == 0 {
		defaultTalentData := []struct {
			Email       string
			Password    string
			FullName    string
			PhoneNumber string
		}{
			{"amanda@gmail.com", "password123", "Jesslyn Amanda", "081292819382"},
			{"john.doe@gmail.com", "johnpassword", "John Doe", "081212345678"},
			{"alice.smith@gmail.com", "alicepass", "Alice Smith", "081298765432"},
			{"bob.jones@gmail.com", "bobsecure", "Bob Jones", "081290123456"},
		}

		for _, talent := range defaultTalentData {
			hashedPassword, err := helper.HashPassword(talent.Password)
			if err != nil {
				log.Printf("Error hashing password for %s: %v\n", talent.Email, err)
				continue
			}

			newTalent := model.Talent{
				TalentID:     uuid.New(),
				Email:        talent.Email,
				Password:     hashedPassword,
				FullName:     talent.FullName,
				HireCount:    0,
				ActiveStatus: true,
				PhoneNumber:  talent.PhoneNumber,
			}

			if err := db.Create(&newTalent).Error; err != nil {
				log.Printf("Error seeding talent %s: %v\n", talent.Email, err)
				return err
			}
		}
	}
	return nil
}

func seedSMEData(db *gorm.DB) error {
	var count int64
	if err := db.Model(&model.SME{}).Count(&count).Error; err != nil {
		return err
	}

	// Only seed data if there are no records
	if count == 0 {
		// Define a slice of dummy SMEs with plain passwords and local UMKM brands
		defaultSMEData := []struct {
			Email       string
			Password    string
			CompanyName string
			CEO         string
			PhoneNumber string
		}{
			{"skintific@gmail.com", "skintific123", "Skintific", "Jessica Lin", "081288273182"},
			{"brodo@gmail.com", "brodo2024", "Brodo", "Yukka Harlanda", "081255512345"},
			{"saffnco@gmail.com", "saffnco2024", "Saff & Co", "Della Putri", "0812123123123"},
			{"erigo@gmail.com", "erigo2024", "Erigo", "Muhammad Sadad", "081298765432"},
			{"sociolla@gmail.com", "sociolla2024", "Sociolla", "John Rasjid", "081234567890"},
		}

		// Iterate over the default SME data and hash their passwords
		for _, sme := range defaultSMEData {
			hashedPassword, err := helper.HashPassword(sme.Password)
			if err != nil {
				log.Printf("Error hashing password for %s: %v\n", sme.Email, err)
				continue
			}
			smeRecord := model.SME{
				SMEID:        uuid.New(),
				Email:        sme.Email,
				Password:     hashedPassword,
				CompanyName:  sme.CompanyName,
				CEO:          sme.CEO,
				ActiveStatus: true,
				PhoneNumber:  sme.PhoneNumber,
			}
			if err := db.Create(&smeRecord).Error; err != nil {
				log.Printf("Error seeding data for %s: %v", sme.CompanyName, err)
				return err
			}
		}
	}

	return nil
}
