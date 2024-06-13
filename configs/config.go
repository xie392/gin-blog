package configs

import (
	"blog/models"
	"fmt"
	"github.com/go-redis/redis"
	"gopkg.in/yaml.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB
var RedisClient *redis.Client

type Config struct {
	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"` // Change to int
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Database string `yaml:"database"`
	} `yaml:"database"`
	Redis struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"` // Change to int
		Password string `yaml:"password"`
	} `yaml:"redis"`
	JWTSecret string `yaml:"jwt_secret"`
}

func GetConfig() (*Config, error) {
	yamlFile, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Fatalf("Failed to read YAML file: %v", err)
	}

	var config Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatalf("Failed to unmarshal YAML: %v", err)
	}

	return &config, nil
}

func InitConfig() {
	cfg, err := GetConfig()
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Database,
	)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	if err := models.AutoMigratePostTable(DB); err != nil {
		panic("failed to migrate Admin table")
	}

	if err := models.AutoMigrateAdminTable(DB); err != nil {
		panic("failed to migrate post table")
	}

	if err != nil {
		log.Fatalf("Auto migration failed: %v", err)
	}

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Redis.Host, cfg.Redis.Port),
		Password: cfg.Redis.Password,
		DB:       0,
	})

	_, err = RedisClient.Ping().Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
}
