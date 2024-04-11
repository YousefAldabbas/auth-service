# Auth Service

## Description
The Auth Service is a GoLang-based microservice that handles user authentication and authorization for your application. It provides secure access control and user management functionalities.

## Folder Directory
- bin: Contains compiled executable files, including the server executable.
- cmd: Contains application-specific code for different executables.
    - server:
        - main.go: Entry point of the server application.
        - routes.go: Defines the routes for the server.
- data: Used to store data files used by the application, such as database files and static data.
- internal: Contains code internal to the application.
    - config:
        - database.go: Sets up the database connection.
        - env.go: Handles environment variables.
- Makefile: Makefile for building, installing, testing, and deploying the application.
- pkg: Contains libraries and packages used by other applications.
    - handler: Contains HTTP handlers for the application.
    - middleware: Contains middleware functions for the application.
    - migrations: Contains database migration scripts.
    - model: Contains data models for the application.
    - repository: Contains code for interacting with the database.
    - utils: Contains utility functions used throughout the application.
- profiles: Contains configuration profiles for different environments (e.g., development, testing, production).
- scripts: Contains scripts for tasks like building, installing, testing, and deploying the application.
- tests: Contains test code for the application.
