# mfc-blogPost Submission
# MFC Posts API
A simple REST API for managing posts, built using Gin and MySQL for Mozilla Firefox Club enrollments 2025 for Backend Subdomain.

### Brief Introduction
This is my first project in Go. I went over the language in 4-5 hours (YouTube, Stack Overflow) and put together this Blog Posting RESTful API with CRUD operations in about 6 hours. This is heavily unsecure and unorganized. The project is a proof of concept. I have worked in PHP before and wanted to take up this challenge. I did not have time to implement the other API task. I have taken the help of a lot of web articles and Golang tutorial and generative AI to help with syntax and bug fixes.

The API allows for basic blog post management, including creating, reading, updating, and deleting posts. It also supports fetching posts by user ID. While functional, the current implementation lacks proper security measures, content validation, and user authentication. Improvements(listed below) could not be implemented partly due to lack of knowledge and mostly due to time constraints.

(p.s I could not install Docker as Docker does not support my computer so please do manual installation !)

## Table of Contents
- [Overview](#overview)
- [Features](#features)
- [Requirements](#requirements)
- [Setup](#setup)
- [API Endpoints](#api-endpoints)
- [Example Use Cases](#example-use-cases)
- [To Be Done](#to-be-done)

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
```env
DB_USER=<user>
DB_PASSWORD=<password>
DB_NAME=<posts>
JWT_SECTET=<key>
```
4. **Run the Application**:
```go
go run main.go
```

## API Endpoints

## API Endpoints

| Endpoint             | Method | Parameters         | Request Body                | Success Response                                  | Error Responses                                                                              |
|----------------------|--------|--------------------|-----------------------------|---------------------------------------------------|----------------------------------------------------------------------------------------------|
| `/posts`             | GET    | None               | None                        | 200 OK - JSON array of posts                      | 500 Internal Server Error - Database error                                                   |
| `/posts/:id`         | GET    | `id` (path parameter, integer) | None                        | 200 OK - JSON object representing the post       | 400 Expectation Failed - ID should be a number<br>404 Not Found - Post not found<br>500 Internal Server Error - Database error  |
| `/posts/user/:user` | GET    | `user` (path parameter, integer) | None                        | 200 OK - JSON array of posts for that user      | 400 Expectation Failed - User should be a number<br>500 Internal Server Error - Database error                                                    |
| `/posts`             | POST   | None               | JSON: `{user: integer, title: string, content: string}` | 201 Created - JSON: `{lastID: integer}` (ID of new post) | 400 Expectation Failed - Invalid arguments in JSON body<br>500 Internal Server Error - Database error  |
| `/posts/:id`         | PUT    | `id` (path parameter, integer) | JSON: `{title: string, content: string}` | 200 OK - Updated post JSON object                   | 400 Expectation Failed - ID should be a number OR Invalid arguments in JSON body<br>500 Internal Server Error - Database error  |
| `/posts/:id`         | DELETE | `id` (path parameter, integer) | None                        | 204 No Content                                  | 400 Expectation Failed - ID should be a number<br>500 Internal Server Error - Database error  |

**Notes:**

*   **Error Responses**: All errors can potentially return a 500 Internal Server Error if there's a database issue. The table lists other more specific errors.
*   **Content-Type**: All `POST` and `PUT` requests should have the `Content-Type` header set to `application/json`.
*   **Validation**: The code lacks explicit input validation. The "Invalid arguments" error typically arises from Gin's `BindJSON` failing, indicating a mismatch between the expected JSON structure and the provided data.
*   **Security:** As highlighted in the file the code is heavily unsecure and lacks authorization or authentication


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

### Creating a Post
```bash
curl -X POST -H "Content-Type: application/json" -d '{"user": 1, "title": "My First Post", "content": "Hello World!"}' http://localhost:8080/posts
```
### Fetching All Posts
```bash
curl http://localhost:8080/posts
```
### Fetching a Specific Post by ID
```bash
curl http://localhost:8080/posts/1
```
### Fetching Posts by User
```bash
curl http://localhost:8080/posts/user/1
```
### Updating a Post
```bash
curl -X PUT -H "Content-Type: application/json" -d '{"title": "Updated Post Title", "content": "This post has been updated!"}' http://localhost:8080/posts/1
```
### Deleting a Post
```bash
curl -X DELETE http://localhost:8080/posts/1
```
## To be Done

- [ ] **Add content validation**  
  Ensure that the `title` and `content` fields in posts are properly validated before being inserted into the database. For example:
  - Parameters should not be empty or exceed a certain character limit.
  - Content should have a minimum length and be checked for invalid inputs (e.g., SQL injection, special characters).

- [ ] **Add login endpoint which validates credentials**  
  Implement a `/login` endpoint that accepts a username and password, validates the credentials against the database, and returns a token (JWT) for authenticated access.

- [ ] **Divide the file into packages**  
  Split the `main.go` file into smaller, more manageable packages. For example:
  - `routes`
  - `controllers`
  - `utils`

- [ ] **Implement user authentication**  
  Add middleware to protect certain endpoints (e.g., creating, updating, or deleting posts).

- [ ] **Implement and Deply using Docker**  
  Deploy using docker with all dependencies and requiremnts in one file.

- [ ] **Implement likes functionality**  
  Add a feature to allow users to "like" posts.
  - Adding endpoints to like/unlike a post.

- [ ] **Performance improvements**  
  - Implementing caching for frequently accessed data like posts, post likes or user sessions.

