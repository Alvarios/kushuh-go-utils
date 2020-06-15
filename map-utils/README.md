# Kushuh go map utils

## Merge

Merge multiple `map[string]interface{}` objects together.

*Variables*
```go
original := map[string]interface{}{
    "key1": "value1",
    "key2": map[string]interface{}{
        "key21": 10,
        "key22": true,
    },
    "key3": map[string]interface{}{
        "key3-1": map[string]interface{}{
            "ultra-nested": []string{},
        },
        "key32": false,
    },
}

replacement1 := map[string]interface{}{
    "key2": "value2",
    "key3": map[string]interface{}{
        "key32": 11,
    },
}

replacement2 := map[string]interface{}{
    "key4": "value4",
}
```

*Function call*
```go
import "github.com/Alvarios/kushuh-go-utils/map-utils"

// Add as many replacement objects as you want.
newMap, err := mapUtils.Merge(original, replacement1, replacement2)
```

Will result in:

```go
map[string]interface{}{
    "key1": "value1",
    "key2": "value2",
    "key3": map[string]interface{}{
        "key31": map[string]interface{}{
            "ultra-nested": []string{},
        }
        "key32": 11,
    },
    "key4": "value4",
}
```

### Flatten

Extract each nested key of a map at top level.

```go
value := map[string]interface{}{
    "key1": "value1",
    "key2": "value2",
    "key3": map[string]interface{}{
        "key31": map[string]interface{}{
            "ultra-nested": []string{},
        }
        "key32": 11,
    },
    "key4": "value4",
}

flattened := mapUtils.Flatten(value)
```

Will result in:

```go
map[string]interface{}{
    "key1": "value1",
    "key2": "value2",
    "key3.key31.ultra-nested": []string{},
    "key3.key32": 11,
    "key4": "value4",
}
```

## Copyright
2020 Kushuh - [MIT license](https://github.com/Alvarios/kushuh-go-utils/blob/master/LICENSE)
