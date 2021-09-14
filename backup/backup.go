package backup

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
	"strings"
)

type backupClient struct {
	client *redis.Client
}

type Backup interface {
	SubscribeToExpiredEvents(ctx context.Context)
}

func NewBackup(client *redis.Client) Backup {
	return &backupClient{client}
}

func (b *backupClient) SubscribeToExpiredEvents(ctx context.Context) {
	keyspace := "__keyevent@0__:expired"
	log.Println("Subscribing to expired events at '", keyspace , "'")
	pubSub := b.client.PSubscribe(ctx, keyspace)
	defer pubSub.Close()
	_, err := pubSub.Receive(ctx)
	if err != nil {
		panic(err)
	}

	receiveChannel := pubSub.Channel()

	for msg := range receiveChannel {
		log.Println("Received expired key: ", msg.Payload)
		err = processExpiredKey(b, ctx, msg.Payload)
		if err != nil {
			log.Println("Error processing expired key", err)
		}
	}
}

func processExpiredKey(b *backupClient, ctx context.Context, expiredKey string) error {
	// get actual key of person object
	key := strings.TrimSuffix(expiredKey,"_expire")
	// fetch person value from REDIS
	person, err := b.client.Get(ctx, key).Result()
	if err != nil {
		return err
	}
	// backup value to S3  using AWS SDK
	log.Println("Backing-up to S3 value:", person)
	// delete key
	return b.client.Del(ctx, key).Err()
}