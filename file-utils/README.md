# Kushuh go file utils

## OpenFromProjectRoot

Open a file relative to project root. Especially useful when working with
complex project trees, where commands such as tests can run from a sub-directory.

It ensures a file opens from project root with go modules, without having to pass
it as an ENV variable (like GOPATH does in non module go packages).

```go
file, err := fileUtils.OpenFromProjectRoot(rootFolderName, relativePath)
```

## FindProjectRoot

Returns the absolute path for the root of a project. It assumes naming
conventions, such as only the root folder is named after the project name.

```go
path, err := file_utils.FindProjectRoot(rootFolderName)
```

## Copyright
2020 Kushuh - [MIT license](https://github.com/Alvarios/kushuh-go-utils/blob/master/LICENSE)
