package worker

import (
	"context"

	"github.com/hibiken/asynq"
	db "github.com/piriya-muaithaisong/Simple-Bank/db/sqlc"
)

type TaskProcessor interface {
	ProcessTaskSendVerifyEmail(ctx context.Context, task asynq.Task) error
}

type RedisTaskProcessor struct {
	server *asynq.Server
	store  db.Store
}

func NewRedisTaskProcessor(redisOpt asynq.RedisClientOpt, store db.Store) TaskProcessor {
	server := asynq.NewServer(
		redisOpt,
		asynq.Config{},
	)

	return &RedisTaskProcessor{
		server: server,
		store:  store,
	}
}
