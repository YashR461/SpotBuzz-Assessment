# SpotBuzz-Assessment
SpotBuzz Assessment

# Features
- Create a new player entry.
- Update player attributes (name and score).
- Delete a player entry.
- Retrieve a list of all players in descending order.
- Fetch a player by rank.
- Retrieve a random player.

# Prerequisites
- Go 1.19 installed on your machine
- Docker (optional).

# Installation
    1. Clone the repository to your local machine:
        git clone https://github.com/YashR461/SpotBuzz-Assessment.git

    2. Navigate to the project directory:
        cd SpotBuzz-Assessment
    3. Install the required Go packages:
        go mod download

# Directory Structure
SpotBuzz-Assessment
├── persistence
    └── mysql.go
    └── init.sql
├── service
    └── playerService.go
├── src
    └── .gitignore
    └── model
        └──player.go
├── main.go
├── go.mod
├── go.sum
├── Dockerfile
├── README.md

# Configuration
Edit the database connection settings in persistence/mysql.go

# Build & Run 
go mod tidy
Run the application using the following command: go run main.go
Or, you can use Docker to run the application in a container:
    docker build -t player-score-management . 
    docker run -p 8080:8080 player-score-management
The application will start and listen on localhost:8080.

Use tools like  Postman to send HTTP requests to the provided endpoints (e.g., POST, PUT, DELETE, GET) to interact with the Player Score Management System.

# How to test Running cloud instance
Service is up and running on Amazon EC2 ubuntu instance at http://ec2-13-51-85-11.eu-north-1.compute.amazonaws.com:3000/

# API Endpoints
1. POST/players – Createsanewentryforaplayer
2. PUT /players/:id – Updates the player attributes. Only name and score can be updated
3. DELETE /players/:id – Deletes the player entry
4. GET /players – Displays the list of all players in descending order
5. GET /players/rank/:val – Fetches the player ranked “val”
6. GET /players/random – Fetches a random player

| HTTP Method |                 Endpoint                                                         |      Description       |
|-------------|----------------------------------------------------------------------------------|------------------------|
| GET         | http://ec2-13-51-85-11.eu-north-1.compute.amazonaws.com:3000/players             | Get all players        |
| GET         | http://ec2-13-51-85-11.eu-north-1.compute.amazonaws.com:3000/players/rank/{rank} | Get player by rank     |
| DELETE      | http://ec2-13-51-85-11.eu-north-1.compute.amazonaws.com:3000/players/{id}        | Deletes a player by id |
| PUT         | http://ec2-13-51-85-11.eu-north-1.compute.amazonaws.com:3000/players/{id}        | Update player by id    |
| GET         | http://ec2-13-51-85-11.eu-north-1.compute.amazonaws.com:3000/players/random      | Get a random player    |
| POST        | http://ec2-13-51-85-11.eu-north-1.compute.amazonaws.com:3000/players             | Adds a player          |
|-------------|----------------------------------------------------------------------------------|------------------------|
