# jobsity-challenge

## Project Setup Instructions

This project consists of a backend (written in Go) and a frontend (built with Angular). The backend relies on a database that can be easily started using Docker. Follow the instructions below to get the entire project up and running.

Prerequisites

Docker and Docker Compose installed
Go installed (version 1.17 or higher)
Node.js and npm installed (latest LTS version recommended)
Step-by-Step Setup

### 1. Start the Database with Docker
To start the database, run the following command in the project root directory where the docker-compose.yml file is located:

```
docker-compose up -d
```

This will launch the database in a detached mode. You can verify that the database is running by checking the Docker containers:

```
docker ps
```

### 2. Install and Run the Backend (Go)
Navigate to the backend directory and install the necessary Go modules:

```
task-manager-api
go mod tidy
```

Then, start the Go backend server:

```
go run main.go
```

The backend will now be running and accessible at http://localhost:3333.


### 3. Install and Run the Frontend (Angular)
Navigate to the frontend project directory and install the dependencies using npm:

```
cd task-manager
npm install
```

After installing the dependencies, start the frontend server:

```
npm start
```

The frontend will be available at http://localhost:4200.

## Accessing the Application

Backend: The backend API will be running on http://localhost:3333.
Frontend: The frontend UI will be accessible on http://localhost:4200.

### Summary of Commands
```
# Start the database with Docker
docker-compose up -d

# Install backend dependencies
cd task-manager-api
go mod tidy

# Run the Go backend
go run main.go

# Install frontend dependencies
cd task-manager
npm install

# Run the Angular frontend
npm start
```

## Troubleshooting

- If there are issues with Docker, ensure that Docker is properly installed and running.
- Make sure the backend and frontend are started in the correct directories.
- If ports 3333 or 4200 are already in use, you may need to stop other services or change the ports in the respective configurations.