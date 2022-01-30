package routeros

import (
	"encoding/json"
	"fmt"
	"strings"
)

func (util *RouterOS) Print(bind interface{}) *RouterOS {
	//set action last command
	util.Query[len(util.Query)-1] += "/print"

	//cek filter
	util.Query = append(util.Query, util.Filter...)

	// Run Query
	util.Run(util.Query)

	// Deteksi Error
	if util.DetectError() {
		return util
	}

	// Parsing Data Mikrotik
	jsonbody, err := json.Marshal(util.Reply.Re[0].Map)
	if err != nil {
		// do error check
		util.Error = err
		return util
	}

	if err := json.Unmarshal(jsonbody, bind); err != nil {
		// do error check
		util.Error = err
		return util
	}
	util.Debug().Msg(fmt.Sprintf("| DEBUG | [%d Fields] %s", len(util.Re[0].Map), strings.Join(util.Query, " ")))

	return util
}
