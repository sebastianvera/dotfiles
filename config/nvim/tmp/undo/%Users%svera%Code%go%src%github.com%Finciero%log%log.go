Vim�UnDo� ��N=�J�!���_�S��[���]Fϑ�   M                                  W|'�    _�                     >       ����                                                                                                                                                                                                                                                                                                                                                             W|'�     �   =   ?   M      		desc = val.Description5�_�                     >       ����                                                                                                                                                                                                                                                                                                                                                             W|'�    �               M   package log       $// TODO(jaguirre): add documentation       import (   	"os"       	"github.com/Finciero/errors"   #	kitlog "github.com/go-kit/kit/log"   )       // Logger ...   type Logger interface {   $	Log(keyvalues ...interface{}) error   %	Info(keyvalues ...interface{}) error   %	Warn(keyvalues ...interface{}) error   1	Error(err error, keyvalues ...interface{}) error       &	With(keyvalues ...interface{}) Logger   }       // Context ...   type Context struct {   	kitlog.Context   }       // NewContext ...   4func NewContext(keyvalues ...interface{}) *Context {   *	logger := kitlog.NewJSONLogger(os.Stderr)   ?	return &Context{*kitlog.NewContext(logger).With(keyvalues...)}   }       // NewRequestContext ...   Mfunc NewRequestContext(requestID string, keyvalues ...interface{}) *Context {   *	logger := kitlog.NewJSONLogger(os.Stderr)   4	ctx := kitlog.NewContext(logger).With(keyvalues...)   4	return &Context{*ctx.With("request_id", requestID)}   }       M// With returns a new Context with keyvals appended to those of the receiver.   9func (ctx *Context) With(keyvals ...interface{}) Logger {   /	return &Context{*ctx.Context.With(keyvals...)}   }       // Info ...   8func (ctx *Context) Info(keyvals ...interface{}) error {   9	return ctx.Context.With("level", "info").Log(keyvals...)   }       // Error ...   Dfunc (ctx *Context) Error(err error, keyvals ...interface{}) error {   	if err == nil {   		return nil   	}       	var desc string       *	if val, ok := (err).(*errors.Error); ok {   		for k, v := range val.Meta {   "			keyvals = append(keyvals, k, v)   		}   		desc = val.Message   		} else {   		desc = err.Error()   	}       	if len(desc) > 0 {   )		keyvals = append(keyvals, "desc", desc)   	}       :	return ctx.Context.With("level", "error").Log(keyvals...)   }       // Warn ...   8func (ctx *Context) Warn(keyvals ...interface{}) error {   <	return ctx.Context.With("level", "warning").Log(keyvals...)   }5��