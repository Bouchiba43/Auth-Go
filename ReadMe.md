# Running the Go Server with CompileDaemon

## Introduction

This guide explains how to use CompileDaemon to automatically build and run your Go server. CompileDaemon is a tool that watches for changes in your Go files, rebuilds the project, and restarts the server automatically.

## Prerequisites

1. **Go Installed**: Ensure you have Go installed on your system. You can check your installation by running:
    ```sh
    go version
    ```
2. **CompileDaemon Installed**: You can install CompileDaemon using the following command:
    ```sh
    go install github.com/githubnemo/CompileDaemon@latest
    ```

## Steps to Run the Server

1. **Navigate to Your Project Directory**: Open your terminal and change to the directory where your Go project is located.
    ```sh
    cd /path/to/your/project
    ```

2. **Organize Your Project Structure**:
    Ensure your `main.go` is located in the `cmd/app` directory (or adjust the path in the commands below to match your structure).

3. **Run CompileDaemon**:
    Use the following command to start CompileDaemon. It will watch your files for changes, build the project, and restart the server automatically.
    ```sh
    CompileDaemon --build="go build -o main ./cmd/app/main.go" --command="./main"
    ```

    - `--build="go build -o main ./cmd/app/main.go"`: This flag specifies the build command. It compiles the Go source file into an executable named `main`.
    - `--command="./main"`: This flag specifies the command to run the compiled executable.

4. **Example Output**:
    After running the command, you should see output similar to this:
    ```sh
    2024/07/08 13:22:59 Running build command!
    2024/07/08 13:22:59 Build ok.
    2024/07/08 13:22:59 Restarting the given command.
    2024/07/08 13:22:59 Command is running.
    ```

## Additional Tips

- **Check Permissions**: Ensure your Go source files and directories have the correct permissions to avoid permission errors.
    ```sh
    chmod +x ./cmd/app/main.go
    ```
- **Verbose Mode**: Run CompileDaemon in verbose mode to get more detailed output.
    ```sh
    CompileDaemon --build="go build -o main ./cmd/app/main.go" --command="./main" --verbose
    ```

- **Environment Variables**: You can set environment variables directly in the command if needed.
    ```sh
    CompileDaemon --build="go build -o main ./cmd/app/main.go" --command="./main" --env="ENV_VAR=value"
    ```

- **Configuration File**: Consider using a configuration file for more complex setups and to keep your command line clean.

## Conclusion

Using CompileDaemon can significantly streamline your development process by automating the build and run cycle of your Go application. It watches for changes in your source files, rebuilds the project, and restarts the server, allowing you to focus on writing code without manually restarting the server.
