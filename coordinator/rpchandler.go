package coordinator

import (
	"context"
	"github.com/anytypeio/any-sync/coordinator/coordinatorproto"
)

type rpcHandler struct {
	c *coordinator
}

func (r *rpcHandler) SpaceStatusCheck(ctx context.Context, request *coordinatorproto.SpaceStatusCheckRequest) (*coordinatorproto.SpaceStatusCheckResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r *rpcHandler) SpaceStatusChange(ctx context.Context, request *coordinatorproto.SpaceStatusChangeRequest) (*coordinatorproto.SpaceStatusChangeResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r *rpcHandler) SpaceSign(ctx context.Context, req *coordinatorproto.SpaceSignRequest) (*coordinatorproto.SpaceSignResponse, error) {
	receipt, err := r.c.SpaceSign(ctx, req.SpaceId, req.Header)
	if err != nil {
		return nil, err
	}
	return &coordinatorproto.SpaceSignResponse{
		Receipt: receipt,
	}, nil
}

func (r *rpcHandler) FileLimitCheck(ctx context.Context, req *coordinatorproto.FileLimitCheckRequest) (*coordinatorproto.FileLimitCheckResponse, error) {
	limit, err := r.c.FileLimitCheck(ctx, req.AccountIdentity, req.SpaceId)
	if err != nil {
		return nil, err
	}
	return &coordinatorproto.FileLimitCheckResponse{
		Limit: limit,
	}, nil
}
