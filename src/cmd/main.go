package main

import (
	"context"
	"flag"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/irdaislakhuafa/go-sdk-starter/src/business/domain"
	"github.com/irdaislakhuafa/go-sdk-starter/src/business/usecase"
	entitygen "github.com/irdaislakhuafa/go-sdk-starter/src/entity/gen"
	"github.com/irdaislakhuafa/go-sdk-starter/src/handler/rest"
	"github.com/irdaislakhuafa/go-sdk-starter/src/utils/config"
	"github.com/irdaislakhuafa/go-sdk-starter/src/utils/connection"
	"github.com/irdaislakhuafa/go-sdk/log"
	"github.com/irdaislakhuafa/go-sdk/smtp"
	"github.com/irdaislakhuafa/go-sdk/storage"
)

const (
	configFileJSON = "etc/cfg/conf.json"
)

func main() {
	// read config file
	cfg, err := config.ReadFileJSON(configFileJSON)
	if err != nil {
		panic(err)
	}

	// initialize db
	db, err := connection.InitMySQL(cfg)
	if err != nil {
		panic(err)
	}

	// initialize log
	l := log.Init(cfg.Log)

	// initialize validator
	v := validator.New(
		validator.WithRequiredStructEnabled(),
	)

	// initialize storage client
	s, err := storage.InitMinio(cfg.Storage)
	if err != nil {
		panic(err)
	}

	// initialize queries
	q := entitygen.New(db)

	// initialize smtp
	smtpGoMail := smtp.InitGoMail(cfg.SMTP)

	// initialize domain
	dom := domain.Init(l, q, db, s)

	// initialize usecase
	uc := usecase.Init(l, cfg, v, dom, smtpGoMail, s)

	// choose running mode
	mode := flag.String("mode", "rest", "Please select running mode: [rest, gql, grpc]")
	flag.Parse()
	*mode = strings.ToLower(*mode)

	switch *mode {
	case "rest":
		r := rest.Init(cfg, l, uc)
		r.Run()
		// TODO: write for grpc and gql
	default:
		l.Fatal(context.Background(), fmt.Sprintf("Running mode '%s' not supported!", *mode))
	}
}
