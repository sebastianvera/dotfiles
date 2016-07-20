Vim�UnDo� �o=�KI��?A��/o�~�!r��   O                                   W���    _�                     :        ����                                                                                                                                                                                                                                                                                                                            :          <          V       W���     �   ;   =          	)�   :   <          +		grpctracer.SetZipkinServiceName(svcName),�   9   ;          #	grpctracer.InitGlobalZipkinTracer(5�_�                     :        ����                                                                                                                                                                                                                                                                                                                            :          <          V       W���    �               O   package main       import (   	"fmt"   	"net"   	"os"       !	"github.com/Finciero/grpctracer"   	"github.com/Finciero/log"    	"github.com/Finciero/mono/kong"    	"github.com/Finciero/mono/mono"    	"github.com/Finciero/utils/env"   )       const (   %	svcPortEnvVar  = "MONO_SERVICE_PORT"   	svcPortDefault = 13000       	svcAddressTmpl = ":%d"   	svcTagTmpl     = "%s:%s"   )       const (   &	kongHostEnvVar  = "KONG_SERVICE_HOST"   #	kongHostDefault = "192.168.99.100"       ,	kongPortEnvVar  = "KONG_SERVICE_PORT_ADMIN"   	kongPortDefault = 8001       	kongURLTmpl = "%s:%d"   )       var (   	svcName    string   	svcVersion string   	svcAddress string       	kongURL string       	logger log.Logger   )       func init() {   5	svcPort := env.GetInt(svcPortEnvVar, svcPortDefault)   7	svcTag := fmt.Sprintf(svcTagTmpl, svcName, svcVersion)   2	svcAddress = fmt.Sprintf(svcAddressTmpl, svcPort)       ;	kongHost := env.GetString(kongHostEnvVar, kongHostDefault)   8	kongPort := env.GetInt(kongPortEnvVar, kongPortDefault)   7	kongURL = fmt.Sprintf(kongURLTmpl, kongHost, kongPort)       	logger = log.NewContext(   		"kong_url", kongURL,   		"svc_tag", svcTag,   		"svc_address", svcAddress,   	)       &	// grpctracer.InitGlobalZipkinTracer(   .	// 	grpctracer.SetZipkinServiceName(svcName),   	// )   }       func main() {   *	lis, err := net.Listen("tcp", svcAddress)   	if err != nil {   		logger.Error(err)   		os.Exit(1)   	}       4	grpcServer := grpctracer.NewServerWithInterceptor()   C	mono.RegisterMonoServer(grpcServer, &server{kong.NewAPI(kongURL)})       "	logger.Info("event", "svc:start")   .	if err := grpcServer.Serve(lis); err != nil {   		logger.Error(err)   		os.Exit(1)   	}   !	logger.Info("event", "svc:stop")   }5��