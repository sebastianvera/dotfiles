Vim�UnDo� �L��n��Q�i�� �4K�d�"OLOw��G��   $                                  Wy�.    _�                            ����                                                                                                                                                                                                                                                                                                                                                             Wy��     �                   �               5�_�                            ����                                                                                                                                                                                                                                                                                                                                                             Wy��     �                 5�_�                            ����                                                                                                                                                                                                                                                                                                                                                             Wy��     �                9func (n *noopLogger) Warn(keyvals ...interface{}) error {�                Efunc (n *noopLogger) Error(err error, keyvals ...interface{}) error {�                9func (n *noopLogger) Info(keyvals ...interface{}) error {�   
             8func (n *noopLogger) Log(keyvals ...interface{}) error {�                >func (n *noopLogger) With(keyvals ...interface{}) log.Logger {�                type noopLogger struct{}5�_�                            ����                                                                                                                                                                                                                                                                                                                                                             Wy��     �                �             5�_�                            ����                                                                                                                                                                                                                                                                                                                                                             Wy�     �                9func (n *NoopLogger) Warn(keyvals ...interface{}) error {�                Efunc (n *NoopLogger) Error(err error, keyvals ...interface{}) error {�                9func (n *NoopLogger) Info(keyvals ...interface{}) error {�                8func (n *NoopLogger) Log(keyvals ...interface{}) error {�                >func (n *NoopLogger) With(keyvals ...interface{}) log.Logger {�                type NoopLogger struct{}�                // NoopLogger 5�_�                            ����                                                                                                                                                                                                                                                                                                                                                             Wy�    �                  package log       // NOOPLogger    type NOOPLogger struct{}       M// With returns a new Context with keyvals appended to those of the receiver.   >func (n *NOOPLogger) With(keyvals ...interface{}) log.Logger {   		return n   }       
// Log ...   8func (n *NOOPLogger) Log(keyvals ...interface{}) error {   	return nil   }       // Info ...   9func (n *NOOPLogger) Info(keyvals ...interface{}) error {   	return nil   }       // Error ...   Efunc (n *NOOPLogger) Error(err error, keyvals ...interface{}) error {   	return nil   }       // Warn ...   9func (n *NOOPLogger) Warn(keyvals ...interface{}) error {   	return nil   }    5�_�      	                     ����                                                                                                                                                                                                                                                                                                                                                             Wy�     �               // NOOPLogger5�_�      
           	          ����                                                                                                                                                                                                                                                                                                                                                             Wy�    �                  package log       import "log"       // NOOPLogger ...   type NOOPLogger struct{}       M// With returns a new Context with keyvals appended to those of the receiver.   >func (n *NOOPLogger) With(keyvals ...interface{}) log.Logger {   		return n   }       
// Log ...   8func (n *NOOPLogger) Log(keyvals ...interface{}) error {   	return nil   }       // Info ...   9func (n *NOOPLogger) Info(keyvals ...interface{}) error {   	return nil   }       // Error ...   Efunc (n *NOOPLogger) Error(err error, keyvals ...interface{}) error {   	return nil   }       // Warn ...   9func (n *NOOPLogger) Warn(keyvals ...interface{}) error {   	return nil   }5�_�   	              
          ����                                                                                                                                                                                                                                                                                                                                                             Wy��     �      	         M// With returns a new Context with keyvals appended to those of the receiver.5�_�   
                        ����                                                                                                                                                                                                                                                                                                                                                             Wy��     �      	         // With returns a NOOPLogger5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             Wy��     �      	         // With 5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             Wy��     �               
// Log ...5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             Wy��     �               // Info ...5�_�                       	    ����                                                                                                                                                                                                                                                                                                                                                             Wy��     �               // Error ...5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             Wy��     �               // Warn ...5�_�                       #    ����                                                                                                                                                                                                                                                                                                                                                             Wy��    �                  package log       import "log"       // NOOPLogger ...   type NOOPLogger struct{}       (// With does nothing and returns it self   >func (n *NOOPLogger) With(keyvals ...interface{}) log.Logger {   		return n   }       #// Log does nothing and returns nil   8func (n *NOOPLogger) Log(keyvals ...interface{}) error {   	return nil   }       $// Info does nothing and returns nil   9func (n *NOOPLogger) Info(keyvals ...interface{}) error {   	return nil   }       %// Error does nothing and returns nil   Efunc (n *NOOPLogger) Error(err error, keyvals ...interface{}) error {   	return nil   }       $// Warn does nothing and returns nil   9func (n *NOOPLogger) Warn(keyvals ...interface{}) error {   	return nil   }5�_�                            ����                                                                                                                                                                                                                                                                                                                                                             Wy��     �      	           �             5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             Wy�     �         !    �      	   !    5�_�                    	        ����                                                                                                                                                                                                                                                                                                                                                             Wy�
     �      	          // NewContext ...5�_�                            ����                                                                                                                                                                                                                                                                                                                                                             Wy�
     �      	   %    �      	   %    5�_�                    
        ����                                                                                                                                                                                                                                                                                                                                                             Wy�     �   	      &      4func NewContext(keyvalues ...interface{}) *Context {5�_�                    
        ����                                                                                                                                                                                                                                                                                                                                                             Wy�     �   	      &      /NewContext(keyvalues ...interface{}) *Context {5�_�                    
        ����                                                                                                                                                                                                                                                                                                                                                             Wy�     �      
   &      func NewNOOPContext()   %(keyvalues ...interface{}) *Context {�   	      &      %(keyvalues ...interface{}) *Context {5�_�                    
       ����                                                                                                                                                                                                                                                                                                                                                             Wy�     �   	   
          *	logger := kitlog.NewJSONLogger(os.Stdout)5�_�                    
       ����                                                                                                                                                                                                                                                                                                                                                             Wy�     �   	      $      ?	return &Context{*kitlog.NewContext(logger).With(keyvalues...)}5�_�                    
       ����                                                                                                                                                                                                                                                                                                                                                             Wy�    �               $   package log       import "log"       // NOOPLogger ...   type NOOPLogger struct{}       // NewContext ...   8func NewNOOPContext(keyvalues ...interface{}) *Context {   	return &NOOPLogger{}   }       (// With does nothing and returns it self   >func (n *NOOPLogger) With(keyvals ...interface{}) log.Logger {   		return n   }       #// Log does nothing and returns nil   8func (n *NOOPLogger) Log(keyvals ...interface{}) error {   	return nil   }       $// Info does nothing and returns nil   9func (n *NOOPLogger) Info(keyvals ...interface{}) error {   	return nil   }       %// Error does nothing and returns nil   Efunc (n *NOOPLogger) Error(err error, keyvals ...interface{}) error {   	return nil   }       $// Warn does nothing and returns nil   9func (n *NOOPLogger) Warn(keyvals ...interface{}) error {   	return nil   }5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             Wy�     �      	   $      // NewContext ...5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             Wy�!    �               $   package log       import "log"       // NOOPLogger ...   type NOOPLogger struct{}       // NewNOOPContext ...   8func NewNOOPContext(keyvalues ...interface{}) *Context {   	return &NOOPLogger{}   }       (// With does nothing and returns it self   >func (n *NOOPLogger) With(keyvals ...interface{}) log.Logger {   		return n   }       #// Log does nothing and returns nil   8func (n *NOOPLogger) Log(keyvals ...interface{}) error {   	return nil   }       $// Info does nothing and returns nil   9func (n *NOOPLogger) Info(keyvals ...interface{}) error {   	return nil   }       %// Error does nothing and returns nil   Efunc (n *NOOPLogger) Error(err error, keyvals ...interface{}) error {   	return nil   }       $// Warn does nothing and returns nil   9func (n *NOOPLogger) Warn(keyvals ...interface{}) error {   	return nil   }5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             Wy�"     �      	   $      // NewNOOPContext ...5�_�                        :    ����                                                                                                                                                                                                                                                                                                                                                             Wy�-    �               $   package log       import "log"       // NOOPLogger ...   type NOOPLogger struct{}       ;// NewNOOPContext returns a NOOPLogger, perfect for testing   8func NewNOOPContext(keyvalues ...interface{}) *Context {   	return &NOOPLogger{}   }       (// With does nothing and returns it self   >func (n *NOOPLogger) With(keyvals ...interface{}) log.Logger {   		return n   }       #// Log does nothing and returns nil   8func (n *NOOPLogger) Log(keyvals ...interface{}) error {   	return nil   }       $// Info does nothing and returns nil   9func (n *NOOPLogger) Info(keyvals ...interface{}) error {   	return nil   }       %// Error does nothing and returns nil   Efunc (n *NOOPLogger) Error(err error, keyvals ...interface{}) error {   	return nil   }       $// Warn does nothing and returns nil   9func (n *NOOPLogger) Warn(keyvals ...interface{}) error {   	return nil   }5�_�                             ����                                                                                                                                                                                                                                                                                                                                                             Wy��     �                 urfave5��