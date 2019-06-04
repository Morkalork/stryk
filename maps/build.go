package maps

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
	"stryk/core"
)

func save(levelMap core.LevelMap, levelNumber int) {
	jsonMap, err := json.Marshal(levelMap)

	if err != nil {
		println(fmt.Sprintf("Error occured while marshalling the levelMap: %s", err))
		return
	}

	fileName := fmt.Sprintf("assets/map-%d.json", levelNumber)
	err = ioutil.WriteFile(fileName, jsonMap, 0777)

	if err != nil {
		println(fmt.Sprintf("Error when writing file to disk: %s (%s)", err, fileName))
		return

	}
}

// A function that runs through all available maps and builds them.
// If you're adding a new map, don't forget to add it into the 'mapFuncs'
// array!
func Build() {
	mapFuncs := map[string]interface{}{
		"map1": Map1,
		"map2": Map2,
	}

	for i := 1; i <= len(mapFuncs); i++ {
		interfaceName := fmt.Sprintf("map%d", i)
		fmt.Println("Building map", interfaceName)
		mapFunc := reflect.ValueOf(mapFuncs[interfaceName])
		in := make([]reflect.Value, 0)
		levelMap := mapFunc.Call(in)[0].Interface().(core.LevelMap)
		save(levelMap, i)
	}
}
