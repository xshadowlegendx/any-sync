package commonspace

import (
	"context"
	"github.com/anytypeio/go-anytype-infrastructure-experiments/common/account"
	"github.com/anytypeio/go-anytype-infrastructure-experiments/common/app"
	"github.com/anytypeio/go-anytype-infrastructure-experiments/common/app/logger"
	"github.com/anytypeio/go-anytype-infrastructure-experiments/common/commonspace/diffservice"
	"github.com/anytypeio/go-anytype-infrastructure-experiments/common/commonspace/storage"
	"github.com/anytypeio/go-anytype-infrastructure-experiments/common/commonspace/syncservice"
	"github.com/anytypeio/go-anytype-infrastructure-experiments/common/commonspace/treegetter"
	"github.com/anytypeio/go-anytype-infrastructure-experiments/common/config"
	"github.com/anytypeio/go-anytype-infrastructure-experiments/common/net/pool"
	"github.com/anytypeio/go-anytype-infrastructure-experiments/common/nodeconf"
)

const CName = "common.commonspace"

var log = logger.NewNamed(CName)

func New() Service {
	return &service{}
}

type Service interface {
	DeriveSpace(ctx context.Context, payload SpaceDerivePayload) (string, error)
	CreateSpace(ctx context.Context, payload SpaceCreatePayload) (string, error)
	NewSpace(ctx context.Context, id string) (sp Space, err error)
	app.Component
}

type service struct {
	config               config.Space
	account              account.Service
	configurationService nodeconf.Service
	storageProvider      storage.SpaceStorageProvider
	treeGetter           treegetter.TreeGetter
	pool                 pool.Pool
}

func (s *service) Init(a *app.App) (err error) {
	s.config = a.MustComponent(config.CName).(*config.Config).Space
	s.account = a.MustComponent(account.CName).(account.Service)
	s.storageProvider = a.MustComponent(storage.CName).(storage.SpaceStorageProvider)
	s.configurationService = a.MustComponent(nodeconf.CName).(nodeconf.Service)
	s.treeGetter = a.MustComponent(treegetter.CName).(treegetter.TreeGetter)
	s.pool = a.MustComponent(pool.CName).(pool.Pool)
	return nil
}

func (s *service) Name() (name string) {
	return CName
}

func (s *service) CreateSpace(ctx context.Context, payload SpaceCreatePayload) (id string, err error) {
	storageCreate, err := storagePayloadForSpaceCreate(payload)
	if err != nil {
		return
	}
	store, err := s.storageProvider.CreateSpaceStorage(storageCreate)
	if err != nil {
		return
	}

	return store.Id(), nil
}

func (s *service) DeriveSpace(ctx context.Context, payload SpaceDerivePayload) (id string, err error) {
	storageCreate, err := storagePayloadForSpaceDerive(payload)
	if err != nil {
		return
	}
	store, err := s.storageProvider.CreateSpaceStorage(storageCreate)
	if err != nil {
		return
	}

	return store.Id(), nil
}

func (s *service) NewSpace(ctx context.Context, id string) (Space, error) {
	st, err := s.storageProvider.SpaceStorage(id)
	if err != nil {
		return nil, err
	}

	lastConfiguration := s.configurationService.GetLast()
	confConnector := nodeconf.NewConfConnector(lastConfiguration, s.pool)
	diffService := diffservice.NewDiffService(id, s.config.SyncPeriod, st, confConnector, s.treeGetter, log)
	syncService := syncservice.NewSyncService(id, confConnector)
	sp := &space{
		id:            id,
		syncService:   syncService,
		diffService:   diffService,
		cache:         s.treeGetter,
		account:       s.account,
		configuration: lastConfiguration,
		storage:       st,
	}
	return sp, nil
}
