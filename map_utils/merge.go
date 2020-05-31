package map_utils

import "errors"

func Merge(origin interface{}, invaders ...interface{}) (interface{}, error) {
	mo, ok := origin.(map[string]interface{})

	if ok == false {
		return nil, errors.New("cannot convert first argument to map object")
	}

	for _, element := range invaders {
		me, ok := element.(map[string]interface{})
		if ok == false {
			return nil, errors.New("argument is not of type map[string]interface{}")
		}

		for key, value := range me {
			mappedValue, ok := value.(map[string]interface{})
			mappedOriginValue, originOk := mo[key].(map[string]interface{})

			if ok {
				if originOk == false {
					mappedOriginValue = map[string]interface{}{}
				}
				newValue, err := Merge(mappedOriginValue, mappedValue)
				if err != nil {
					return nil, err
				}
				mo[key] = newValue
			} else {
				mo[key] = value
			}
		}
	}

	return mo, nil
}