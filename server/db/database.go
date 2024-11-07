package db

import (
	"bongserver/utils"
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func Connect(filename string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "bongserver.db")
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func CreateTable(conn *sql.DB) error {
	result, err := conn.Exec(`
		CREATE TABLE IF NOT EXISTS Issues (
			-- Student Information
			studentname VARCHAR(255),
			professorname VARCHAR(255),
			section VARCHAR(255),
			year VARCHAR(4),

			-- Lab Information
			labname VARCHAR(255),
			pcnumber INT,

			-- Issue Details
			description TEXT,
			timestamp DATETIME
		);
	`)

	rowsAffected, err := result.RowsAffected()
	utils.LogDebug("Rows affected: %d", rowsAffected)

	return err
}

func checkConnection(conn *sql.DB) error {
	if conn == nil {
		log.Fatalf("Connection is nil")
	}
	return nil
}

func InsertIssue(conn *sql.DB, issue *Issue) error {
	checkConnection(conn)
	conn.Begin()

	result, err := conn.Exec(`
		INSERT INTO Issues (
		  studentname, professorname, section, year, labname, pcnumber, description, timestamp
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`, issue.StudentName, issue.ProfessorName, issue.Section, issue.SchoolYear, issue.LabName, issue.PcNumber, issue.Description, issue.Timestamp)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	utils.LogDebug("Rows affected: %d", rowsAffected)

	return err
}

func PrintIssues(conn *sql.DB) error {
	checkConnection(conn)

	rows, err := conn.Query("SELECT * FROM Issues")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var studentName string
		var professorName string
		var section string
		var year string
		var labName string
		var pcNumber int
		var description string
		var timestamp string

		err := rows.Scan(&studentName, &professorName, &section, &year, &labName, &pcNumber, &description, &timestamp)
		if err != nil {
			return err
		}

		utils.LogDebug(`StudentName: %s, ProfessorName: %s, Section: %s, Year: %s, LabName: %s, PcNumber: %d, Description: %s, Timestamp: %s`, studentName, professorName, section, year, labName, pcNumber, description, timestamp)
	}

	return nil
}
