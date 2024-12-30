package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 文章结构体
type Post struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// 模拟的文章数据
var posts = []Post{
	{ID: 1, Title: "Go语言入门", Content: "Go语言是一门简洁的编程语言..."},
	{ID: 2, Title: "Gin框架教程", Content: "Gin是一个高性能的Go语言Web框架..."},
}

func main() {
	// 创建Gin引擎
	r := gin.Default()

	// 首页路由，显示欢迎消息
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "欢迎来到我的博客!",
		})
	})

	// 获取所有文章
	r.GET("/posts", func(c *gin.Context) {
		c.JSON(http.StatusOK, posts)
	})

	// 获取单篇文章
	r.GET("/posts/:id", func(c *gin.Context) {
		id := c.Param("id")
		for _, post := range posts {
			if string(post.ID) == id {
				c.JSON(http.StatusOK, post)
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"message": "文章未找到"})
	})

	// 启动服务器
	r.Run(":8680")
}
