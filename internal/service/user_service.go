package service

import (
	"QtCloudPan/internal/repository"
	"QtCloudPan/pkg/utils"
	"errors"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"time"
)

// RegisterRequest 表示用户注册请求
type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// RegisterResponse 表示用户注册响应
type RegisterResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type LoginResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Token   string `json:"token"`
}

// LoginRequest 表示用户登录请求
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// RegisterUser 处理用户注册逻辑
func RegisterUser(req RegisterRequest) RegisterResponse {
	// 在这里处理用户注册逻辑，比如存储到数据库
	err := repository.SaveUserToDB(req.Username, req.Password, req.Email)
	if err != nil {
		// 对error进行类型断言
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) {
			switch mysqlErr.Number {
			case 1062:
				return RegisterResponse{
					Status:  2, // 2 表示用户名重复
					Message: "Registration failed: " + err.Error(),
				}

			default:
				fmt.Printf("MySQL Error [%d]: %s\n", mysqlErr.Number, mysqlErr.Message)
				return RegisterResponse{
					Status:  0, // 0 表示注册失败
					Message: "Registration failed: " + err.Error(),
				}
			}
		} else {
			// 其他错误类型
			fmt.Println("Failed to execute query:", err)
		}
	}

	return RegisterResponse{
		Status:  1, // 1 表示注册成功
		Message: "Registration successful",
	}
}

func LoginUser(req LoginRequest) LoginResponse {
	// 查询数据库，验证用户登录信息
	err := repository.QueryUserFromDB(req.Username, req.Password)
	if err != nil {
		return LoginResponse{
			Status:  0,
			Message: "Login failed: " + err.Error(),
			Token:   "",
		}
	}

	// 生成JWT
	token, err := utils.GenerateToken(req.Username, 24*7*time.Hour)
	if err != nil {
		return LoginResponse{
			Status:  0,
			Message: "Login failed: " + err.Error(),
			Token:   "",
		}
	}
	// 将token放入响应中
	return LoginResponse{
		Status:  1,
		Message: "Login successful",
		Token:   token,
	}

}
