package map_utils

func Merge(origin map[string]interface{}, invaders ...map[string]interface{}) map[string]interface{} {
	for _, element := range invaders {
		for key, value := range element {
			mappedValue, ok := value.(map[string]interface{})
			mappedOriginValue, originOk := origin[key].(map[string]interface{})

			if ok {
				if originOk == false {
					mappedOriginValue = map[string]interface{}{}
				}
				origin[key] = Merge(mappedOriginValue, mappedValue)
			} else {
				origin[key] = value
			}
		}
	}

	return origin
}