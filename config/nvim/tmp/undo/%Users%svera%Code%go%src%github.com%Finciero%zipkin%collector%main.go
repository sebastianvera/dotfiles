Vim�UnDo� �Wpę��v+�(eo���ï^
�?�IOb�V��              ;                       WY�    _�                            ����                                                                                                                                                                                                                                                                                                                                                             WY��    �                  package main       import (   $	"github.com/spacemonkeygo/spacelog"   *	"gopkg.in/spacemonkeygo/monitor.v1/trace"   )       func main() {   6	spacelog.MustSetup("collector", spacelog.SetupConfig{   		Level:  "info",   <		Format: "{{ColorizeLevel .Level}}{{.Message}}{{.Reset}}"})       =	collector, err := trace.NewScribeCollector("localhost:9410")   	if err != nil {   		panic(err)   	}   	defer collector.Close()       :	panic(trace.RedirectPackets("127.0.0.1:8082", collector))   }5�_�                       -    ����                                                                                                                                                                                                                                                                                                                                                             WY�     �               =	collector, err := trace.NewScribeCollector("localhost:9410")5�_�                       5    ����                                                                                                                                                                                                                                                                                                                                                             WY�    �                  package main       import (   $	"github.com/spacemonkeygo/spacelog"   *	"gopkg.in/spacemonkeygo/monitor.v1/trace"   )       func main() {   6	spacelog.MustSetup("collector", spacelog.SetupConfig{   		Level:  "info",   <		Format: "{{ColorizeLevel .Level}}{{.Message}}{{.Reset}}"})       =	collector, err := trace.NewScribeCollector("127.0.0.1:9410")   	if err != nil {   		panic(err)   	}   	defer collector.Close()       :	panic(trace.RedirectPackets("127.0.0.1:8082", collector))   }5�_�                       -    ����                                                                                                                                                                                                                                                                                                                                                             WY�     �               =	collector, err := trace.NewScribeCollector("127.0.0.1:9410")5�_�                       5    ����                                                                                                                                                                                                                                                                                                                                                             WY�    �                  package main       import (   $	"github.com/spacemonkeygo/spacelog"   *	"gopkg.in/spacemonkeygo/monitor.v1/trace"   )       func main() {   6	spacelog.MustSetup("collector", spacelog.SetupConfig{   		Level:  "info",   <		Format: "{{ColorizeLevel .Level}}{{.Message}}{{.Reset}}"})       =	collector, err := trace.NewScribeCollector("localhost:9410")   	if err != nil {   		panic(err)   	}   	defer collector.Close()       :	panic(trace.RedirectPackets("127.0.0.1:8082", collector))   }5�_�                       7    ����                                                                                                                                                                                                                                                                                                                                                             WY�@     �               =	collector, err := trace.NewScribeCollector("localhost:9410")5�_�                       7    ����                                                                                                                                                                                                                                                                                                                                                             WY�A    �                  package main       import (   $	"github.com/spacemonkeygo/spacelog"   *	"gopkg.in/spacemonkeygo/monitor.v1/trace"   )       func main() {   6	spacelog.MustSetup("collector", spacelog.SetupConfig{   		Level:  "info",   <		Format: "{{ColorizeLevel .Level}}{{.Message}}{{.Reset}}"})       =	collector, err := trace.NewScribeCollector("localhost:7410")   	if err != nil {   		panic(err)   	}   	defer collector.Close()       :	panic(trace.RedirectPackets("127.0.0.1:8082", collector))   }5�_�      	                 ;    ����                                                                                                                                                                                                                                                                                                                                                             WY�$     �               =	collector, err := trace.NewScribeCollector("localhost:7410")5�_�      
           	      7    ����                                                                                                                                                                                                                                                                                                                                                             WY�-     �               C	collector, err := trace.NewScribeCollector("localhost:7410/spans")5�_�   	              
      7    ����                                                                                                                                                                                                                                                                                                                                                             WY�/    �                  package main       import (   $	"github.com/spacemonkeygo/spacelog"   *	"gopkg.in/spacemonkeygo/monitor.v1/trace"   )       func main() {   6	spacelog.MustSetup("collector", spacelog.SetupConfig{   		Level:  "info",   <		Format: "{{ColorizeLevel .Level}}{{.Message}}{{.Reset}}"})       C	collector, err := trace.NewScribeCollector("localhost:9410/spans")   	if err != nil {   		panic(err)   	}   	defer collector.Close()       :	panic(trace.RedirectPackets("127.0.0.1:8082", collector))   }5�_�   
                    ;    ����                                                                                                                                                                                                                                                                                                                                                             WY�     �               C	collector, err := trace.NewScribeCollector("localhost:9410/spans")5�_�                      ;    ����                                                                                                                                                                                                                                                                                                                                                             WY�     �               B	collector, err := trace.NewScribeCollector("localhost:9410spans")5�_�                       H    ����                                                                                                                                                                                                                                                                                                                                                             WY�     �               O	collector, err := trace.NewScribeCollector("localhost:9410/api/v1/spansspans")5�_�                        H    ����                                                                                                                                                                                                                                                                                                                                                             WY�    �                  package main       import (   $	"github.com/spacemonkeygo/spacelog"   *	"gopkg.in/spacemonkeygo/monitor.v1/trace"   )       func main() {   6	spacelog.MustSetup("collector", spacelog.SetupConfig{   		Level:  "info",   <		Format: "{{ColorizeLevel .Level}}{{.Message}}{{.Reset}}"})       J	collector, err := trace.NewScribeCollector("localhost:9410/api/v1/spans")   	if err != nil {   		panic(err)   	}   	defer collector.Close()       :	panic(trace.RedirectPackets("127.0.0.1:8082", collector))   }5�_�                       ;    ����                                                                                                                                                                                                                                                                                                                                                             WY�     �               I	collector, err := trace.NewScribeCollector("localhost:9410/api/v1/spans`       1 new message since 17:05   Mark as read (esc)Mark as read   Files   ")5��