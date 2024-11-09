package db

import (
	"bongserver/protobuf"
	"bongserver/utils"
	"database/sql"
	"encoding/json"
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

/*
TODO: Make a table for a class

CREATE TABLE IF NOT EXISTS class (

	id INTEGER PRIMARY KEY AUTOINCREMENT,
	year INTEGER NOT NULL,
	section TEXT NOT NULL,
	course TEXT NOT NULL,
	professor TEXT NOT NULL

);
*/
func CreateTable(conn *sql.DB) error {
	result, err := conn.Exec(`
		CREATE TABLE IF NOT EXISTS Students (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			firstName TEXT NOT NULL,
			lastName TEXT NOT NULL,
			year INTEGER NOT NULL,
			section TEXT NOT NULL,
			course TEXT NOT NULL,
			professor TEXT NOT NULL
		);

		CREATE TABLE IF NOT EXISTS Issues (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			studentId INTEGER NOT NULL,
			labRoom TEXT NOT NULL,
			pcNumber INTEGER NOT NULL,
			concern TEXT NOT NULL,
			timestamp TEXT NOT NULL,
			issues TEXT NOT NULL,
			status INTEGER NOT NULL DEFAULT 0,
			FOREIGN KEY (studentId) REFERENCES students(id)
		);

		CREATE TABLE IF NOT EXISTS class (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			year INTEGER NOT NULL,
			section TEXT NOT NULL,
			course TEXT NOT NULL,
			professor TEXT NOT NULL
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

func getStudentID(db *sql.DB, student *protobuf.Student) (int64, bool, error) {
	checkConnection(db)

	var id int64
	err := db.QueryRow(`SELECT id FROM Students WHERE firstName = ? AND lastName = ? AND year = ? AND section = ? AND course = ? AND professor = ?`,
		student.FirstName, student.LastName, student.Year, student.Section, student.Course, student.Professor).Scan(&id)
	if err == sql.ErrNoRows {
		return 0, false, nil
	}
	return id, err == nil, err
}

func insertStudent(db *sql.DB, student *protobuf.Student) (int64, error) {
	checkConnection(db)

	result, err := db.Exec(`INSERT INTO Students (firstName, lastName, year, section, course, professor) VALUES (?, ?, ?, ?, ?, ?)`,
		student.FirstName, student.LastName, student.Year, student.Section, student.Course, student.Professor)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func GetIssues(db *sql.DB) (*protobuf.IssueList, error) {
	checkConnection(db)

	result, err := db.Query(`SELECT s.firstName, s.lastName, s.year, s.section, s.course, s.professor, i.labRoom, i.pcNumber, i.concern, i.timestamp, i.issues, i.status FROM Students s INNER JOIN Issues i ON s.id = i.studentId`)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	var issues []*protobuf.Issue
	for result.Next() {
		// TODO: rename idk
		issue := &protobuf.Issue{Student: &protobuf.Student{}}
		var idk string

		err := result.Scan(
			&issue.Student.FirstName,
			&issue.Student.LastName,
			&issue.Student.Year,
			&issue.Student.Section,
			&issue.Student.Course,
			&issue.Student.Professor,
			&issue.LabRoom,
			&issue.PcNumber,
			&issue.Concern,
			&issue.Timestamp,
			&idk,
			&issue.Status,
		)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal([]byte(idk), &issue.Issues)
		if err != nil {
			return nil, err
		}

		issues = append(issues, issue)
	}

	// Check for errors from iterating over rows
	if err := result.Err(); err != nil {
		return nil, err
	}

	return &protobuf.IssueList{Issues: issues}, nil
}

func UpdateIssueStatus(db *sql.DB, id int64, status int32) error {
	checkConnection(db)

	_, err := db.Exec(`UPDATE Issues SET status = ? WHERE id = ?`, status, id)
	return err
}

// TODO: return the id of the inserted issue
func InsertIssue(db *sql.DB, issue *protobuf.Issue) error {
	checkConnection(db)

	issuesJSON, err := json.Marshal(issue.Issues)
	if err != nil {
		return err
	}

	studentID, exists, err := getStudentID(db, issue.Student)
	if err != nil {
		return err
	}

	if !exists {
		studentID, err = insertStudent(db, issue.Student)
		if err != nil {
			return err
		}
	}

	_, err = db.Exec(`INSERT INTO Issues (studentId, labRoom, pcNumber, concern, timestamp, issues) VALUES (?, ?, ?, ?, ?, ?)`,
		studentID, issue.LabRoom, issue.PcNumber, issue.Concern, issue.Timestamp, string(issuesJSON))

	return err
}
