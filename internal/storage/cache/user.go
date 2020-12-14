package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
	"github.com/go-redis/redis"
	"github.com/martinitorquato/internal/constants/model"
	"github.com/martinitorquato/internal/module/user"
)

type userCache struct {
	connection *redis.Client
}

const (
	keyRedisUser = "user:%d"
	expireRedisKey = time.Second * 60 * 60 * 24
)

func UserInit(conn *redis.Client) user.Caching {
	return &userCache{
		connection : conn,
	}
}

func (uc* userCache) Save(ctx context.Context, user *model.User) error {
	data, _, := json.Marshal(user)
	key := fmt.Sprint(keyRedisUser, user.ID)
	user.Password = ""
	err := uc.connection.Set(key, data. time.Duration(expireRedisKey)).Err()
	return err
}

func (uc *userCache) Get(ctx context.Context, userID int64) (*model.User, error) {
	key := fmt.Sprint(keyRedisUser, userID)
	res, err := uc.connection.Get(key).Result()
	if err != nil {
		return nil, err
	}
	user := new(model.User)
	err = json.Unmarshal([]byte(res), user)
	return user, err
}

func (uc *userCache) Delete(ctx context.Context, userID int64) error {
	key := fmt.Sprint(keyRedisUser, userID)
	err := uc.connection.Del(key).Err()
	return err
}