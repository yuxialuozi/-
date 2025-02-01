package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"google.golang.org/grpc"
	"log"
	"net"
	"simpledouyin/constants" // 导入数据库常量配置
	"simpledouyin/generated/douyin/auth"
	server "simpledouyin/service"
)

// 启动 gRPC 服务器
func startGRPCServer() {
	// 数据库连接字符串
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s",
		constants.DBUser,
		constants.DBPassword,
		constants.DBHost,
		constants.DBPort,
		constants.DBName,
		constants.DBCharset,
		constants.DBParseTime,
		constants.DBLoc,
	)

	// 连接到数据库
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	// 启动 gRPC 服务器
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	auth.RegisterAuthServiceServer(s, &server.AuthServer{DB: db}) // 注册 gRPC 服务

	log.Println("gRPC Server listening on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// 启动 Gin HTTP 服务器
func startHTTPServer() {
	// 创建 Gin 路由
	r := gin.Default()

	// 创建 gRPC 客户端连接
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := auth.NewAuthServiceClient(conn)

	// 登录接口
	r.POST("/login", func(c *gin.Context) {
		req := &auth.LoginRequest{Username: "test", Password: "password123"}
		resp, err := client.Login(c, req)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, resp)
	})

	// 启动 Gin HTTP 服务器
	r.Run(":8080")
}

func main() {
	// 启动 gRPC 和 HTTP 服务器
	go startGRPCServer()
	startHTTPServer()
}
