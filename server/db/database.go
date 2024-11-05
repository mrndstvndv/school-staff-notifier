package db

import (
	"bongserver/utils"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	cfg := mysql.Config{
		User:                 os.Getenv("DBUSER"),
		Passwd:               os.Getenv("DBPASS"),
		AllowNativePasswords: true,
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "techalert",
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
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
		CREATE TABLE IF NOT EXISTS issue (
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

func InsertIssue(conn *sql.DB, issue *Issue) error {
	if conn == nil {
		log.Fatalf("Connection is nil")
		return fmt.Errorf("Connection is nil")
	}

	result, err := conn.Exec(`
		INSERT INTO issue (
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
