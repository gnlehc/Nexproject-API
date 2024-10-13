package repository

import (
	"loom/database"
	"loom/model"
)

type JobRepository interface {
	FindAllJobs() ([]model.Job, error)
	CreateJob(job *model.Job) error
	SaveJob(job *model.Job) error
}

type JobRepo struct{}

func (j *JobRepo) FindAllJobs() ([]model.Job, error) {
	var jobs []model.Job
	if err := database.GlobalDB.Find(&jobs).Error; err != nil {
		return nil, err
	}
	return jobs, nil
}

func (j *JobRepo) CreateJob(job *model.Job) error {
	return database.GlobalDB.Create(job).Error
}

func (j *JobRepo) SaveJob(job *model.Job) error {
	return database.GlobalDB.Save(job).Error
}
