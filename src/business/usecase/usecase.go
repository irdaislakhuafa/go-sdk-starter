package usecase

import (
	"github.com/go-playground/validator/v10"
	"github.com/irdaislakhuafa/go-sdk-starter/src/business/domain"
	"github.com/irdaislakhuafa/go-sdk-starter/src/business/usecase/todo"
	"github.com/irdaislakhuafa/go-sdk-starter/src/entity"
	"github.com/irdaislakhuafa/go-sdk-starter/src/utils/config"
	"github.com/irdaislakhuafa/go-sdk/caches"
	"github.com/irdaislakhuafa/go-sdk/log"
	"github.com/irdaislakhuafa/go-sdk/smtp"
	"github.com/irdaislakhuafa/go-sdk/storage"
)

type (
	Usecase struct {
		Todo todo.Interface
	}
)

func Init(
	log log.Interface,
	cfg config.Config,
	val *validator.Validate,
	dom *domain.Domain,
	smtp smtp.Interface,
	storage storage.Interface,
	cache caches.Interface[entity.Cache],
) *Usecase {
	return &Usecase{
		Todo: todo.Init(log, cfg, val, dom),
	}
}
