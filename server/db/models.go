package db

type Issue struct {
	StudentName   string `json:"studentName"`
	ProfessorName string `json:"professorName"`
	Section       string `json:"section"`
	LabName       string `json:"labName"`
	Description   string `json:"description"`
	Timestamp     string `json:"timestamp"`
	SchoolYear    int    `json:"schoolYear"`
	PcNumber      int    `json:"pcNumber"`
}
