package connection

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/irdaislakhuafa/go-sdk-starter/src/utils/config"
	"github.com/irdaislakhuafa/go-sdk/codes"
	"github.com/irdaislakhuafa/go-sdk/errors"
	"github.com/irdaislakhuafa/go-sdk/strformat"
)

func InitMySQL(cfg config.Config) (*sql.DB, error) {
	ds := "{{ .Username }}:{{ .Password }}@tcp({{ .Host }}:{{ .Port }})/{{ .Name }}?parseTime=true"
	ds = strformat.TWE(ds, cfg.DB.Master)
	db, err := sql.Open(DriverNameMySQL, ds)
	if err != nil {
		return nil, errors.NewWithCode(codes.CodeSQLInit, "cannot connect to db: %s", err.Error())
	}

	return db, nil
}
