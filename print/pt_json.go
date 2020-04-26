package print

import (
	"encoding/json"
	"fmt"
)

func PtJson(obj interface{}) {
	bytes, _ := json.MarshalIndent(obj, "", "    ")
	fmt.Println(string(bytes))
}
