package interfaces

import "context"

// SQLHandler interface
type SQLHandler interface {
	Get(context.Context, interface{}, string, ...interface{}) error
}
