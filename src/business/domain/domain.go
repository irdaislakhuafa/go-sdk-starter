package domain

import (
	"database/sql"

	"github.com/irdaislakhuafa/go-sdk-starter/src/business/domain/todo"
	entitygen "github.com/irdaislakhuafa/go-sdk-starter/src/entity/gen"
	"github.com/irdaislakhuafa/go-sdk/log"
)

type (
	Domain struct {
		Todo todo.Interface
	}
)

func Init(log log.Interface, queries *entitygen.Queries, db *sql.DB) *Domain {
	return &Domain{
		Todo: todo.Init(log, queries, db),
	}
}
