package db

import (
	"bongserver/protobuf"
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
	CREATE TABLE IF NOT EXISTS students (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		first_name TEXT NOT NULL,
		last_name TEXT NOT NULL,
		section TEXT NOT NULL,
		course TEXT NOT NULL,
		professor TEXT NOT NULL,
		year INTEGER NOT NULL
	);

	CREATE TABLE IF NOT EXISTS issues (
		id INTEGER PRIMARY KEY,
		lab_room TEXT NOT NULL,
		concern TEXT NOT NULL,
		timestamp TEXT NOT NULL,
		pc_number INTEGER NOT NULL,
		student_id INTEGER NOT NULL,
		urgency INTEGER NOT NULL,
		FOREIGN KEY (student_id) REFERENCES students(id)
	);

	CREATE TABLE IF NOT EXISTS faulty_components (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		fixed BOOLEAN NOT NULL DEFAULT 0,
		parent_id INTEGER NOT NULL,
		FOREIGN KEY (parent_id) REFERENCES issues(id)
	);

	-- Create indexes for better query performance
	CREATE INDEX idx_issues_student ON issues(student_id);
	CREATE INDEX idx_components_parent ON faulty_components(parent_id);
	`)

	if err != nil {
		return err
	}

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

	if (student) == nil {
		log.Fatalf("Student is nil")
	}

	var id int64
	err := db.QueryRow(`SELECT id FROM students WHERE first_name = ? AND last_name = ? AND year = ? AND section = ? AND course = ? AND professor = ?`,
		student.FirstName, student.LastName, student.Year, student.Section, student.Course, student.Professor).Scan(&id)

	if err == sql.ErrNoRows {
		return 0, false, nil
	}
	return id, err == nil, err
}

func insertStudent(db *sql.DB, student *protobuf.Student) (int64, error) {
	checkConnection(db)

	result, err := db.Exec(`INSERT INTO students (first_name, last_name, year, section, course, professor) VALUES (?, ?, ?, ?, ?, ?)`,
		student.FirstName, student.LastName, student.Year, student.Section, student.Course, student.Professor)
	if err != nil {
		return 0, err
	}
	issueID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return issueID, nil
}

// insertFaultyComponents inserts all faulty components for a given issue
func insertFaultyComponents(db *sql.DB, issueID int64, components []*protobuf.FaultyComponent) ([]int64, error) {
	if len(components) == 0 {
		return nil, nil
	}

	var ids []int64

	log.Printf("hello?")

	for _, component := range components {
		result, err := db.Exec(`INSERT INTO faulty_components (name, fixed, parent_id) VALUES (?, ?, ?)`,
			component.Name, component.Fixed, issueID)

		faultyComponentId, err := result.LastInsertId()
		ids = append(ids, faultyComponentId)

		if err != nil {
			return nil, err
		}
	}

	log.Printf("ids: %v", ids)

	return ids, nil
}

func GetIssues(db *sql.DB) (*protobuf.IssueList, error) {
	checkConnection(db)

	result, err := db.Query(`SELECT s.first_name, s.last_name, s.year, s.section, s.course, s.professor, i.lab_room, i.pc_number, i.concern, i.timestamp, i.id, i.urgency FROM students s INNER JOIN issues i ON s.id = i.student_id`)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	var issues []*protobuf.Issue
	for result.Next() {
		issue := &protobuf.Issue{Student: &protobuf.Student{}}

		log.Printf("%v", &issue.Urgency)

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
			&issue.Id,
			&issue.Urgency,
		)
		if err != nil {
			return nil, err
		}

		log.Print("Got here\n")

		faultyComponents, err := getFaultyComponents(db, issue.Id)
		if err != nil {
			return nil, err
		}

		issue.FaultyComponents = faultyComponents
		log.Printf("%d Faulty Components: %v", issue.Id, faultyComponents)

		issues = append(issues, issue)
	}

	// Check for errors from iterating over rows
	if err := result.Err(); err != nil {
		return nil, err
	}

	return &protobuf.IssueList{Issues: issues}, nil
}

func getFaultyComponents(db *sql.DB, issueID int64) ([]*protobuf.FaultyComponent, error) {
	rows, err := db.Query(`
		SELECT id, name, fixed, parent_id 
		FROM faulty_components 
		WHERE parent_id = ?
	`, issueID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var components []*protobuf.FaultyComponent

	for rows.Next() {
		component := &protobuf.FaultyComponent{}
		err := rows.Scan(&component.Id, &component.Name, &component.Fixed, &component.ParentId)
		if err != nil {
			return nil, err
		}

		components = append(components, component)
	}

	log.Printf("Components: %v", components)

	return components, nil
}

// INFO: a status of 1 means its fixed
func UpdateComponentStatus(db *sql.DB, id string, status string) error {
	checkConnection(db)

	result, err := db.Exec(`UPDATE faulty_components SET fixed = ? WHERE id = ?`, status, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	log.Printf("Rows affected: %d", rowsAffected)
	return nil
}

func InsertIssue(db *sql.DB, issue *protobuf.Issue) (int64, []int64, error) {
	checkConnection(db)

	studentID, exists, err := getStudentID(db, issue.Student)
	if err != nil {
		return 0, nil, err
	}

	if !exists {
		studentID, err = insertStudent(db, issue.Student)
		if err != nil {
			return 0, nil, err
		}
	}

	result, err := db.Exec(`INSERT INTO issues (student_id, lab_room, pc_number, concern, timestamp, urgency) VALUES (?, ?, ?, ?, ?, ?)`,
		studentID, issue.LabRoom, issue.PcNumber, issue.Concern, issue.Timestamp, issue.Urgency)
	if err != nil {
		return 0, nil, err
	}

	issueId, err := result.LastInsertId()
	if err != nil {
		return 0, nil, err
	}

	ids, err := insertFaultyComponents(db, issueId, issue.FaultyComponents)
	if err != nil {
		return 0, nil, err
	}

	return issueId, ids, nil
}
