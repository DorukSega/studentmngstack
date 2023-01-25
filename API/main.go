package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gin-gonic/gin"
)

var server = "localhost"
var port = 53561

var db *sql.DB

type Department struct {
	Id           int
	BuildingName string
	Name         string
}

type Teacher struct {
	Id         int
	Name       string
	Surname    string
	Isdean     bool
	Wage       int
	Department Department
}

type PStudent struct {
	CourseName string
	Id         int
	Name       string
	Surname    string
	Grade      float64
}

type Student struct {
	Id         int
	Name       string
	Surname    string
	Grade      float64
	Year       int
	Department Department
}

type Course struct {
	Id           int
	Name         string
	Class        Class
	Teacher      Teacher
	Limit        int
	Total_Hours  int
	StartingDate string
	CourseBook   string
}

type Class struct {
	Number     int
	Floor      int
	Department Department
}

func main() {
	var err error

	// Create connection string
	connString := fmt.Sprintf("server=%s;port=%d;trusted_connection=yes;",
		server, port)

	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: " + err.Error())
	}
	log.Printf("Connected!\n")

	//HTTP SERVER
	r := gin.Default()

	r.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello")
	})
	r.GET("/departments", func(c *gin.Context) {
		c.JSON(200, getDepartments())
	})
	r.GET("/teachers", func(c *gin.Context) {
		c.JSON(200, getTeachers())
	})
	r.GET("/students", func(c *gin.Context) {
		c.JSON(200, getStudents())
	})
	r.GET("/courses", func(c *gin.Context) {
		c.JSON(200, getCourses())
	})
	r.GET("/classes", func(c *gin.Context) {
		c.JSON(200, getClasses())
	})
	r.GET("/passingstudents", func(c *gin.Context) {
		c.JSON(200, getPassingStudents())
	})

	r.GET("/student/add", func(c *gin.Context) {
		db.Exec("insert into SchoolProject.dbo.Student(StudentID,StudentName,StudentSurname,Grade,StudentYear,DepartmentID) values (@p1,@p2,@p3,@p4,@p5,@p6)",
			c.Query("id"),
			c.Query("name"),
			c.Query("surname"),
			c.Query("grade"),
			c.Query("year"),
			c.Query("depid"))

		c.String(200, "done")
	})

	r.GET("/teacher/add", func(c *gin.Context) {
		db.Exec("insert into SchoolProject.dbo.Teacher(TeacherID,TeacherName,TeacherSurname,IsDean,Wage,DepartmentID) values (@p1,@p2,@p3,@p4,@p5,@p6)",
			c.Query("id"),
			c.Query("name"),
			c.Query("surname"),
			c.Query("isdean"),
			c.Query("wage"),
			c.Query("depid"))

		c.String(200, "done")
	})

	r.GET("/department/add", func(c *gin.Context) {
		db.Exec("insert into SchoolProject.dbo.Department(DepartmentID,BuildingName,DepartmentName) values (@p1,@p2,@p3)",
			c.Query("id"),
			c.Query("bname"),
			c.Query("dname"))

		c.String(200, "done")
	})

	r.GET("/courses/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err == nil {
			c.JSON(200, getCoursesOfStudent(id))
		}
	})

	r.GET("/remove/student/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err == nil {
			db.Exec("delete from SchoolProject.dbo.Student where StudentID = @p1",
				id)
			c.String(200, "done")
		}
	})
	r.GET("/remove/teacher/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err == nil {
			db.Exec("delete from SchoolProject.dbo.Teacher where TeacherID = @p1",
				id)
			c.String(200, "done")
		}
	})
	r.GET("/remove/department/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err == nil {
			db.Exec("delete from SchoolProject.dbo.Department where DepartmentID = @p1",
				id)
			c.String(200, "done")
		}
	})

	r.GET("/update/students", func(c *gin.Context) {
		var tag = ""
		switch c.Query("tag") {
		case "Id":
			tag = "StudentID"
		case "Name":
			tag = "StudentName"
		case "Surname":
			tag = "StudentSurname"
		case "Grade":
			tag = "Grade"
		case "Year":
			tag = "StudentYear"
		}
		cmnd := fmt.Sprintf("update SchoolProject.dbo.Student set %v = '%v'  where StudentID = %v", tag, c.Query("change"), c.Query("id"))
		db.Exec(cmnd)

		c.String(200, "done")
	})
	r.GET("/update/teachers", func(c *gin.Context) {
		var tag = ""
		switch c.Query("tag") {
		case "Id":
			tag = "TeacherID"
		case "Name":
			tag = "TeacherName"
		case "Surname":
			tag = "TeacherSurname"
		case "Wage":
			tag = "Wage"
		}
		cmnd := fmt.Sprintf("update SchoolProject.dbo.Teacher set %v = '%v'  where TeacherID = %v", tag, c.Query("change"), c.Query("id"))
		db.Exec(cmnd)

		c.String(200, "done")
	})
	r.GET("/update/departments", func(c *gin.Context) {
		var tag = ""
		switch c.Query("tag") {
		case "Id":
			tag = "DepartmentID"
		case "Name":
			tag = "DepartmentName"
		case "BuildingName":
			tag = "BuildingName"
		}
		cmnd := fmt.Sprintf("update SchoolProject.dbo.Department set %v = '%v'  where DepartmentID = %v", tag, c.Query("change"), c.Query("id"))
		db.Exec(cmnd)

		c.String(200, "done")
	})
	r.GET("/update/courses", func(c *gin.Context) {
		var tag = ""
		switch c.Query("tag") {
		case "Name":
			tag = "CourseName"
		case "Limit":
			tag = "Limit"
		case "Total_Hours":
			tag = "Total_Hours"
		case "CourseBook":
			tag = "CourseBook"
		}
		cmnd := fmt.Sprintf("update SchoolProject.dbo.Course set %v = '%v'  where CourseID = %v", tag, c.Query("change"), c.Query("id"))
		db.Exec(cmnd)

		c.String(200, "done")
	})
	r.GET("/update/classes", func(c *gin.Context) {

		cmnd := fmt.Sprintf("update SchoolProject.dbo.Class set ClassFloor = '%v'  where Number = %v", c.Query("change"), c.Query("id"))
		db.Exec(cmnd)
		c.String(200, "done")
	})

	r.Run(":8090")
	//////////////////////////////////////

	// Close the database connection pool after program executes
	defer db.Close()

}

func getDepartment(id int) Department {
	rows, err := db.Query("select * from SchoolProject.dbo.Department where DepartmentID = @p1", id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	ct := Department{}
	rows.Next()
	rows.Scan(&ct.Id, &ct.BuildingName, &ct.Name)

	return ct
}

func getCourse(id int) Course {
	rows, err := db.Query("select * from SchoolProject.dbo.Course where CourseID = @p1", id)
	if err != nil {
		fmt.Println("Error at Course")
		panic(err)
	}
	defer rows.Close()

	ct := Course{}
	rows.Next()
	var tid, cid int
	rows.Scan(&ct.Id, &ct.Total_Hours, &ct.CourseBook, &ct.StartingDate, &ct.Limit, &cid, &ct.Name, &tid)

	ct.Teacher = getTeacher(tid)
	ct.Class = getClass(cid)

	return ct
}

func getClass(id int) Class {
	rows, err := db.Query("select * from SchoolProject.dbo.Class where Number = @p1", id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	ct := Class{}
	rows.Next()

	var deptId int
	err2 := rows.Scan(&ct.Number, &ct.Floor, &deptId)
	if err2 != nil {
		log.Fatal(err)
	}

	ct.Department = getDepartment(deptId)
	return ct
}

func getTeacher(id int) Teacher {
	rows, err := db.Query("select * from SchoolProject.dbo.Teacher where TeacherID = @p1", id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	ct := Teacher{}
	rows.Next()

	var deptId int

	err2 := rows.Scan(&ct.Id, &ct.Name, &ct.Surname, &ct.Isdean, &ct.Wage, &deptId)
	if err2 != nil {
		log.Fatal(err)
	}

	ct.Department = getDepartment(deptId)
	return ct
}

func getStudent(id int) Student {
	rows, err := db.Query("select * from SchoolProject.dbo.Student where StudentID = @p1", id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	ct := Student{}
	rows.Next()
	var deptId int

	err2 := rows.Scan(&ct.Id, &ct.Name, &ct.Surname, &ct.Grade, &deptId, &ct.Year)
	if err2 != nil {
		log.Fatal(err)
	}
	ct.Department = getDepartment(deptId)
	return ct
}

// id is student id
func getCoursesOfStudent(id int) []Course {
	rows, err := db.Query("exec SchoolProject.dbo.GetCoursesOf @p1", id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var result []Course
	for rows.Next() {
		ct := Course{}
		var tid, cid int
		rows.Scan(&ct.Id, &ct.Total_Hours, &ct.CourseBook, &ct.StartingDate, &ct.Limit, &cid, &ct.Name, &tid)
		ct.Teacher = getTeacher(tid)
		ct.Class = getClass(cid)
		result = append(result, ct)
	}
	return result
}

func getDepartments() []Department {
	result := []Department{}
	rows, err := db.Query("select * from SchoolProject.dbo.Department")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		ct := Department{}
		err := rows.Scan(&ct.Id, &ct.BuildingName, &ct.Name)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, ct)
	}
	return result
}

func getTeachers() []Teacher {
	result := []Teacher{}
	rows, err := db.Query("select * from SchoolProject.dbo.Teacher")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		ct := Teacher{}
		var deptId int
		err := rows.Scan(&ct.Id, &ct.Name, &ct.Surname, &ct.Isdean, &ct.Wage, &deptId)
		if err != nil {
			log.Fatal(err)
		}
		ct.Department = getDepartment(deptId)

		result = append(result, ct)
	}
	return result
}

func getStudents() []Student {
	result := []Student{}
	rows, err := db.Query("select * from SchoolProject.dbo.Student")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		ct := Student{}
		var deptId int
		err := rows.Scan(&ct.Id, &ct.Name, &ct.Surname, &ct.Grade, &deptId, &ct.Year)
		if err != nil {
			log.Fatal(err)
		}
		ct.Department = getDepartment(deptId)
		result = append(result, ct)
	}
	return result
}

func getPassingStudents() []PStudent {
	result := []PStudent{}
	rows, err := db.Query("select * from SchoolProject.dbo.PassingStudents")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		ct := PStudent{}
		err := rows.Scan(&ct.CourseName, &ct.Id, &ct.Name, &ct.Surname, &ct.Grade)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, ct)
	}
	return result
}

func getCourses() []Course {

	result := []Course{}
	rows, err := db.Query("select * from SchoolProject.dbo.Course")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		ct := Course{}
		var tid, cid int
		rows.Scan(&ct.Id, &ct.Total_Hours, &ct.CourseBook, &ct.StartingDate, &ct.Limit, &cid, &ct.Name, &tid)
		ct.Teacher = getTeacher(tid)
		ct.Class = getClass(cid)
		result = append(result, ct)
	}
	return result
}

func getClasses() []Class {
	result := []Class{}
	rows, err := db.Query("select * from SchoolProject.dbo.Class")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		ct := Class{}
		var deptId int
		err2 := rows.Scan(&ct.Number, &ct.Floor, &deptId)
		if err2 != nil {
			log.Fatal(err)
		}
		ct.Department = getDepartment(deptId)
		result = append(result, ct)
	}
	return result

}
