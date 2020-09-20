package actors

import (
	"github.com/kooinam/fab.io/helpers"
)

// Context used to represent actor execution context with data
type Context struct {
	params *helpers.Dictionary
}

// makeContext use to instantiate controller context instance
func makeContext(params helpers.H) *Context {
	context := &Context{
		params: helpers.MakeDictionary(params),
	}

	return context
}

// Params used to retrieve params value
func (context *Context) Params(key string) interface{} {
	return context.params.Value(key)
}

// ParamsStr used to retrieve params value in string
func (context *Context) ParamsStr(key string) string {
	return context.params.ValueStr(key)
}

// ParamsInt used to retrieve params value in int
func (context *Context) ParamsInt(key string, fallback int) int {
	return context.params.ValueInt(key, fallback)
}

// ParamsFloat64 used to retrieve params value in float64
func (context *Context) ParamsFloat64(key string, fallback float64) float64 {
	return context.params.ValueFloat64(key, fallback)
}

// ParamsBool used to retrieve params value in bool
func (context *Context) ParamsBool(key string) bool {
	return context.params.ValueBool(key)
}
