package config

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"simpledouyin/logging"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var EnvConfig = envConfigSchema{}

var dsn string

// 结构体定义，包含数据库和缓存配置
type envConfigSchema struct {
	CONSUL_ADDR     string `env:"CONSUL_ADDR,DREAM_SERVICE_DISCOVERY_URI"`
	EXPORT_ENDPOINT string

	MYSQL_HOST     string
	MYSQL_PORT     string
	MYSQL_USER     string
	MYSQL_PASSWORD string
	MYSQL_DATABASE string

	REDIS_HOST            string
	REDIS_PORT            string
	REDIS_DATABASE        string
	MAX_REQUEST_BODY_SIZE int

	LOCAL_FS_LOCATION string
	LOCAL_FS_BASEURL  string

	S3_ENDPOINT_URL string
	S3_PUBLIC_URL   string
	S3_BUCKET       string
	S3_SECRET_ID    string
	S3_SECRET_KEY   string
	S3_PATH_STYLE   string

	UNSPLASH_ACCESS_KEY string
	STORAGE_TYPE        string

	REDIS_PASSWORD string
	REDIS_DB       int
	REDIS_ADDR     string
}

// 获取数据库连接字符串
func (s *envConfigSchema) GetDSN() string {
	return dsn
}

// 初始化默认配置值和环境变量加载
func init() {
	EnvConfig = defaultConfig

	// 校验配置项
	envValidate()

	// 构建数据库连接字符串
	dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		EnvConfig.MYSQL_USER,
		EnvConfig.MYSQL_PASSWORD,
		EnvConfig.MYSQL_HOST,
		EnvConfig.MYSQL_PORT,
		EnvConfig.MYSQL_DATABASE)
}

// 默认配置，若环境变量没有提供，将使用这些默认值
var defaultConfig = envConfigSchema{
	MYSQL_HOST:     "localhost",
	MYSQL_PORT:     "3306",
	MYSQL_USER:     "root",
	MYSQL_PASSWORD: "Lycdemima1@",
	MYSQL_DATABASE: "douyin",

	REDIS_HOST:     "localhost",
	REDIS_PORT:     "6379",
	REDIS_DATABASE: "0",

	EXPORT_ENDPOINT: "127.0.0.1:4317",
	CONSUL_ADDR:     "127.0.0.1:8500",

	MAX_REQUEST_BODY_SIZE: 200 * 1024 * 1024,

	LOCAL_FS_LOCATION: "/tmp",
	LOCAL_FS_BASEURL:  "http://localhost/",

	STORAGE_TYPE: "s3",

	S3_ENDPOINT_URL: "http://localhost:9000",
	S3_PUBLIC_URL:   "http://localhost:9000",
	S3_BUCKET:       "bucket",
	S3_SECRET_ID:    "minio",
	S3_SECRET_KEY:   "12345678",
	S3_PATH_STYLE:   "true",

	UNSPLASH_ACCESS_KEY: "access_key",
}

func envInit() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, ignored")
	}

	// 使用反射读取 defaultConfig 并加载环境变量
	v := reflect.ValueOf(defaultConfig)
	typeOfV := v.Type()

	for i := 0; i < v.NumField(); i++ {
		fieldName := typeOfV.Field(i).Name
		fieldType := typeOfV.Field(i).Type
		fieldValue := v.Field(i).Interface()

		// 获取字段的环境变量名称（如果有的话）
		envNameAlt := []string{fieldName}
		if fieldTag, ok := typeOfV.Field(i).Tag.Lookup("env"); ok && len(fieldTag) > 0 {
			envNameAlt = append(envNameAlt, strings.Split(fieldTag, ",")...)
		}

		// 根据字段类型进行处理
		switch fieldType {
		case reflect.TypeOf(""): // 字符串类型
			configDefaultValue := fieldValue.(string)
			envValue := resolveEnv(envNameAlt, configDefaultValue)
			if len(envValue) > 0 {
				reflect.ValueOf(&EnvConfig).Elem().Field(i).SetString(envValue)
			}
		case reflect.TypeOf(0): // 整数类型，处理整数值
			configDefaultValue := fieldValue.(int)                                    // 获取默认值时是整数类型
			envValue := resolveEnv(envNameAlt, fmt.Sprintf("%d", configDefaultValue)) // 转为字符串
			if len(envValue) > 0 {
				// 解析为整数
				intValue, err := strconv.Atoi(envValue)
				if err == nil {
					reflect.ValueOf(&EnvConfig).Elem().Field(i).SetInt(int64(intValue))
				} else {
					// 如果转换失败，则使用默认值
					reflect.ValueOf(&EnvConfig).Elem().Field(i).SetInt(int64(configDefaultValue))
				}
			}
		}
	}
}

// 根据多个环境变量名称查找配置值
func resolveEnv(configKeys []string, defaultValue string) string {
	for _, item := range configKeys {
		envValue := os.Getenv(item)
		if envValue != "" {
			return envValue
		}
	}
	return defaultValue
}

// 校验配置项的正确性
func envValidate() {
	if EnvConfig.MYSQL_HOST == "" || EnvConfig.MYSQL_PORT == "" || EnvConfig.MYSQL_USER == "" || EnvConfig.MYSQL_PASSWORD == "" || EnvConfig.MYSQL_DATABASE == "" {
		logging.Logger.WithFields(logrus.Fields{
			"MYSQL_HOST":     EnvConfig.MYSQL_HOST,
			"MYSQL_PORT":     EnvConfig.MYSQL_PORT,
			"MYSQL_USER":     EnvConfig.MYSQL_USER,
			"MYSQL_PASSWORD": EnvConfig.MYSQL_PASSWORD,
			"MYSQL_DATABASE": EnvConfig.MYSQL_DATABASE,
		}).Fatal("Missing essential MySQL configuration values")
	}

	if EnvConfig.REDIS_HOST == "" || EnvConfig.REDIS_PORT == "" {
		logging.Logger.WithFields(logrus.Fields{
			"REDIS_HOST": EnvConfig.REDIS_HOST,
			"REDIS_PORT": EnvConfig.REDIS_PORT,
		}).Fatal("Missing essential Redis configuration values")
	}
}
