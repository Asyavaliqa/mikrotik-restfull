package routeros

import (
	"fmt"
	"reflect"
)

func (util *RouterOS) Add(data interface{}) *RouterOS {
	//set action last command
	util.Query[len(util.Query)-1] += "/add"

	var stmt reflect.Value = reflect.ValueOf(data)
	if stmt.IsNil() && stmt.CanAddr() {
		stmt.Set(reflect.New(stmt.Type()))
	}

	typeOfS := stmt.Elem().Type()
	for i := 0; i < stmt.Elem().NumField(); i++ {
		if stmt.Elem().Field(i).CanInterface() {
			kName := typeOfS.Field(i).Tag.Get("json")
			kValue := stmt.Elem().Field(i)
			util.Query = append(util.Query, fmt.Sprintf("=%s=%s", kName, kValue)) //query[kName] = kValue.String()
		}
	}

	fmt.Println(util.Query)
	util.Run(util.Query)

	return util
}
