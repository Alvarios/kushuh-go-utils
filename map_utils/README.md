# Kushuh go map utils

## Merge

Merge multiple `map[string]interface{}` objects together.

*Variables*
```go
original := map[string]interface{}{
    "key1": "original value",
    "key2": map[string]interface{}{
        "key2-1": 10,
        "key2-2": true
    },
    "key3": map[string]interface{}{
        "key3-1": map[string]interface{}{
            "ultra-nested": []string{}
        }
        "key3-2": false,
    },
}

replacement1 := map[string]interface{}{
    "key2": "Maliwan",
    "key3": map[string]interface{}{
        "key3-2": 11,
    },
}

replacement2 := map[string]interface{}{
    "key4": "a new key",
}
```

*Function call*
```go
import "github.com/Alvarios/kushuh-go-utils/map_utils"

// Add as many replacement objects as you want.
newMap := map_utils.merge(original, replacement1, replacement2)
```

Will result in:

```go
map[string]interface{}{
    "key1": "original value",
    "key2": "Maliwan",
    "key3": map[string]interface{}{
        "key3-1": map[string]interface{}{
            "ultra-nested": []string{}
        }
        "key3-2": 11,
    },
    "key4": "a new key",
}
```

## Copyright
2020 Kushuh - [MIT license](https://github.com/Alvarios/kushuh-go-utils/blob/master/LICENSE)
