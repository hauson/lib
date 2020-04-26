package exception

import (
	"reflect"
	"fmt"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func Assert(got, expect interface{}) {
	if !reflect.DeepEqual(got, expect) {
		panic(fmt.Sprintf("got:%v not equal expect:%v", got, expect))
	}
}
