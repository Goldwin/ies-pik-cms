package mongo

import (
	"context"

	"github.com/Goldwin/ies-pik-cms/pkg/auth/commands"
	"github.com/Goldwin/ies-pik-cms/pkg/common/worker"
	"go.mongodb.org/mongo-driver/mongo"
)

type authUnitOfWorkImpl struct {
	db             *mongo.Database
	useTransaction bool
}

// Execute implements worker.UnitOfWork.
func (a *authUnitOfWorkImpl) Execute(ctx context.Context, op worker.AtomicOperation[commands.CommandContext]) error {
	db := a.db
	session, err := db.Client().StartSession()
	if err != nil {
		return err
	}
	defer session.EndSession(ctx)

	//transaction can only be done if mongo use replica set. otherwise it will fail.
	if a.useTransaction {
		_, err = session.WithTransaction(ctx, func(sessionContext mongo.SessionContext) (interface{}, error) {
			c := NewCommandContext(sessionContext, db)
			return nil, op(c)
		})
		return err
	}

	c := NewCommandContext(ctx, db)
	return op(c)
}

func NewUnitOfWork(db *mongo.Database, useTransaction bool) worker.UnitOfWork[commands.CommandContext] {
	return &authUnitOfWorkImpl{
		db:             db,
		useTransaction: useTransaction,
	}
}
