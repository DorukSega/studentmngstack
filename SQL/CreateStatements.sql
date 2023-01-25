CREATE TABLE Department(
	DepartmentID INT PRIMARY KEY NOT NULL,
	BuildingName VARCHAR(20) NOT NULL,
	DepartmentName VarChar(100) NOT NULL
);

CREATE TABLE Student (
	StudentID INT PRIMARY KEY NOT NULL,
	StudentName VARCHAR(50) NOT NULL,
	StudentSurname VARCHAR(50) NOT NULL,
	Grade INT NULL,
	StudentYear INT NOT NULL,
	DepartmentID INT FOREIGN KEY REFERENCES Department(DepartmentID)
);

CREATE TABLE Class (
	Number INT PRIMARY KEY NOT NULL,
	ClassFloor INT NULL,
	DepartmentID INT FOREIGN KEY REFERENCES Department(DepartmentID)
);

CREATE TABLE Teacher(
	TeacherID INT PRIMARY KEY NOT NULL,
	TeacherName VARCHAR(50) NOT NULL,
	TeacherSurname VARCHAR(50) NOT NULL,
	IsDean BIT,
	Wage INT NULL,
	DepartmentID INT FOREIGN KEY REFERENCES Department(DepartmentID)
);

CREATE TABLE Course(
	CourseID INT PRIMARY KEY NOT NULL,
	Total_Hours INT NULL,
	CourseBook VARCHAR(30),
	StartingDate DATE,
	Limit INT,
	Number INT FOREIGN KEY REFERENCES Class(Number),
	CourseName VARCHAR(50) NOT NULL,
	TeacherID INT FOREIGN KEY REFERENCES Teacher(TeacherID),
);

CREATE TABLE Takes (
	StudentID INT FOREIGN KEY REFERENCES Student(StudentID),
	CourseID INT FOREIGN KEY REFERENCES Course(CourseID),
);