package controllers

import (
	"github.com/newrelic/go-agent"
	"github.com/revel/revel"
	"github.com/tamurayoshiya/modules/server-engine/newrelic"
)

type RelicController struct {
	*revel.Controller
}

func (r *RelicController) GetRelicApplication() newrelic.Application {
	if app, ok := revel.CurrentEngine.(*revelnewrelic.ServerNewRelic); ok {
		return app.NewRelicApp
	}
	return nil
}
