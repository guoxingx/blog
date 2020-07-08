package webserver

import (
	"database/sql"
	"net/http"

	"backend/httpwrap"
	"backend/logs"
	"backend/utils/mysql"
	"go.uber.org/zap"
)

var (
	logger  *zap.SugaredLogger
	auth    *Authorizer
	db      *sql.DB
	statics string
)

// Run webserver
func Run(cfg Config) error {
	var err error

	auth, _ = NewAuthorizer("whenwearebothcats", cfg.Users)
	statics = cfg.Statics

	if cfg.CORS != "" {
		httpwrap.SetCors(cfg.CORS)
	}

	db, err = mysql.New(cfg.Mysql)
	if err != nil {
		logger.Fatalf("failed to connect mysql: %s", err)
	}

	// api handler
	http.HandleFunc("/api/blogs", httpwrap.Handler(handleBlogs).Base)
	http.HandleFunc("/api/blogs/", httpwrap.Handler(handleBlog).Base)

	logger.Infof("webserver listening, host: %s", cfg.Host)
	return http.ListenAndServe(cfg.Host, nil)
}

// Config for webserver
type Config struct {
	Host    string            `json:"host"`
	CORS    string            `json:"cors"`
	Users   map[string]string `json:"users"`
	Mysql   mysql.Config      `json:"mysql"`
	Statics string            `json:"statics"`
}

func init() {
	logger = logs.GetLogger()
}
