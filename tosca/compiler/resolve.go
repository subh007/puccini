package compiler

import (
	"github.com/tliron/puccini/clout"
	"github.com/tliron/puccini/js"
	"github.com/tliron/puccini/tosca/problems"
)

func Resolve(clout_ *clout.Clout, problems_ *problems.Problems) *clout.Clout {
	context := js.NewContext("tosca.resolve", log, true, "yaml", "")
	if err := context.Exec(clout_, "tosca.resolve", map[string]interface{}{"problems": problems_}); err != nil {
		problems_.ReportError(err)
	}

	return clout_
}
