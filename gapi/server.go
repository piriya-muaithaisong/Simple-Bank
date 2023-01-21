package gapi

import (
	"fmt"

	db "github.com/piriya-muaithaisong/Simple-Bank/db/sqlc"
	"github.com/piriya-muaithaisong/Simple-Bank/pb"
	"github.com/piriya-muaithaisong/Simple-Bank/token"
	"github.com/piriya-muaithaisong/Simple-Bank/utils"
)

type Server struct {
	pb.UnimplementedSimpleBankServer
	config     utils.Config
	store      db.Store
	tokenMaker token.Maker
}

func NewServer(config utils.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)

	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		store:      store,
		tokenMaker: tokenMaker,
		config:     config,
	}

	return server, nil
}
