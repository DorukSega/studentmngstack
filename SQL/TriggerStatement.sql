CREATE TRIGGER LimitOfCourse
    ON Takes
    AFTER INSERT AS
    DECLARE @CourseID INT;
    DECLARE @Limit INT;
    DECLARE @Current INT;
    BEGIN
        SELECT @CourseID = CourseID FROM inserted;
        SELECT @Limit =Limit FROM Course AS C WHERE C.CourseID = @CourseID;
        SELECT @Current = COUNT(*) FROM Takes AS T WHERE T.CourseID = @CourseID;
        IF  @Limit < @Current
            UPDATE Course SET Limit = Limit + 1 WHERE CourseID = @CourseID;
    END;

CREATE TRIGGER DeanLimit
        ON Teacher
        AFTER  INSERT AS
        DECLARE @DepartmentID INT;
        DECLARE @Current INT;
    BEGIN
        IF (SELECT IsDean FROM inserted) = 1
            BEGIN
                SELECT @DepartmentID = DepartmentID FROM inserted;
                SELECT @Current = COUNT(*) FROM Teacher AS T WHERE T.DepartmentID = @DepartmentID and T.IsDean = 1;
                IF  @Current > 0
                    BEGIN
                        UPDATE Teacher SET IsDean=0 WHERE IsDean = 1;
                        INSERT INTO Teacher
                        SELECT * FROM inserted;
                    END
            END

    END;


