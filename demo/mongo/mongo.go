package mongo

import (
	"context"
	"github.com/Allen9012/broker"
	mongo_broker "github.com/Allen9012/broker/mongo"
	"sync"
)

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/8/30
  @desc:
  @modified by:
**/

var (
	Client        *mongo_broker.Client
	onceInitMongo sync.Once
)

func init() {
	onceInitMongo.Do(func() {
		ctx := context.Background()
		tc := &mongo_broker.Client{
			BaseComponent: broker.NewBaseComponent(),
			RealCli: mongo_broker.NewClient(ctx, &mongo_broker.Config{
				URI:         "mongodb://localhost:27017",
				MinPoolSize: 3,
				MaxPoolSize: 3000,
			}),
		}

		tc.Launch()
		defer tc.Stop()
	})
}
