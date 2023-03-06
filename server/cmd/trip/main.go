package main

import (
	"context"
	"net"
	"strconv"

	"github.com/CyanAsterisk/FreeCar/server/cmd/trip/global"
	"github.com/CyanAsterisk/FreeCar/server/cmd/trip/initialize"
	"github.com/CyanAsterisk/FreeCar/server/cmd/trip/tool/car"
	"github.com/CyanAsterisk/FreeCar/server/cmd/trip/tool/poi"
	"github.com/CyanAsterisk/FreeCar/server/cmd/trip/tool/profile"
	"github.com/CyanAsterisk/FreeCar/server/shared/consts"
	"github.com/CyanAsterisk/FreeCar/server/shared/kitex_gen/trip/tripservice"
	middleware2 "github.com/CyanAsterisk/FreeCar/server/shared/middleware"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/utils"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
)

func main() {
	// initialization
	initialize.InitLogger()
	IP, Port := initialize.InitFlag()
	r, info := initialize.InitNacos(Port)
	initialize.InitDB()
	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(global.ServerConfig.Name),
		provider.WithExportEndpoint(global.ServerConfig.OtelInfo.EndPoint),
		provider.WithInsecure(),
	)
	defer p.Shutdown(context.Background())
	initialize.InitCar()
	initialize.InitProfile()

	impl := new(TripServiceImpl)
	impl.CarManager = &car.Manager{
		CarService: global.CarClient,
	}
	impl.ProfileManager = &profile.Manager{
		ProfileService: global.ProfileClient,
	}
	impl.POIManager = &poi.Manager{}
	// Create new server.
	srv := tripservice.NewServer(impl,
		server.WithServiceAddr(utils.NewNetAddr(consts.TCP, net.JoinHostPort(IP, strconv.Itoa(Port)))),
		server.WithRegistry(r),
		server.WithRegistryInfo(info),
		server.WithLimit(&limit.Option{MaxConnections: 2000, MaxQPS: 500}),
		server.WithMiddleware(middleware2.CommonMiddleware),
		server.WithMiddleware(middleware2.ServerMiddleware),
		server.WithSuite(tracing.NewServerSuite()),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: global.ServerConfig.Name}),
	)

	err := srv.Run()
	if err != nil {
		klog.Fatal(err)
	}
}
