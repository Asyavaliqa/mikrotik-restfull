package routeros

import (
	"fmt"
	"github.com/didintri196/mikrotik-restfull/domain/models"
	"github.com/rs/zerolog"
	"gopkg.in/routeros.v2"
	"os"
	"strings"
)

type RouterOS struct {
	models.Auth
	*routeros.Reply
	Error error
	zerolog.Logger
	Query  []string
	Filter []string
}

type Configs struct {
	Auth  models.Auth
	Debug bool
}

func New(config Configs) RouterOS {
	zerolog.SetGlobalLevel(zerolog.NoLevel)
	if config.Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	return RouterOS{config.Auth, &routeros.Reply{}, nil, RegisterLog(), []string{}, []string{}}
}

func (util *RouterOS) Run(query []string) *RouterOS {
	host := fmt.Sprintf("%s:%s", util.Ip, util.Port)
	c, err := routeros.Dial(host, util.Username, util.Password)
	if err != nil {
		util.Debug().Msg(fmt.Sprintf("| ERROR | %s", err.Error()))
		util.Error = err
		return util
	}
	defer c.Close()

	re, err := c.RunArgs(query)
	if err != nil {
		util.Error = err
		return util
	}
	util.Reply = re
	return util
}

func (util *RouterOS) Command(query string) *RouterOS {
	util.Query = []string{query}
	return util
}

func (util *RouterOS) DetectError() bool {
	if util.Error != nil {
		return true
	}
	return false
}

func RegisterLog() zerolog.Logger {
	output := zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: "2006-01-02 15:04:05"}
	output.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf(""))
	}
	output.FormatMessage = func(i interface{}) string {
		return fmt.Sprintf("%s", i)
	}
	output.FormatFieldName = func(i interface{}) string {
		return fmt.Sprintf("%s:", i)
	}
	output.FormatFieldValue = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("%s", i))
	}

	return zerolog.New(output).With().Timestamp().Logger()
}
