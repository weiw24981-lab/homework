package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	dbfactory "practise/dbfactory"
	model "practise/model"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var db *gorm.DB

var config model.Config

func init() {
	// 设置配置文件名称（不含扩展名）
	viper.SetConfigName("config")
	// 设置配置文件类型
	viper.SetConfigType("yaml")
	// 添加配置文件搜索路径
	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME/.app")

	// 读取环境变量
	viper.AutomaticEnv()
	viper.SetEnvPrefix("APP")

	// 设置默认值
	viper.SetDefault("server.port", "8080")
	viper.SetDefault("server.host", "0.0.0.0")
	viper.SetDefault("server.mode", "debug")

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Warning: Error reading config file: %v", err)
		log.Println("Using default values and environment variables")
	} else {
		log.Printf("Using config file: %s", viper.ConfigFileUsed())
	}
}

// LoadConfig 加载并解析配置文件，将 Viper 中的配置数据映射到 Config 结构体
// 返回解析后的配置对象指针和可能的错误
// 配置数据来源包括：配置文件、环境变量和默认值（在 init 函数中已设置）
func LoadConfig() (*model.Config, error) {
	// 使用 viper.Unmarshal 将配置数据解析到 config 结构体中
	// 如果解析失败，返回错误
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	// 返回解析成功的配置对象
	return &config, nil
}

// ========== JWT Claims ==========
type Claims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func main() {
	config, err := LoadConfig()

	db := dbfactory.NewTestDB(config)
	if err := db.AutoMigrate(&model.User{}, &model.Post{}, &model.Comment{}, &model.Tag{}); err != nil {
		log.Fatalf("auto migrate: %v", err)
	}

	// var us model.User
	// if err := db.Where(" name = ? and password = ?", "admin", "admin123").First(&us).Error; err != nil {
	// 	log.Fatalf("test db: %v", err)
	// }

	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// 设置 Gin 模式
	gin.SetMode(config.Server.Mode)
	r := gin.Default()

	// ========== 公开路由 ==========
	r.POST("/api/login", login)
	r.GET("/api/public", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "public endpoint"})
	})

	// ========== 需要认证的路由 ==========
	api := r.Group("/api")
	api.Use(authMiddleware())
	{
		api.POST("/tags", createTag)
		api.PUT("/tag/:id", updateTag)
		api.GET("/tag/:id", getTag)
		api.GET("/tags", getTags)
		api.POST("/posts", createPost)

		api.GET("/protected", func(c *gin.Context) {
			userID, _ := c.Get("userID")
			username, _ := c.Get("username")
			c.JSON(http.StatusOK, gin.H{
				"message":  "protected endpoint",
				"user_id":  userID,
				"username": username,
			})
		})

		api.GET("/profile", getProfile)
	}

	addr := fmt.Sprintf("%s:%s", config.Server.Host, config.Server.Port)
	log.Printf("Server starting on %s", addr)
	r.Run(addr)
}

// ========== 生成 Token ==========
func generateToken(userID uint, username string) (string, error) {
	ex := config.JWT.Expire
	claims := Claims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(ex) * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(config.JWT.Secret)
}

// ========== 解析 Token ==========
func parseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return config.JWT.Secret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

// ========== 认证中间件 ==========
func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从 Header 获取 Token
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization header required",
			})
			c.Abort()
			return
		}

		// 提取 Token（Bearer <token>）
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid authorization header format",
			})
			c.Abort()
			return
		}

		tokenString := parts[1]

		// 验证 Token
		claims, err := parseToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid token",
			})
			c.Abort()
			return
		}

		// 将用户信息存储到 Context
		c.Set("userID", claims.UserID)
		c.Set("username", claims.Username)

		c.Next()
	}
}

// ========== tag接口 ==========
func createTag(c *gin.Context) {
	var req struct {
		Name string `json:"name" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var tag = model.Tag{}
	tag.Name = req.Name
	if err := db.Create(&tag).Error; err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	success(c, gin.H{
		"tag": tag,
	})
}

func updateTag(c *gin.Context) {
	var req struct {
		Name string `json:"name" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	var tag = model.Tag{}
	tag.Name = req.Name
	tag.ID = uint(id)
	if err := db.Model(&tag).Updates(&tag).Error; err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	success(c, gin.H{
		"tag": tag,
	})
}

func getTag(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	var tag = model.Tag{}
	tag.ID = uint(id)
	if err := db.Model(&tag).First(&tag).Error; err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	success(c, gin.H{
		"tag": tag,
	})
}

func getTags(c *gin.Context) {
	var tags []model.Tag
	if err := db.Find(&tags).Error; err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	success(c, gin.H{
		"tags": tags,
	})
}

// ========== post接口 ==========
func createPost(c *gin.Context) {
	var req struct {
		Title   string `json:"title" binding:"required"`
		Content string `json:"content" binding:"required"`
		Tags    []int  `json:"tags"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var post = model.Post{}
	var tags []model.Tag
	post.Title = req.Title
	post.Content = req.Content
	tagids := req.Tags
	if tagids != nil {
		for _, tagid := range tagids {
			tag := model.Tag{}
			tag.ID = uint(tagid)
			tags = append(tags, tag)
		}
		post.Tags = tags
	}
	if err := db.Create(&post).Error; err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	success(c, gin.H{
		"post": post,
	})
}

func updatePost(c *gin.Context) {
	var req struct {
		Title   string `json:"title" binding:"required"`
		Content string `json:"content" binding:"required"`
		Tags    []int  `json:"tags"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	var post = model.Post{}
	var tags []model.Tag
	post.Title = req.Title
	post.Content = req.Content
	post.ID = uint(id)
	tagids := req.Tags
	if tagids != nil {
		for _, tagid := range tagids {
			tag := model.Tag{}
			tag.ID = uint(tagid)
			tags = append(tags, tag)
		}
		post.Tags = tags
	}

	if err := db.Model(&post).Updates(&post).Error; err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	success(c, gin.H{
		"post": post,
	})
}

func getPost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	var post = model.Post{}
	post.ID = uint(id)
	if err := db.Model(&post).First(&post).Error; err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	success(c, gin.H{
		"post": post,
	})
}

func getPosts(c *gin.Context) {
	var posts []model.Post
	if err := db.Find(&posts).Error; err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	success(c, gin.H{
		"posts": posts,
	})
}

// ========== 登录接口 ==========
func login(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// 验证用户名密码（示例，实际应从数据库查询）
	var us model.User
	if err := db.Where(" name = ? and password = ?", req.Username, req.Password).First(&us).Error; err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := generateToken(us.ID, us.Name)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	success(c, gin.H{
		"token": token,
		"user": gin.H{
			"id":       us.ID,
			"username": us.Name,
		},
	})

	// c.JSON(http.StatusUnauthorized, gin.H{
	// 	"error": "Invalid credentials",
	// })
}

// ========== 获取用户信息 ==========
func getProfile(c *gin.Context) {
	userID, _ := c.Get("userID")
	username, _ := c.Get("username")

	c.JSON(http.StatusOK, gin.H{
		"id":       userID,
		"username": username,
	})
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "success",
		Data:    data,
	})
}

func errorResponse(c *gin.Context, code int, message string) {
	c.JSON(code, Response{
		Code:    code,
		Message: message,
	})
}
