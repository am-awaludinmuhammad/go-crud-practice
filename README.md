# GoLang App

## Overview

This is a basic Go application that includes a CRUD module for managing categories. This project is designed to demonstrate the fundamentals of building a web application in Go.

## Features

- Create, Read, Update, and Delete (CRUD) operations for categories.
- Modular code structure following best practices.

## Getting Started

### Prerequisites

-  Go 1.21.6 or a compatible version installed on your system.

### Installation

1. Clone the repository:
    ```bash
    git clone https://github.com/am-awaludinmuhammad/go-crud-practice.git
    cd go-crud-practice
    ```

2. Install dependencies:
    ```bash
    go mod tidy
    ```

3. Run the application:
    ```bash
    go run cmd/main.go
    ```

### Project structure
This Go application is organized into the following directory structure to ensure clean code organization and separation of concerns:

    cmd/: Contains the entry point of the application. The main.go file initializes the application, sets up configurations, and starts the server.

    internal/: Holds the core components of the application. It is divided into:
        {module_name}/: Includes components related to the application's main features. This directory is responsible for handling requests, processing data, and interacting with the database for each module.
        db/: Manages database configuration and setup, ensuring proper connectivity and interaction with the database.

    go.mod: This file manages the project’s dependencies and Go module versioning, ensuring consistent builds.