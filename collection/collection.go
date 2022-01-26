package collection

import (
	"encoding/json"
	"fmt"
)

type coll map[string]interface{}

type CollSets []coll

/* ================================== */
type collSetsCollection interface {
	WhereStr(string, string) collSetsCollection
	Groupby(string) collSetsCollection
	OrderBy() collSetsCollection
	Get() collSetsCollection
}

// WhereStr 相等的值回傳
func (c CollSets) WhereStr(compareKey, mapValue string) collSetsCollection {
	var o CollSets

	// 這層是coll
	for _, subColl := range c {
		if subColl[compareKey].(string) == mapValue {
			o = append(o, subColl)
		}
	}

	return o
}

// Groupby 指定的字串做聚集
func (c CollSets) Groupby(groupKey string) collSetsCollection {

	// 紀錄有曾出現的groupby key
	var groupDataList []string
	var o CollSets

	// 這層是coll
	for _, subColl := range c {

		// 不在groupby key的紀錄
		if StringInSlice(subColl[groupKey].(string), groupDataList) == false {
			// 先記錄在一個slice
			groupDataList = append(groupDataList, subColl[groupKey].(string))

			// 建立一個新的sub object (coll) 取代 map[string]interface{} 的 interface{}
			newColl := make(coll, 1)
			newColl["groupKey"] = subColl[groupKey].(string)

			// interface{} 被換成 []coll{}
			arr := []coll{}
			arr = append(arr, subColl)

			newColl["object"] = arr
			o = append(o, newColl)

		} else {
			// 在groupby key的紀錄
			for oKey, oValue := range o {

				if oValue["groupKey"] == subColl[groupKey].(string) {

					arr := []coll{}
					arr = append(arr, subColl)

					//I don't know that interface{} is replaced by []coll{}, and it is still considered to be interface{}, so it is asserted to be []coll{}
					o[oKey]["object"] = append(o[oKey]["object"].([]coll), subColl)
				}
			}
		}
	}

	return o
}

// OrderBy TODO
func (c CollSets) OrderBy() collSetsCollection {
	fmt.Println("The orderby condition is processed here(TODO)")
	return c
}

// Get return self
func (c CollSets) Get() collSetsCollection {
	a, _ := json.Marshal(&c)
	fmt.Println(string(a))
	return c
}
