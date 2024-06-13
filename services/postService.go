package services

import (
	"blog/configs"
	"blog/models"
	"context"
	"encoding/json"
	"fmt"
	"time"
)

func GetPost(id int) (models.Post, error) {
	var post models.Post
	ctx := context.Background()
	key := fmt.Sprintf("post:%d", id)
	result, err := configs.RedisClient.WithContext(ctx).Get(key).Result()
	if err == nil {
		// 如果在 Redis 中找到了数据，将其反序列化为 Post 对象并返回
		err := json.Unmarshal([]byte(result), &post)
		if err != nil {
			return models.Post{}, err
		}
		return post, nil
	}

	// 如果 Redis 中没有找到数据，则从数据库中获取，并存入 Redis
	err = configs.DB.First(&post, id).Error
	if err != nil {
		return models.Post{}, err
	}

	// 将获取到的 post 对象序列化为 JSON，并存入 Redis
	postData, err := json.Marshal(post)
	if err != nil {
		return models.Post{}, err
	}
	err = configs.RedisClient.WithContext(ctx).Set(key, postData, 24*time.Hour).Err()
	if err != nil {
		fmt.Println("Error caching post in Redis:", err)
	}

	return post, nil
}

func CreatePost(post models.Post) (models.Post, error) {
	ctx := context.Background()

	err := configs.DB.Create(&post).Error
	if err != nil {
		fmt.Println("post service create post error", err)
		return models.Post{}, err
	}

	// 创建成功后，将新创建的 post 对象存入 Redis
	key := fmt.Sprintf("post:%d", post.ID)
	postData, err := json.Marshal(post)
	if err != nil {
		return models.Post{}, err
	}
	err = configs.RedisClient.WithContext(ctx).Set(key, postData, 24*time.Hour).Err()
	if err != nil {
		fmt.Println("Error caching post in Redis:", err)
	}

	return post, nil
}

func GetAllPosts() ([]models.Post, error) {
	var posts []models.Post
	ctx := context.Background()

	// 尝试从 Redis 中获取数据
	result, err := configs.RedisClient.WithContext(ctx).Get("all_posts").Result()
	if err == nil {
		// 如果在 Redis 中找到了数据，将其反序列化为 Post 对象并返回
		err := json.Unmarshal([]byte(result), &posts)
		if err != nil {
			return []models.Post{}, err
		}
		return posts, nil
	}

	// 如果 Redis 中没有找到数据，则从数据库中获取，并存入 Redis
	err = configs.DB.Find(&posts).Error
	if err != nil {
		return []models.Post{}, err
	}

	// 将获取到的 posts 对象序列化为 JSON，并存入 Redis
	postData, err := json.Marshal(posts)
	if err != nil {
		return []models.Post{}, err
	}
	err = configs.RedisClient.WithContext(ctx).Set("all_posts", postData, 24*time.Hour).Err()
	if err != nil {
		fmt.Println("Error caching all posts in Redis:", err)
	}

	return posts, nil
}

func UpdatePost(post models.Post) (models.Post, error) {
	ctx := context.Background()

	err := configs.DB.Save(&post).Error
	if err != nil {
		return models.Post{}, err
	}

	// 更新 Redis 中的数据
	key := fmt.Sprintf("post:%d", post.ID)
	postData, err := json.Marshal(post)
	if err != nil {
		return models.Post{}, err
	}
	err = configs.RedisClient.WithContext(ctx).Set(key, postData, 24*time.Hour).Err()
	if err != nil {
		fmt.Println("Error updating post in Redis:", err)
	}

	return post, nil
}

func DeletePost(id int) error {
	ctx := context.Background()

	err := configs.DB.Delete(&models.Post{}, id).Error
	if err != nil {
		return err
	}

	// 删除 Redis 中的数据
	key := fmt.Sprintf("post:%d", id)
	err = configs.RedisClient.WithContext(ctx).Del(key).Err()
	if err != nil {
		fmt.Println("Error deleting post from Redis:", err)
	}

	return nil
}
