# Kushuh go slice utils

## Compare

Check if two slices are equal in Go. Both slice can be of any type, but
their elements should be sorted in the same order.

```go
slice1 := []int{2,1,3}
slice2 := []int{2,1,3}
slice1 := []int{3,2,1}

isEqual := sliceUtils.Compare(slice1, slice2) // return true

isEqual = sliceUtils.Compare(slice1, slice3) // return false
```

## Copyright
2020 Kushuh - [MIT license](https://github.com/Alvarios/kushuh-go-utils/blob/master/LICENSE)
