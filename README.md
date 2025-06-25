# todo-api
A simple Todo APIs in Go with implementation of CRUD based REST APIs and JWT Middleware 

Implementation of below package

1. godotenv
2. golang-jwt
3. net/http
4. gorilla/mux
5. gorm
6. driver/postgres

Make a .env file with below configuration for working project on your local

DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=admin
DB_NAME=tododb
JWT_SECRET=my_secret_key

REST APIs endpoints are as below 
1. Generate JWT Token | POST | (http://localhost:8000/login) | use generated token for further API request in Authorization Bearer Header
2. Read All Todos | GET | (http://localhost:8000/api/todos)
3. Read Specific Todo | GET | (http://localhost:8000/api/todo/:id)
4. Create Todo | POST | (http://localhost:8000/api/todos)
5. Update Todo | PUT | (http://localhost:8000/api/todos/:id)
6. Delete Todo | DELETE | (http://localhost:8000/api/todos/:id)