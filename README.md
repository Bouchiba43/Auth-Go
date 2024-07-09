# Building a Go Application with CompileDaemon, Docker, and Docker Compose

## Introduction

This guide explains how to set up and build a Go application using CompileDaemon for local development and Docker/Docker Compose for containerized deployment.

### Prerequisites

1. **Go Installed**: Ensure Go is installed on your system. Verify your installation by running:
    ```sh
    go version
    ```

2. **CompileDaemon Installed**: Install CompileDaemon using:
    ```sh
    go install github.com/githubnemo/CompileDaemon@latest
    ```

3. **Docker Installed**: Ensure Docker is installed on your system. Check your installation with:
    ```sh
    docker --version
    ```

4. **Docker Compose Installed** (for Docker Compose only): Check Docker Compose installation:
    ```sh
    docker-compose --version
    ```

5. **Setting Environment Variables**:
    - Create a `.env` file in your project directory and set the following variables:
      ```dotenv
      PORT=1323
      # PostgreSQL connection variables
      DB="host=# user=# password=# dbname=# port=# sslmode=#"
      JWT_SECRET=
      ```

    - Replace `host=#`, `user=#`, `password=#`, `dbname=#`, `port=#`, and `sslmode=#` with your PostgreSQL connection details.

    - Set `JWT_SECRET` using a Linux command:
      ```sh
      openssl rand -base64 32
      ```

## Steps to Build and Run the Go Application

### Local Development with CompileDaemon

1. **Navigate to Your Project Directory**: Open your terminal and change to your Go project directory.
    ```sh
    cd /path/to/your/project
    ```

2. **Organize Your Project Structure**:
    Ensure your `main.go` file is located in the `cmd/app` directory (or adjust paths accordingly).

3. **Run CompileDaemon**:
    Start CompileDaemon to automatically build and restart your server on file changes.
    ```sh
    CompileDaemon --build="go build -o main ./cmd/app/main.go" --command="./main"
    ```

    - `--build="go build -o main ./cmd/app/main.go"`: Specifies the build command to compile `main.go` into an executable named `main`.
    - `--command="./main"`: Defines the command to run the compiled executable.

4. **Verify Output**:
    Upon running CompileDaemon, you should see output indicating successful builds and server restarts on file changes.

### Containerized Deployment with Docker and Docker Compose

#### Docker Setup

1. **Ensure Dockerfile**:
    Confirm you have a `Dockerfile` in your project directory 

2. **Build Docker Image**:
    Build the Docker image using the existing `Dockerfile`:
    ```sh
    docker build -t goland-app:latest .
    ```

#### Running Docker Container

1. **Run Docker Container**:
    Start a Docker container based on the built image:
    ```sh
    docker run -d --name auth-go -p 1323:1323 goland-app:latest
    ```

    - `-d`: Runs the container in detached mode.
    - `--name auth-go`: Assigns the name `auth-go` to the container.
    - `-p 1323:1323`: Maps port 8080 from the container to your host machine.

### Using Docker Compose

1. **Ensure `docker-compose.yml`**:
    Confirm you have a `docker-compose.yml` file in your project directory 

2. **Build and Start Docker Compose**:
    Build and start your application using Docker Compose:
    ```sh
    docker-compose up --build
    ```

    This command builds the Docker image defined in your `Dockerfile`, starts the container, and exposes port 1323.

## Conclusion

Integrating CompileDaemon for local development and Docker/Docker Compose for containerized deployment simplifies the build and deployment process of your Go applications. It enhances development efficiency and ensures consistency across different environments.
