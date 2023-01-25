INSERT INTO Department
VALUES (1000,'Block A', 'Department Of Mathematics');

INSERT INTO Department
VALUES (1001,'Block B', 'Department Of Biology');

INSERT INTO Department
VALUES (1002,'Block C', 'Department Of Economics');

INSERT INTO Class
VALUES (100,1,1000);

INSERT INTO Class
VALUES (101,1,1000);

INSERT INTO Class
VALUES (102,2,1000);

INSERT INTO Class
VALUES (200,1,1001);

INSERT INTO Class
VALUES (201,1,1001);

INSERT INTO Class
VALUES (202,1,1001);

INSERT INTO Class
VALUES (300,1,1002);

INSERT INTO Class
VALUES (301,2,1002);

INSERT INTO Class
VALUES (302,3,1002);

INSERT INTO Teacher
VALUES (910,'Rich','Evans',1,20000,1001);

INSERT INTO Teacher
VALUES (911,'Mike','Deniro',0,5000,1001);

INSERT INTO Teacher
VALUES (912,'Adam','Smith',0,5000,1001);

INSERT INTO Teacher
VALUES (900,'David','Lynch',1,20000,1000);

INSERT INTO Teacher
VALUES (901,'Denis','Villenue',0,10000,1000);

INSERT INTO Teacher
VALUES (902,'Liam','Neelson',0,10000,1000);

INSERT INTO Teacher
VALUES (920,'Kevin','Jenkins',1,25000,1002);

INSERT INTO Teacher
VALUES (921,'John','DOE',0,10000,1002);

INSERT INTO Teacher
VALUES (922,'Dave','Riller',0,10000,1002);

INSERT INTO Course
VALUES (8000,40,'Thomas Calculus 4th Edition','05-09-2020',30,100,'Calculus 1',900);

INSERT INTO Course
VALUES (8001,30,'Differential Equations 2th Edition','03-09-2020',20,101,'Differential Equations',901);

INSERT INTO Course
VALUES (8002,30,'Linear Algebra International Edition','07-09-2020',25,102,'Linear Algebra',902);

INSERT INTO Course
VALUES (8001,30,'Differential Equations 2th','03-09-2020',20,101,'Differential Equations',901);

INSERT INTO Course
VALUES (8002,30,'Linear Algebra International Edition','07-09-2020',25,102,'Linear Algebra',902);

INSERT INTO Course
VALUES (8002,30,'Linear Algebra Second Ed','07-09-2020',25,102,'Linear Algebra',902)

INSERT INTO Course
VALUES (8100,40,'Introduction to Biology','05-09-2020',30,200,'Introductory Biology',910);

INSERT INTO Course
VALUES (8101,30,'Genetics 101','04-09-2020',20,201,'Genetics',911);

INSERT INTO Course
VALUES (8102,30,'Cellular Integration','07-09-2020',25,202,'Cell Biology',912);

INSERT INTO Course
VALUES (8200,40,'Economic Theory','06-09-2020',30,300,'Microeconomic Theory',920);

INSERT INTO Course
VALUES (8201,25,'Macroeconomy Solutions','04-09-2020',20,301,'Principles of Macroeconomics',921);

INSERT INTO Course
VALUES (8202,30,'How to Control the Market','03-09-2020',25,302,'Market Design',922);
