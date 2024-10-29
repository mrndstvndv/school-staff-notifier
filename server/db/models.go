package db

type Issue struct {
	Timestamp string
	Description string
}

type Computer struct {
	Number string
	Issues []Issue
}

func (computer Computer) isFunctioning() bool {
	return len(computer.Issues) == 0
}

type Lab struct {
	Name    string
	Computers []Computer
}

type Library struct {
	Labs []Lab
}

