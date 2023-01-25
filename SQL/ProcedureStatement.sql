CREATE PROCEDURE GetStudentsOf @TeacherID INT
AS
SELECT S.StudentID, S.StudentName, S.StudentSurname, S.Grade FROM Student as S, Takes as T,Course as C
         WHERE S.StudentID = T.StudentID and C.CourseID = T.CourseID AND C.TeacherID = @TeacherID
GO;

EXEC GetStudentsOf @TeacherID = 920;


CREATE PROCEDURE GetCoursesOf @StudentID INT
AS
SELECT DISTINCT C.CourseID,C.Total_Hours,C.CourseBook,C.StartingDate,C.Limit,C.Number,C.CourseName,C.TeacherID FROM  Takes as T,Course as C
WHERE @StudentID = T.StudentID and C.CourseID = T.CourseID
go;

exec GetCoursesOf 2020100;