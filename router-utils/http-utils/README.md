# Kushuh go http utils

## Post

Send a post request from go, in `application/json` format.

```go
res, err := httpUtils.Post(url, payload)
```

## PostForm

Send a post request from go, with a form content.

```go
files := []httpUtils.File{
    {
        Path: '/path/to/my/file',
        Key: 'file'
    }
}

res, err := httpUtils.PostForm(url, data, files)
```

## Copyright
2020 Kushuh - [MIT license](https://github.com/Alvarios/kushuh-go-utils/blob/master/LICENSE)
