CREATE VIEW CalculusStudents AS
    SELECT S.StudentID,StudentName,StudentSurname, Grade,StudentYear FROM Takes AS T,Student AS S
    WHERE T.StudentID= S.StudentID AND T.CourseID = 8000;

CREATE VIEW PassingStudents AS
    SELECT C.CourseName,S.StudentID,StudentName,StudentSurname, Grade FROM Takes AS T,Student AS S, Course AS C
    WHERE T.StudentID= S.StudentID AND T.CourseID = C.CourseID AND Grade>3;