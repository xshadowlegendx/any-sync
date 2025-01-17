package nameserviceclient

import (
	"context"

	"github.com/anyproto/any-sync/app"
	"github.com/anyproto/any-sync/app/logger"
	"github.com/anyproto/any-sync/net/pool"
	"github.com/anyproto/any-sync/net/rpc/rpcerr"
	"github.com/anyproto/any-sync/nodeconf"
	"go.uber.org/zap"

	nsp "github.com/anyproto/any-sync/nameservice/nameserviceproto"
)

const CName = "nameservice.nameserviceclient"

var log = logger.NewNamed(CName)

/*
 * This client component can be used to access the Any Naming System (any-ns)
 * from other components.
 */
type AnyNsClientServiceBase interface {
	IsNameAvailable(ctx context.Context, in *nsp.NameAvailableRequest) (out *nsp.NameAvailableResponse, err error)
	// reverse resolve
	GetNameByAddress(ctx context.Context, in *nsp.NameByAddressRequest) (out *nsp.NameByAddressResponse, err error)

	app.Component
}

type AnyNsClientService interface {
	// AccountAbstractions methods:
	GetUserAccount(ctx context.Context, in *nsp.GetUserAccountRequest) (out *nsp.UserAccount, err error)
	AdminFundUserAccount(ctx context.Context, in *nsp.AdminFundUserAccountRequestSigned) (out *nsp.OperationResponse, err error)

	GetOperation(ctx context.Context, in *nsp.GetOperationStatusRequest) (out *nsp.OperationResponse, err error)
	CreateOperation(ctx context.Context, in *nsp.CreateUserOperationRequestSigned) (out *nsp.OperationResponse, err error)

	AnyNsClientServiceBase
}

type service struct {
	pool     pool.Pool
	nodeconf nodeconf.Service
}

func (s *service) Init(a *app.App) (err error) {
	s.pool = a.MustComponent(pool.CName).(pool.Pool)
	s.nodeconf = a.MustComponent(nodeconf.CName).(nodeconf.Service)
	return nil
}

func (s *service) Name() (name string) {
	return CName
}

func New() AnyNsClientService {
	return new(service)
}

func (s *service) doClient(ctx context.Context, fn func(cl nsp.DRPCAnynsClient) error) error {
	// it will try to connect to the Naming Node
	// please enable "namingNode" type of node in the config (in the network.nodes array)
	peer, err := s.pool.Get(ctx, s.nodeconf.NamingNodePeers()[0])
	log.Info("trying to connect to namingNode peer: ", zap.Any("peer", peer))

	if err != nil {
		return err
	}

	dc, err := peer.AcquireDrpcConn(ctx)
	if err != nil {
		return err
	}
	defer peer.ReleaseDrpcConn(dc)

	return fn(nsp.NewDRPCAnynsClient(dc))
}

func (s *service) doClientAA(ctx context.Context, fn func(cl nsp.DRPCAnynsAccountAbstractionClient) error) error {
	// it will try to connect to the Naming Node
	// please enable "namingNode" type of node in the config (in the network.nodes array)
	peer, err := s.pool.Get(ctx, s.nodeconf.NamingNodePeers()[0])
	log.Info("trying to connect to namingNode peer: ", zap.Any("peer", peer))

	if err != nil {
		return err
	}

	dc, err := peer.AcquireDrpcConn(ctx)
	if err != nil {
		return err
	}
	defer peer.ReleaseDrpcConn(dc)

	return fn(nsp.NewDRPCAnynsAccountAbstractionClient(dc))
}

func (s *service) IsNameAvailable(ctx context.Context, in *nsp.NameAvailableRequest) (out *nsp.NameAvailableResponse, err error) {
	err = s.doClient(ctx, func(cl nsp.DRPCAnynsClient) error {
		if out, err = cl.IsNameAvailable(ctx, in); err != nil {
			return rpcerr.Unwrap(err)
		}
		return nil
	})
	return
}

func (s *service) GetNameByAddress(ctx context.Context, in *nsp.NameByAddressRequest) (out *nsp.NameByAddressResponse, err error) {
	err = s.doClient(ctx, func(cl nsp.DRPCAnynsClient) error {
		if out, err = cl.GetNameByAddress(ctx, in); err != nil {
			return rpcerr.Unwrap(err)
		}
		return nil
	})
	return
}

// AA
func (s *service) GetUserAccount(ctx context.Context, in *nsp.GetUserAccountRequest) (out *nsp.UserAccount, err error) {
	err = s.doClientAA(ctx, func(cl nsp.DRPCAnynsAccountAbstractionClient) error {
		if out, err = cl.GetUserAccount(ctx, in); err != nil {
			return rpcerr.Unwrap(err)
		}
		return nil
	})
	return
}

func (s *service) AdminFundUserAccount(ctx context.Context, in *nsp.AdminFundUserAccountRequestSigned) (out *nsp.OperationResponse, err error) {
	err = s.doClientAA(ctx, func(cl nsp.DRPCAnynsAccountAbstractionClient) error {
		if out, err = cl.AdminFundUserAccount(ctx, in); err != nil {
			return rpcerr.Unwrap(err)
		}
		return nil
	})
	return
}

func (s *service) GetOperation(ctx context.Context, in *nsp.GetOperationStatusRequest) (out *nsp.OperationResponse, err error) {
	err = s.doClientAA(ctx, func(cl nsp.DRPCAnynsAccountAbstractionClient) error {
		if out, err = cl.GetOperation(ctx, in); err != nil {
			return rpcerr.Unwrap(err)
		}
		return nil
	})
	return
}

func (s *service) CreateOperation(ctx context.Context, in *nsp.CreateUserOperationRequestSigned) (out *nsp.OperationResponse, err error) {
	err = s.doClientAA(ctx, func(cl nsp.DRPCAnynsAccountAbstractionClient) error {
		if out, err = cl.CreateUserOperation(ctx, in); err != nil {
			return rpcerr.Unwrap(err)
		}
		return nil
	})
	return
}
