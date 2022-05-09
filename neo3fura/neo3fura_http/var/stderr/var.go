package stderr

import (
	"neo3fura_http/lib/jsonrpc2"
)

// ErrUnknown ...
var ErrUnknown = jsonrpc2.NewError(-32001, "unknown error")

// ErrInvalidArgs ...
var ErrInvalidArgs = jsonrpc2.NewError(-32001, "invalid args")

// ErrUnsupportedMethod ...
var ErrUnsupportedMethod = jsonrpc2.NewError(-32001, "unsupported method")

// ErrNotFound ...
var ErrNotFound = jsonrpc2.NewError(-32001, "not found")

// ErrZero
var ErrZero = jsonrpc2.NewError(-32001, "txid cannot be zero")

// FindDocumentErr
var ErrFind = jsonrpc2.NewError(-32001, "find document(s) error")

// InsertJobErr
var ErrInsert = jsonrpc2.NewError(-32001, "insert job error")

// InsertJobErr
var ErrArgsInner = jsonrpc2.NewError(-32001, "start must not bigger end")

var ErrAMarketConfig = jsonrpc2.NewError(-32001, "market config error")

var ErrPrice = jsonrpc2.NewError(-32001, "asset convent price error")

var ErrData = jsonrpc2.NewError(-32001, "data conversion error")

var ErrWhiteList = jsonrpc2.NewError(-32001, "whiteList is null")

var ErrInsertDocument = jsonrpc2.NewError(-32001, "insert document error")

var ErrFind1 = jsonrpc2.NewError(-32001, "find document(s) error1")

var ErrFind2 = jsonrpc2.NewError(-32001, "find document(s) error2")

var ErrFind3 = jsonrpc2.NewError(-32001, "find document(s) error3")

var ErrFind4 = jsonrpc2.NewError(-32001, "find document(s) error4")

var ErrFind5 = jsonrpc2.NewError(-32001, "find document(s) error5")

var ErrFind6 = jsonrpc2.NewError(-32001, "find document(s) error6")

var ErrFind7 = jsonrpc2.NewError(-32001, "find document(s) error7")

var ErrFind8 = jsonrpc2.NewError(-32001, "find document(s) error8")

var ErrFind9 = jsonrpc2.NewError(-32001, "find document(s) error9")

var ErrFind10 = jsonrpc2.NewError(-32001, "find document(s) error10")

var ErrFind11 = jsonrpc2.NewError(-32001, "find document(s) error11")

var ErrFind12 = jsonrpc2.NewError(-32001, "find document(s) error12")

var ErrFind13 = jsonrpc2.NewError(-32001, "find document(s) error13")

var ErrFind14 = jsonrpc2.NewError(-32001, "find document(s) error14")

var ErrFind15 = jsonrpc2.NewError(-32001, "find document(s) error15")

var ErrFind16 = jsonrpc2.NewError(-32001, "find document(s) error16")

var ErrFind17 = jsonrpc2.NewError(-32001, "find document(s) error17")

var ErrFind18 = jsonrpc2.NewError(-32001, "find document(s) error18")

var ErrFind19 = jsonrpc2.NewError(-32001, "find document(s) error19")
