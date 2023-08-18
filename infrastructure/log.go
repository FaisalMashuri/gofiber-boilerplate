package infrastructure

import (
	"fmt"
	"github.com/olivere/elastic/v7"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gofiber-boilerplate/config"
	"gopkg.in/sohlich/elogrus.v7"
	"sync"
)

type LogCustom struct {
	Logrus *logrus.Logger
	WhoAmI iAm
}
type iAm struct {
	Name string
	Host string
	Port string
}

var instance *LogCustom
var once sync.Once

type stackTracer interface {
	StackTrace() errors.StackTrace
}

func NewLogCustom() *LogCustom {
	var log *logrus.Logger
	log = logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{PrettyPrint: true})

	client, err := elastic.NewClient(elastic.SetURL(
		fmt.Sprintf("http://%v:%v", config.AppConfig.Log.Host, config.AppConfig.Log.Port)),
		elastic.SetSniff(false),
		elastic.SetBasicAuth(config.AppConfig.Log.Username, config.AppConfig.Log.Password))
	if err != nil {
		selfLogError(err, "infrastructure/log: elastic client", log)
		return instance
	} else {
		hook, err := elogrus.NewAsyncElasticHook(
			client, config.AppConfig.Log.Host, logrus.DebugLevel, config.AppConfig.Log.Index)
		if err != nil {
			selfLogError(err, "infrastructure/log: elastic client", log)
		}
		log.Hooks.Add(hook)
	}

	once.Do(func() {
		instance = &LogCustom{
			Logrus: log,
			WhoAmI: iAm{
				Name: config.AppConfig.AppConfig.Name,
				Host: config.AppConfig.AppConfig.Host,
				Port: config.AppConfig.AppConfig.Port,
			},
		}
	})
	return instance
}

func selfLogError(err error, description string, log *logrus.Logger) {
	err = errors.WithStack(err)
	st := err.(stackTracer).StackTrace()
	stFormat := fmt.Sprintf("%+v", st[1:2])

	log.WithFields(logrus.Fields{
		"error_cause":   stFormat,
		"error_message": err.Error(),
	}).Error(description)
}

func (l *LogCustom) Fatal(err error, description string, traceHeader map[string]string) {
	err = errors.WithStack(err)
	st := err.(stackTracer).StackTrace()
	stFormat := fmt.Sprintf("%+v", st[1:2])

	l.Logrus.WithFields(logrus.Fields{
		"whoami":        l.WhoAmI,
		"trace_header":  traceHeader,
		"error_cause":   stFormat,
		"error_message": err.Error(),
	}).Fatal(description)
}
