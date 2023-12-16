# Go Todo List REST API
## Overview
This Go Todo List REST API provides a simple and efficient way to manage a todo list by modifying JSON data. The API is built with the Gin web framework and allows for operations like retrieving, adding, and updating todo items.

## Features
List Todos: Retrieve a list of all todo items.
Get Todo: Fetch a single todo item by its ID.
Add Todo: Add a new todo item to the list.
Toggle Todo Status: Toggle the completion status of a todo item.
Getting Started
Prerequisites
Go (Golang) installed on your machine.
Basic understanding of REST APIs and HTTP methods.
Installation
Clone the repository to your local machine:

```bash
git clone [repository-url]
```

Navigate to the project directory:


``` bash
cd [project-directory]
```

Running the API
To start the server, run:

```bash
go run main.go
```

The server will start on localhost:9090.

## API Endpoints
1. List Todos
Endpoint: /todos
Method: GET
Description: Fetches all todo items.

2. Get Todo
Endpoint: /todos/:id
Method: GET
Description: Fetches a single todo item by ID.

3. Add Todo
Endpoint: /todos
Method: POST
Description: Adds a new todo item. The item should be passed in the request body in JSON format.

4. Toggle Todo Status
Endpoint: /todos/:id
Method: PATCH
Description: Toggles the completion status of a todo item.


## Examples

### Add Todo
Request:

```json
POST /todos
Content-Type: application/json

{
  "id": "4",
  "item": "Buy Groceries",
  "completed": false
}
```

Response:

```json
Status: 201 Created

{
  "id": "4",
  "item": "Buy Groceries",
  "completed": false
}
```

### Toggle Todo Status
Request:

```json
PATCH /todos/1
```
Response:

```json
Status: 200 OK

{
  "id": "1",
  "item": "Clean Room",
  "completed": true
}
```
Contributing
Contributions to improve the API are welcome. Please follow the standard pull request process to propose changes.

License
This project is licensed under the MIT License.
