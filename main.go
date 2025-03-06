/*
ToDo
-- add content validation
-- add login endpoint
-- divide the file into packages
-- implement user authentication
-- implement likes
-- performance improvements
*/
package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB

type post struct {
	ID         int    `json:"id,omitempty"`
	User       int    `json:"user"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	LastUpdate string `json:"lastupdated,omitempty"`
	Like       int    `json:"likes,omitempty"`
}

func getPosts(c *gin.Context) {
	var posts []post

	rows, err := db.Query("SELECT id, user, title, content, lastupdate, likes FROM posts")
	if err != nil {
		log.Printf("Database error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var pst post
		if err := rows.Scan(&pst.ID, &pst.User, &pst.Title, &pst.Content, &pst.LastUpdate, &pst.Like); err != nil {
			log.Printf("Database error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
			return
		}
		posts = append(posts, pst)
	}
	if err := rows.Err(); err != nil {
		log.Printf("Database error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	c.JSON(http.StatusOK, posts)
}

func postPosts(c *gin.Context) {
	var newPost post

	if err := c.BindJSON(&newPost); err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{"error": "Invalid arguments"})
		return
	}
	result, err := db.Exec("INSERT INTO posts (user, title, content) VALUES (?, ?, ?)", newPost.User, newPost.Title, newPost.Content)
	if err != nil {
		log.Printf("Database error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Printf("Database error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"lastID": id})
}

func getPostsByID(c *gin.Context) {
	id := c.Param("id")
	val, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{"error": "ID should be a number"})
		return
	}
	var pst post

	err = db.QueryRow("SELECT id, user, title, content, lastupdate, likes FROM posts WHERE id = ?", val).Scan(&pst.ID, &pst.User, &pst.Title, &pst.Content, &pst.LastUpdate, &pst.Like)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		} else {
			log.Printf("Database error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		}
		return
	}
	c.JSON(http.StatusOK, pst)
}

func deletePostsById(c *gin.Context) {
	id := c.Param("id")
	val, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{"error": "ID should be a number"})
		return
	}
	_, err = db.Exec("DELETE FROM posts WHERE id = ?", val)
	if err != nil {
		log.Printf("Database error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	c.Status(http.StatusNoContent)
}

func editPostsById(c *gin.Context) {
	id := c.Param("id")
	val, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{"error": "ID should be a number"})
		return
	}

	var newPost post
	if err = c.BindJSON(&newPost); err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{"error": "Invalid arguments"})
		return
	}

	currentTime := time.Now().Format("2006-01-02 15:04:05")
	_, err = db.Exec("UPDATE posts SET title = ?, content = ?, lastupdate = ? WHERE id = ?", newPost.Title, newPost.Content, currentTime, val)
	if err != nil {
		log.Printf("Database error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	c.JSON(http.StatusOK, newPost)
}

func getPostsByUser(c *gin.Context) {
	user := c.Param("user")
	val, err := strconv.Atoi(user)
	if err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{"error": "User should be a number"})
		return
	}
	var posts []post

	rows, err := db.Query("SELECT id, user, title, content, lastupdate, likes FROM posts WHERE user = ?", val)
	if err != nil {
		log.Printf("Database error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var pst post
		if err := rows.Scan(&pst.ID, &pst.User, &pst.Title, &pst.Content, &pst.LastUpdate, &pst.Like); err != nil {
			log.Printf("Database error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
			return
		}
		posts = append(posts, pst)
	}
	if err := rows.Err(); err != nil {
		log.Printf("Database error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	c.JSON(http.StatusOK, posts)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := mysql.Config{
		User:                 os.Getenv("DB_USER"),
		Passwd:               os.Getenv("DB_PASSWORD"),
		DBName:               os.Getenv("DB_NAME"),
		AllowNativePasswords: true,
	}
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	router := gin.Default()
	router.GET("/posts", getPosts)
	router.GET("/posts/:id", getPostsByID)
	router.GET("/posts/user/:user", getPostsByUser)
	router.POST("/posts", postPosts)
	router.PUT("/posts/:id", editPostsById)
	router.DELETE("/posts/:id", deletePostsById)
	router.Run("localhost:8080")
}
