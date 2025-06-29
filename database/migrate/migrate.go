package migrate

import (
	"fmt"
	"log"
	"nexproject/helper"
	"nexproject/model"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func DBMigrate(db *gorm.DB) error {
	migrator := db.Migrator()
	// Migrate table structure
	if !migrator.HasTable(&model.SME{}) {
		if err := db.AutoMigrate(&model.SME{}); err != nil {
			return err
		}
	}

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
	if !migrator.HasTable(&model.Project{}) {
		if err := db.AutoMigrate(&model.Project{}); err != nil {
			return err
		}
	}
	if !migrator.HasTable(&model.Learning{}) {
		if err := db.AutoMigrate(&model.Learning{}); err != nil {
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
	if err := SeedSkillsData(db); err != nil {
		return err
	}
	if err := seedApplicationStatusData(db); err != nil {
		return err
	}
	if err := seedProjectData(db); err != nil {
		return err
	}
	if err := seedLearningData(db); err != nil {
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

func seedProjectData(db *gorm.DB) error {
	var count int64
	if err := db.Model(&model.Project{}).Count(&count).Error; err != nil {
		return err
	}

	if count == 0 {
		var sme model.SME
		if err := db.First(&sme).Error; err != nil {
			return err
		}

		tx := db.Begin()

		projectID := uuid.New()
		project := model.Project{
			ProjectID:          projectID,
			ProjectName:        "NextGen Commerce Platform",
			ProjectDescription: "A modern e-commerce platform with integrated payment.",
			SMEID:              sme.SMEID,
		}

		if err := tx.Create(&project).Error; err != nil {
			tx.Rollback()
			return err
		}

		jobsData := []struct {
			Title       string
			Description string
			Type        string
			Qualif      string
			Arrangement string
			Wage        string
			Active      bool
			Location    string
			Skills      []string
		}{
			{
				"Backend Engineer", "Develop microservices in Go.", "Full-time",
				"3+ years backend development", "Remote", "IDR 15.000.000/month", true, "Jakarta",
				[]string{"Spring Boot", "Golang", "Rust"},
			},
			{
				"Product Manager", "Lead product development lifecycle.", "Contract",
				"Experience in managing agile teams", "Hybrid", "IDR 20.000.000/month", true, "Bandung",
				[]string{"Excel", "PPT", "Power BI"},
			},
		}

		for _, j := range jobsData {
			var skillModels []model.Skill

			for _, skillName := range j.Skills {
				var skill model.Skill
				if err := db.Where("skill_name = ?", skillName).First(&skill).Error; err != nil {
					// skill belum ada → buat baru
					skill = model.Skill{
						SkillID:   uuid.New(),
						SkillName: skillName,
					}
					if err := tx.Create(&skill).Error; err != nil {
						tx.Rollback()
						return fmt.Errorf("failed creating skill: %w", err)
					}
				}
				skillModels = append(skillModels, skill)
			}

			job := model.Job{
				JobID:          uuid.New(),
				ProjectID:      projectID,
				JobTitle:       j.Title,
				JobDescription: j.Description,
				JobType:        j.Type,
				Qualification:  j.Qualif,
				JobArrangement: j.Arrangement,
				Wage:           j.Wage,
				Active:         j.Active,
				CreatedAt:      time.Now(),
				Location:       j.Location,
				Skills:         skillModels,
			}

			if err := tx.Create(&job).Error; err != nil {
				tx.Rollback()
				return fmt.Errorf("failed creating job: %w", err)
			}
		}

		if err := tx.Commit().Error; err != nil {
			return fmt.Errorf("failed committing transaction: %w", err)
		}

		log.Println("Successfully seeded Project with Jobs and Skills.")
	}

	return nil
}

func seedLearningData(db *gorm.DB) error {
	var count int64
	if err := db.Model(&model.Learning{}).Count(&count).Error; err != nil {
		return err
	}

	if count == 0 {
		var skills []model.Skill
		if err := db.Limit(5).Find(&skills).Error; err != nil {
			return err
		}

		learnings := []model.Learning{
			{
				LearningID:    uuid.New(),
				Title:         "Introduction to Digital Marketing",
				Content:       "This learning module introduces basic digital marketing concepts.",
				ImageCoverURL: "https://example.com/cover1.png",
				Skills:        skills[:2],
			},
			{
				LearningID:    uuid.New(),
				Title:         "Basic Data Analysis",
				Content:       "Learn about data cleaning, visualization, and interpretation.",
				ImageCoverURL: "https://example.com/cover2.png",
				Skills:        skills[2:5],
			},
		}

		for _, learning := range learnings {
			if err := db.Create(&learning).Error; err != nil {
				log.Printf("Error seeding learning %s: %v", learning.Title, err)
				return err
			}
		}

		log.Println("Successfully seeded Learning data.")
	}

	return nil
}
