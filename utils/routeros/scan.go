package routeros

import "encoding/json"

func (util *RouterOS) Scan(bind interface{}) *RouterOS {
	//set action last command
	util.Query[len(util.Query)-1] += "/print"

	// Run Query
	util.Run(util.Query)

	// Deteksi Error
	if util.DetectError() {
		return util
	}

	var dataArr []interface{}
	for i := 0; i < len(util.Reply.Re); i++ {
		dataArr = append(dataArr, util.Reply.Re[i].Map)
	}

	jsonbody, err := json.Marshal(dataArr)
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
	return util
}
