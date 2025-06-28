package request

type LearningBySkillsRequestDTO struct {
	SkillIDs []string
}

type LearningBySearchRequestDTO struct {
	Keyword string
}

type LearningDetailRequestDTO struct {
	LearningID string
}
