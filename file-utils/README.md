# Kushuh go file utils

## Upsert

Open a file, and create it if it doesn't exist already. If path is incomplete
(intermediate folders are missing), they will also be created automatically.

```go
file, err := fileUtils.Upsert(pathToFile)
```

## Delete

Delete a file.

```go
err := fileUtils.Delete(pathToFile)
```

## OpenFromProjectRoot

Open a file relative to project root, when a command is ran from a sub-directory
(which happens with tests and some other commands).

It ensures a file opens from project root with go modules, without having to pass
it as an ENV variable (like GOPATH does in non module go packages). It assumes
the project folder has an unique name in the entire project tree.

```go
file, err := fileUtils.OpenFromProjectRoot(rootFolderName, relativePath)
```

## FindProjectRoot

Returns the absolute path for the root of a project. It assumes
the project folder has an unique name in the entire project tree.

```go
path, err := fileUtils.FindProjectRoot(rootFolderName)
```

## Copyright
2020 Kushuh - [MIT license](https://github.com/Alvarios/kushuh-go-utils/blob/master/LICENSE)
