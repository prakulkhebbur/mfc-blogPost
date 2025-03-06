# mfc-blogPost
# MFC Posts API
A simple REST API for managing posts, built using Gin and MySQL.

## Table of Contents
- [Overview](#overview)
- [Features](#features)
- [Requirements](#requirements)
- [Setup](#setup)
- [API Endpoints](#api-endpoints)
- [Example Use Cases](#example-use-cases)

## Overview
This project provides a basic API for creating, reading, updating, and deleting (CRUD) posts. It uses Gin as the web framework and MySQL for database operations.

## Features
- **CRUD Operations**: Supports creating, reading, updating, and deleting posts.
- **User-specific Posts**: Allows fetching posts by user ID.
- **Error Handling**: Includes robust error handling for database operations and invalid requests.

## Requirements
- **Go**: Version 1.18 or later.
- **Gin**: For building the web API.
- **MySQL**: For database operations.
- **MySQL Driver**: The `github.com/go-sql-driver/mysql` package.

## Setup
1. **Install Go**: Ensure Go is installed on your system.
2. **Get Required Packages**:
```go
go get github.com/gin-gonic/gin
```
```go
go get github.com/go-sql-driver/mysql
```
3. **Configure MySQL**:
- Create a MySQL database named `mfc-posts` (use the SQL provided to import a sample DB with 3 posts).
- Update the `main.go` file with your MySQL credentials (create a .env file).
4. **Run the Application**:
```go
go run main.go
```

## API Endpoints
### GET /posts
- **Description**: Fetch all posts.
- **Response**: JSON array of posts.

### GET /posts/:id
- **Description**: Fetch a post by ID.
- **Response**: JSON object representing the post.

### GET /posts/user/:user
- **Description**: Fetch posts by user ID.
- **Response**: JSON array of posts.

### POST /posts
- **Description**: Create a new post.
- **Request Body**: JSON object with `user`, `title`, and `content`.
- **Response**: JSON object with the ID of the newly created post.

### PUT /posts/:id
- **Description**: Update an existing post.
- **Request Body**: JSON object with updated `title` and `content`.
- **Response**: Success message.

### DELETE /posts/:id
- **Description**: Delete a post by ID.
- **Response**: Success message.

## Example Use Cases
- **Creating a Post**:
```bash
curl -X POST -H "Content-Type: application/json" -d '{"user": 1, "title": "My First Post", "content": "Hello World!"}' http://localhost:8080/posts
```
- **Fetching Posts by User**:
```bash
curl http://localhost:8080/posts/user/1
```


