package warehouseserver

import (
	"context"
	"log"
	"math/rand"
	"time"

	pb "github.com/seblw/twirp_gowroc/rpc/warehouse"
	"github.com/twitchtv/twirp"
)

type Server struct{}

func (s *Server) GetOrders(ctx context.Context, req *pb.GetOrdersReq) (resp *pb.GetOrdersResp, err error) {
	if req.GetStatus() == "" {
		return nil, twirp.InvalidArgumentError("status", "cannot be empty!")
	}

	t, ok := getReqRecv(ctx)
	if ok {
		log.Printf("Took %v since received request", time.Since(t))
	}

	return &pb.GetOrdersResp{
		Orders: []*pb.Order{
			{
				Id:     rand.Int31n(10),
				Type:   []string{"clothes", "food", "books", "toys", "electronics"}[rand.Intn(4)],
				Status: req.GetStatus(),
			},
			{
				Id:     rand.Int31n(10) + 10,
				Type:   []string{"clothes", "food", "books", "toys", "electronics"}[rand.Intn(4)],
				Status: req.GetStatus(),
			},
		},
	}, nil
}

func getReqRecv(ctx context.Context) (time.Time, bool) {
	t, ok := ctx.Value("req_recv").(time.Time)
	return t, ok
}
