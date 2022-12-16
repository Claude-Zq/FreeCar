package global

import (
	"github.com/CyanAsterisk/FreeCar/server/cmd/car/config"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	DB           *mongo.Collection
	ServerConfig config.ServerConfig
	NacosConfig  config.NacosConfig
)
