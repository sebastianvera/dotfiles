Vim�UnDo� �*a� y��=��M����]�Q�~���   =                                  WxY�    _�                             ����                                                                                                                                                                                                                                                                                                                                                             WxY4     �                   �               5�_�                            ����                                                                                                                                                                                                                                                                                                                                                             WxY5     �          /       �          .    5�_�                            ����                                                                                                                                                                                                                                                                                                                                                             WxY8    �               /   package handler       // ContextHandler ...   type ContextHandler interface {   N	ServeHTTPContext(ctx context.Context, w http.ResponseWriter, r *http.Request)   }       // ContextHandlerFunc ...   Qtype ContextHandlerFunc func(context.Context, http.ResponseWriter, *http.Request)       // ServeHTTPContext ...   nfunc (h ContextHandlerFunc) ServeHTTPContext(ctx context.Context, rw http.ResponseWriter, req *http.Request) {   	h(ctx, rw, req)   }       // Adapter ...   type Adapter struct {   	logger  log.Logger   	handler ContextHandler   }       Efunc (c *Adapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {   *	requestID := r.Header.Get("X-Request-Id")   1	logger := c.logger.With("request_id", requestID)   	ctx := context.WithValue(   		context.Background(),   		"logger",   			logger,   	)    	logger.Info("request", "start")   $	defer logger.Info("request", "end")   &	c.handler.ServeHTTPContext(ctx, w, r)   }       // Handler ...   type Handler struct {   	Logger log.Logger   }       // Wrap ...   =func (h *Handler) Wrap(chf ContextHandlerFunc) http.Handler {   	return &Adapter{   		logger:  h.Logger,   		handler: chf,   	}   }    5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             WxY=     �         5      	"log"5�_�                            ����                                                                                                                                                                                                                                                                                                                                                             WxYA    �               5   package handler       import (   	"github.com/Finciero/log"   	"net/http"       	"golang.org/x/net/context"   )       // ContextHandler ...   type ContextHandler interface {   N	ServeHTTPContext(ctx context.Context, w http.ResponseWriter, r *http.Request)   }       // ContextHandlerFunc ...   Qtype ContextHandlerFunc func(context.Context, http.ResponseWriter, *http.Request)       // ServeHTTPContext ...   nfunc (h ContextHandlerFunc) ServeHTTPContext(ctx context.Context, rw http.ResponseWriter, req *http.Request) {   	h(ctx, rw, req)   }       // Adapter ...   type Adapter struct {   	logger  log.Logger   	handler ContextHandler   }       Efunc (c *Adapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {   *	requestID := r.Header.Get("X-Request-Id")   1	logger := c.logger.With("request_id", requestID)   	ctx := context.WithValue(   		context.Background(),   		"logger",   			logger,   	)    	logger.Info("request", "start")   $	defer logger.Info("request", "end")   &	c.handler.ServeHTTPContext(ctx, w, r)   }       // Handler ...   type Handler struct {   	Logger log.Logger   }       // Wrap ...   =func (h *Handler) Wrap(chf ContextHandlerFunc) http.Handler {   	return &Adapter{   		logger:  h.Logger,   		handler: chf,   	}   }5�_�                            ����                                                                                                                                                                                                                                                                                                                                                             WxYH    �               6   package handler       import (   	"net/http"       	"github.com/Finciero/log"       	"golang.org/x/net/context"   )       // ContextHandler ...   type ContextHandler interface {   N	ServeHTTPContext(ctx context.Context, w http.ResponseWriter, r *http.Request)   }       // ContextHandlerFunc ...   Qtype ContextHandlerFunc func(context.Context, http.ResponseWriter, *http.Request)       // ServeHTTPContext ...   nfunc (h ContextHandlerFunc) ServeHTTPContext(ctx context.Context, rw http.ResponseWriter, req *http.Request) {   	h(ctx, rw, req)   }       // Adapter ...   type Adapter struct {   	logger  log.Logger   	handler ContextHandler   }       Efunc (c *Adapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {   *	requestID := r.Header.Get("X-Request-Id")   1	logger := c.logger.With("request_id", requestID)   	ctx := context.WithValue(   		context.Background(),   		"logger",   			logger,   	)    	logger.Info("request", "start")   $	defer logger.Info("request", "end")   &	c.handler.ServeHTTPContext(ctx, w, r)   }       // Handler ...   type Handler struct {   	Logger log.Logger   }       // Wrap ...   =func (h *Handler) Wrap(chf ContextHandlerFunc) http.Handler {   	return &Adapter{   		logger:  h.Logger,   		handler: chf,   	}   }5�_�                            ����                                                                                                                                                                                                                                                                                                                                                             WxY~     �         7       �         6    5�_�      	                     ����                                                                                                                                                                                                                                                                                                                                                             WxY�     �         9       �         8    5�_�      
           	          ����                                                                                                                                                                                                                                                                                                                                                             WxY�     �          9      func New5�_�   	              
          ����                                                                                                                                                                                                                                                                                                                                                             WxY�     �          9      func New(logger log.Logger)5�_�   
                        ����                                                                                                                                                                                                                                                                                                                                                             WxY�     �          9      #func New(logger log.Logger) Handler5�_�                       $    ����                                                                                                                                                                                                                                                                                                                                                             WxY�     �       "   =      		�       "   <    �      "   ;      	�      !   :    �      !   9      $func New(logger log.Logger) *Handler5�_�                    !       ����                                                                                                                                                                                                                                                                                                                                                             WxY�    �               =   package handler       import (   	"net/http"       	"github.com/Finciero/log"       	"golang.org/x/net/context"   )       // ContextHandler ...   type ContextHandler interface {   N	ServeHTTPContext(ctx context.Context, w http.ResponseWriter, r *http.Request)   }       // ContextHandlerFunc ...   Qtype ContextHandlerFunc func(context.Context, http.ResponseWriter, *http.Request)       // ServeHTTPContext ...   nfunc (h ContextHandlerFunc) ServeHTTPContext(ctx context.Context, rw http.ResponseWriter, req *http.Request) {   	h(ctx, rw, req)   }       // Adapter ...   type Adapter struct {   	logger  log.Logger   	handler ContextHandler   }       
// New ...   &func New(logger log.Logger) *Handler {   	return &Handler{   		logger: logger,   	}   }       Efunc (c *Adapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {   *	requestID := r.Header.Get("X-Request-Id")   1	logger := c.logger.With("request_id", requestID)   	ctx := context.WithValue(   		context.Background(),   		"logger",   			logger,   	)    	logger.Info("request", "start")   $	defer logger.Info("request", "end")   &	c.handler.ServeHTTPContext(ctx, w, r)   }       // Handler ...   type Handler struct {   	Logger log.Logger   }       // Wrap ...   =func (h *Handler) Wrap(chf ContextHandlerFunc) http.Handler {   	return &Adapter{   		logger:  h.Logger,   		handler: chf,   	}   }5�_�                    !       ����                                                                                                                                                                                                                                                                                                                                                             WxY�     �       "   =      		logger: logger,5�_�                     !       ����                                                                                                                                                                                                                                                                                                                                                             WxY�    �               =   package handler       import (   	"net/http"       	"github.com/Finciero/log"       	"golang.org/x/net/context"   )       // ContextHandler ...   type ContextHandler interface {   N	ServeHTTPContext(ctx context.Context, w http.ResponseWriter, r *http.Request)   }       // ContextHandlerFunc ...   Qtype ContextHandlerFunc func(context.Context, http.ResponseWriter, *http.Request)       // ServeHTTPContext ...   nfunc (h ContextHandlerFunc) ServeHTTPContext(ctx context.Context, rw http.ResponseWriter, req *http.Request) {   	h(ctx, rw, req)   }       // Adapter ...   type Adapter struct {   	logger  log.Logger   	handler ContextHandler   }       
// New ...   &func New(logger log.Logger) *Handler {   	return &Handler{   		Logger: logger,   	}   }       Efunc (c *Adapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {   *	requestID := r.Header.Get("X-Request-Id")   1	logger := c.logger.With("request_id", requestID)   	ctx := context.WithValue(   		context.Background(),   		"logger",   			logger,   	)    	logger.Info("request", "start")   $	defer logger.Info("request", "end")   &	c.handler.ServeHTTPContext(ctx, w, r)   }       // Handler ...   type Handler struct {   	Logger log.Logger   }       // Wrap ...   =func (h *Handler) Wrap(chf ContextHandlerFunc) http.Handler {   	return &Adapter{   		logger:  h.Logger,   		handler: chf,   	}   }5��