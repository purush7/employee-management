### Employee Management


> The internal directory named as internal_ext instead of internal, as internal won't be visible to everyone 


#### Problem:



You are tasked with creating a Go application that manages a simple employee database or in-memory store. Additionally, you need to implement a RESTful API with pagination for listing employee records.

Requirements:

1. Employee Struct:

- Define a struct named `Employee` with the following fields:

- `ID` (int): Unique identifier for the employee.

- `Name` (string): Name of the employee.

- `Position` (string): Position/title of the employee.

- `Salary` (float64): Salary of the employee.

2. CRUD Operations:

- Implement functions/methods to perform CRUD operations on the employee database or in-memory store:

- `CreateEmployee`: Adds a new employee to the database or store.

- `GetEmployeeByID`: Retrieves an employee from the database or store by ID.

- `UpdateEmployee`: Updates the details of an existing employee.

- `DeleteEmployee`: Deletes an employee from the database or store by ID.

3. Concurrency:

- Ensure that the application is safe for concurrent use by using appropriate synchronization mechanisms.

4. Testing:

- Write unit tests to cover the CRUD operations and ensure the correctness of the implementation.

5. RESTful API with Pagination:

- Implement a RESTful API for listing employee records with pagination.

- The API should provide endpoints for listing employees with support for pagination.

- Each page should contain a configurable number of records.

- Implement proper error handling and response formatting for the API endpoints.


#### Unit testing

To run the test, use command
```
go test -v ./...
```

To get the test coverage, use below commands

to output the coverage into c.out file
```
go test -coverprofile=c.out -v ./...        
```

and to view the coverage, use 
```
go tool cover -html=c.out
```
which opens on browser with more detail information

Currently the test coverage is 100% in 1 file and other has 95%


### Run and Testing:

- Run command 
    ```
    go build cmd/server/main.go
    ```
    in main directory and it will generate the main binary and run
    ```
    ./main
    ```
- Allow the connections for `./main` binary on 8080
- There are 2 files in `./data` which are environment `Beta` and the collection `employee management`
- Set the environment `Beta` and call the apis, `Create`, `Update`, `Get`, `Delete` and `Get` list api(which has pagination support)

