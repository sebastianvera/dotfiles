Vim�UnDo� ��=D��y�L 2��IK؆����m�N�It�L                    	       	   	   	    WJ|_    _�                             ����                                                                                                                                                                                                                                                                                                                                                             WJ|6     �                   5�_�                            ����                                                                                                                                                                                                                                                                                                                                                             WJ|9     �                  -me "*.t1" -exec rename 's/\.t1$/.t2/' '{}' \;5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             WJ|<     �                  7find . -name "*.t1" -exec rename 's/\.t1$/.t2/' '{}' \;5�_�                       &    ����                                                                                                                                                                                                                                                                                                                                                             WJ|>     �                  7find . -name "*.js" -exec rename 's/\.t1$/.t2/' '{}' \;5�_�                       ,    ����                                                                                                                                                                                                                                                                                                                                                             WJ|@    �                  7find . -name "*.js" -exec rename 's/\.js$/.t2/' '{}' \;5�_�                            ����                                                                                                                                                                                                                                                                                                                                                             WJ|R     �                  7find . -name "*.js" -exec rename 's/\.js$/.ts/' '{}' \;5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             WJ|W     �                  Vfind . -name "*.t1" -exec bash -c 'mv "$1" "$(sed "s/\.t1$/.t2/" <<< "$1")"' - '{}' \;5�_�      	                 7    ����                                                                                                                                                                                                                                                                                                                                                             WJ|\     �                  Vfind . -name "*.js" -exec bash -c 'mv "$1" "$(sed "s/\.t1$/.t2/" <<< "$1")"' - '{}' \;5�_�                  	      =    ����                                                                                                                                                                                                                                                                                                                                                             WJ|^    �                  Vfind . -name "*.js" -exec bash -c 'mv "$1" "$(sed "s/\.js$/.t2/" <<< "$1")"' - '{}' \;5��