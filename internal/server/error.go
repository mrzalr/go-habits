package server

import "errors"

var ErrInvalidPort = errors.New("SERVER: invalid port, please check your config.yaml")
