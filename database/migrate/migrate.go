package migrate

import (
	"log"
	"loom/helper"
	"loom/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func DBMigrate(db *gorm.DB) error {
	migrator := db.Migrator()
	// Migrate table structure
	if !migrator.HasTable(&model.SMEType{}) {
		if err := db.AutoMigrate(&model.SMEType{}); err != nil {
			return err
		}
	}
	// log.Println("Checking if SME table exists...")
	if !migrator.HasTable(&model.SME{}) {
		// log.Println("SME table does not exist, migrating...")
		if err := db.AutoMigrate(&model.SME{}); err != nil {
			return err
		}
		// log.Println("Successfully migrated SME table.")
	}
	// else {
	// 	log.Println("SME table already exists.")
	// }

	if !migrator.HasTable(&model.Skill{}) {
		if err := db.AutoMigrate(&model.Skill{}); err != nil {
			return err
		}
	}

	if !migrator.HasTable(&model.Portofolio{}) {
		if err := db.AutoMigrate(&model.Portofolio{}); err != nil {
			return err
		}
	}

	if !migrator.HasTable(&model.Talent{}) {
		if err := db.AutoMigrate(&model.Talent{}); err != nil {
			return err
		}
	}

	if !migrator.HasTable(&model.Job{}) {
		if err := db.AutoMigrate(&model.Job{}); err != nil {
			return err
		}
	}
	if !migrator.HasTable(&model.ApplicationStatus{}) {
		if err := db.AutoMigrate(&model.ApplicationStatus{}); err != nil {
			return err
		}
	}
	if !migrator.HasTable(&model.TrApplication{}) {
		if err := db.AutoMigrate(&model.TrApplication{}); err != nil {
			return err
		}
	}

	// Seed dummy data
	if err := seedTalentData(db); err != nil {
		return err
	}
	if err := seedSMEData(db); err != nil {
		return err
	}
	if err := SeedSMETypesData(db); err != nil {
		return err
	}
	if err := SeedSkillsData(db); err != nil {
		return err
	}
	if err := seedApplicationStatusData(db); err != nil {
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

	if count == 0 {
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
func SeedSMETypesData(db *gorm.DB) error {
	smeTypes := []model.SMEType{
		{SMETypeID: uuid.New(), SMETypeName: "Beauty, Skincare, Makeup", SMETypeDescription: "Products and services focused on personal grooming and aesthetics."},
		{SMETypeID: uuid.New(), SMETypeName: "Health and Wellness", SMETypeDescription: "Promoting physical and mental well-being through various services and products."},
		{SMETypeID: uuid.New(), SMETypeName: "Fashion and Apparel", SMETypeDescription: "Involves the design, production, and retail of clothing and accessories."},
		{SMETypeID: uuid.New(), SMETypeName: "Home Decor", SMETypeDescription: "Specializes in furnishings and accessories for enhancing living spaces."},
		{SMETypeID: uuid.New(), SMETypeName: "Fitness and Sports", SMETypeDescription: "Offers gym services, fitness classes, and sports equipment."},
		{SMETypeID: uuid.New(), SMETypeName: "Food and Beverage", SMETypeDescription: "Includes restaurants, cafes, catering, and food products."},
		{SMETypeID: uuid.New(), SMETypeName: "Technology and Gadgets", SMETypeDescription: "Selling electronic devices and tech-related services."},
		{SMETypeID: uuid.New(), SMETypeName: "Travel and Hospitality", SMETypeDescription: "Focuses on travel agencies, hotels, and travel experiences."},
		{SMETypeID: uuid.New(), SMETypeName: "Pet Care", SMETypeDescription: "Provides grooming, health products, and training for pets."},
		{SMETypeID: uuid.New(), SMETypeName: "Education and E-Learning", SMETypeDescription: "Offers educational materials, online courses, and tutoring services."},
		{SMETypeID: uuid.New(), SMETypeName: "Sustainable Products", SMETypeDescription: "Concentrates on eco-friendly products and practices."},
		{SMETypeID: uuid.New(), SMETypeName: "Real Estate", SMETypeDescription: "Involves buying, selling, and managing properties."},
		{SMETypeID: uuid.New(), SMETypeName: "Automotive", SMETypeDescription: "Focused on vehicle sales, maintenance, and accessories."},
		{SMETypeID: uuid.New(), SMETypeName: "Financial Services", SMETypeDescription: "Includes banking, insurance, and investment services."},
		{SMETypeID: uuid.New(), SMETypeName: "Event Planning", SMETypeDescription: "Organizing and coordinating events such as weddings and corporate functions."},
		{SMETypeID: uuid.New(), SMETypeName: "Digital Marketing", SMETypeDescription: "Services focused on promoting businesses online through various strategies."},
		{SMETypeID: uuid.New(), SMETypeName: "Arts and Crafts", SMETypeDescription: "Selling handmade goods, art, and craft supplies."},
		{SMETypeID: uuid.New(), SMETypeName: "Consulting Services", SMETypeDescription: "Providing expert advice in various fields, such as business, finance, or IT."},
		{SMETypeID: uuid.New(), SMETypeName: "Online Retail", SMETypeDescription: "E-commerce businesses selling products through online platforms."},
		{SMETypeID: uuid.New(), SMETypeName: "Landscaping and Gardening", SMETypeDescription: "Services related to gardening, landscaping design, and maintenance."},
	}

	var count int64
	if err := db.Model(&model.SMEType{}).Count(&count).Error; err != nil {
		return err
	}

	if count == 0 {
		for _, smeType := range smeTypes {
			if err := db.Create(&smeType).Error; err != nil {
				log.Printf("Error seeding SME type %s: %v\n", smeType.SMETypeName, err)
				return err
			}
		}
		log.Println("Successfully seeded SME types.")
	} else {
		log.Println("SME types already exist, skipping seeding.")
	}

	return nil
}
func SeedSkillsData(db *gorm.DB) error {
	skills := []model.Skill{
		{SkillID: uuid.New(), SkillName: "Business Management"},
		{SkillID: uuid.New(), SkillName: "Digital Marketing"},
		{SkillID: uuid.New(), SkillName: "Data Analysis"},
		{SkillID: uuid.New(), SkillName: "Software Development"},
		{SkillID: uuid.New(), SkillName: "Project Management"},
		{SkillID: uuid.New(), SkillName: "Graphic Design"},
		{SkillID: uuid.New(), SkillName: "Financial Analysis"},
		{SkillID: uuid.New(), SkillName: "Communication"},
		{SkillID: uuid.New(), SkillName: "Sales Strategy"},
		{SkillID: uuid.New(), SkillName: "Customer Service"},
		{SkillID: uuid.New(), SkillName: "Social Media Management"},
		{SkillID: uuid.New(), SkillName: "Leadership"},
		{SkillID: uuid.New(), SkillName: "UX/UI Design"},
		{SkillID: uuid.New(), SkillName: "Negotiation"},
		{SkillID: uuid.New(), SkillName: "Public Speaking"},
		{SkillID: uuid.New(), SkillName: "Networking"},
		{SkillID: uuid.New(), SkillName: "Coding"},
		{SkillID: uuid.New(), SkillName: "Content Creation"},
		{SkillID: uuid.New(), SkillName: "Research"},
		{SkillID: uuid.New(), SkillName: "Time Management"},
		{SkillID: uuid.New(), SkillName: "Web Development"},
		{SkillID: uuid.New(), SkillName: "SEO (Search Engine Optimization)"},
		{SkillID: uuid.New(), SkillName: "Copywriting"},
		{SkillID: uuid.New(), SkillName: "Video Editing"},
		{SkillID: uuid.New(), SkillName: "Photography"},
		{SkillID: uuid.New(), SkillName: "Virtual Assistance"},
		{SkillID: uuid.New(), SkillName: "Translation"},
		{SkillID: uuid.New(), SkillName: "App Development"},
		{SkillID: uuid.New(), SkillName: "E-commerce Management"},
		{SkillID: uuid.New(), SkillName: "Affiliate Marketing"},
	}

	var count int64
	if err := db.Model(&model.Skill{}).Count(&count).Error; err != nil {
		return err
	}

	if count == 0 {
		for _, skill := range skills {
			if err := db.Create(&skill).Error; err != nil {
				log.Printf("Error seeding skill %s: %v\n", skill.SkillName, err)
				return err
			}
		}
		log.Println("Successfully seeded skills.")
	} else {
		log.Println("Skills already exist, skipping seeding.")
	}

	return nil
}

func seedApplicationStatusData(db *gorm.DB) error {
	var count int64
	if err := db.Model(&model.ApplicationStatus{}).Count(&count).Error; err != nil {
		return err
	}

	if count == 0 {
		defaultStatusData := []struct {
			StatusID int
			Status   string
		}{
			{1, "New"},
			{2, "Reviewed"},
			{3, "Interviewed"},
			{4, "Hired"},
			{5, "Rejected"},
		}

		for _, status := range defaultStatusData {
			newStatus := model.ApplicationStatus{
				StatusID: status.StatusID,
				Status:   status.Status,
			}

			if err := db.Create(&newStatus).Error; err != nil {
				return err
			}
		}
	}
	return nil
}
