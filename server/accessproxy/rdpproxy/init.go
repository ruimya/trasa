package rdpproxy

import (
	"github.com/seknox/trasa/server/api/logs"
	"github.com/seknox/trasa/server/consts"
	"github.com/seknox/trasa/server/global"
	"github.com/seknox/trasa/server/models"
)

//InitStore initialises package state
func InitStore(state *global.State, policyFunc models.CheckPolicyFunc) {
	Store = GWStore{
		proxy:           NewProxy(),
		State:           state,
		checkPolicyFunc: policyFunc,
	}
}

//Store is the package state variable which contains database connections
var Store GWAdapter

type GWStore struct {
	*global.State
	proxy           *Proxy
	checkPolicyFunc models.CheckPolicyFunc
}

type GWAdapter interface {
	CheckPolicy(params *models.ConnectionParams, policy *models.Policy, adhoc bool) (bool, consts.FailedReason)
	uploadSessionLog(authlog *logs.AuthLog) error
}
