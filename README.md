# Go Template Project

This is a template project for Go applications. It provides a basic structure for your Go application and includes a number of features to help you get started quickly.

## Project Structure

The project is structured as follows:

- `api/`: Contains the OpenAPI (Swagger) specification for your application.
- `assets/`: Static files that your application might need.
- `build/`: Contains build-related files.
- `cmd/`: The main entry point for your application.
- `configs/`: Configuration files for your application.
- `deployments/`: Kubernetes deployment files.
- `internal/`: The main application code.
- `pkg/`: Libraries and packages that can be used by other services.

## Getting Started

To get started with this project, you need to have Go installed on your machine. Once you have Go installed, you can clone this repository and start developing your application.

```sh
git clone https://github.com/dkdkv/go_template.git
cd go_template
go run cmd/main/main.go
