# errors

[![Go Reference](https://pkg.go.dev/badge/github.com/kubasobon/errors.svg)](https://pkg.go.dev/github.com/kubasobon/errors)

A small package to make working with errors in Golang nicer. For use in
personal projects.

```golang
err := errors.New("something went wrong: %v", value)
```

```golang
err := errors.NewOfKind(errors.ConfigError, "expected field was empty")
```

```golang
const invalidValueError errors.ErrorKind = "InvalidValueError"
err := errors.NewOfKind(invalidValueError, "expected X, got %#q", v)
```

```golang
err := errors.Mask(err)
```
