# Student Management Stack
Very basic [Three Tier Architecture](https://en.wikipedia.org/wiki/Multitier_architecture#Three-tier_architecture) - FRONTEND & API & DMBS SERVER - implemented for a school project.

![Example image showcasing the frontend interface](https://i.imgur.com/a8KAx0L.png)

Consists of SQL statements written for MSQL, API server written on GO that serves the information on a MSQL server and a simple website that serves as a admin panel.

Written and implemented very quickly so some parts of the code quite hacky and messy.

## Build Instructions
Since this is a three tier architecture, testing on local is bit tricky.

API and SQL server must be running to be able to use Web Interface.

Dependencies:
- Go environment
- Go packages that are initilized in go.mod file

Steps after dependencies are met:
1. Start the Database
2. inside main.go in the API folder change the port accordingly
3. inside API folder run following command "go run .\main.go"
	this will start a local instance of the api
4. inside WEB folder start the index.html
	if working with a local api server, cors bypass is needed
	if api is hosted, change `const url = "http://127.0.0.1:8090"` inside main.js to the host ip address
