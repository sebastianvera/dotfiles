Vim�UnDo� �� �q�@�X���Ϻ�4:�����0|IbNu�(   �   +      expect(reply).to.have.been.calledOnce   [   '                       W"�F    _�                     O        ����                                                                                                                                                                                                                                                                                                                            O          W          V       W"�    �   N   \   �    �   O   P   �    �   N   O       	         const reply = sinon.spy()         const req = {           headers: {},           url: {             path: '/foo/',   
        },         }   &      apiVersionMiddleware(req, reply)   <      expect(reply).to.have.been.calledWith(Boom.notFound())5�_�                    O        ����                                                                                                                                                                                                                                                                                                                            O          O          V       W"�D     �   N   Q   �    �   O   P   �    �   N   O                const reply = sinon.spy()5�_�                    Z        ����                                                                                                                                                                                                                                                                                                                            Z          [          V       W"�K     �   Y   \   �    �   Z   [   �    �   Y   Z          &      apiVersionMiddleware(req, reply)   <      expect(reply).to.have.been.calledWith(Boom.notFound())5�_�                    \        ����                                                                                                                                                                                                                                                                                                                            Z           [   (       V       W"�L    �   [   \           5�_�                   _       ����                                                                                                                                                                                                                                                                                                                            a   
       a   
       V   
    W"��    �   ^   `   �      E  describe("when a request have a wrong have accept headers", () => {�   _   `   �    5�_�                    o       ����                                                                                                                                                                                                                                                                                                                            a   
       a   
       V   
    W"��    �   n   p   �      K  describe("when a request have a wrong version in accept headers", () => {�   o   p   �    5�_�                    >        ����                                                                                                                                                                                                                                                                                                                            >           >           V        W"��     �   =   @   �    �   >   ?   �    �   =   >                const reply = sinon.spy()5�_�                    J   '    ����                                                                                                                                                                                                                                                                                                                            >           ?   -       V        W"��     �   I   K   �      <      expect(reply).to.have.been.calledWith(Boom.notFound())5�_�                    J   '    ����                                                                                                                                                                                                                                                                                                                            >           ?   -       V        W"��    �   I   K   �      +      expect(reply).to.have.been.calledonce5�_�                    J       ����                                                                                                                                                                                                                                                                                                                            >           ?   -       V        W"��    �   I   K   �      +      expect(reply).to.have.been.calledOnce5�_�                    <       ����                                                                                                                                                                                                                                                                                                                            >           ?   -       V        W"��   	 �   ;   =   �      N  describe("when is an options request and have wrong accept headers", () => {�   <   =   �    5�_�                    =       ����                                                                                                                                                                                                                                                                                                                            >           ?   -       V        W"��   
 �   <   >   �      1    it('should return a 404 status code', () => {5�_�                    O       ����                                                                                                                                                                                                                                                                                                                            >           ?   -       V        W"��    �   N   P   �      1    it('should return a 404 status code', () => {5�_�                    V       ����                                                                                                                                                                                                                                                                                                                            >           ?   -       V        W"�     �   U   W   �              method: 'options',5�_�                    \       ����                                                                                                                                                                                                                                                                                                                            >           ?   -       V        W"�     �   [   ]   �      )      expect(spy).to.have.been.calledOnce5�_�                    P       ����                                                                                                                                                                                                                                                                                                                            Q          P          V       W"�     �   O   Q   �    �   P   Q   �    �   O   P          *      const reply = { continue: () => {} }   .      const spy = sinon.spy(reply, 'continue')5�_�                    [       ����                                                                                                                                                                                                                                                                                                                            P           P          V       W"�      �   Z   \   �      -      expect(spy).not.to.have.been.calledOnce5�_�                    [       ����                                                                                                                                                                                                                                                                                                                            P           P          V       W"�%    �   Z   \   �      /      expect(reply).not.to.have.been.calledOnce5�_�                    O       ����                                                                                                                                                                                                                                                                                                                            P           P          V       W"�1    �   N   P   �      &    it('should call continue', () => {5�_�                    [   '    ����                                                                                                                                                                                                                                                                                                                            P           P          V       W"�=     �   Z   \   �      +      expect(reply).to.have.been.calledOnce5�_�                     [   ;    ����                                                                                                                                                                                                                                                                                                                            P           P          V       W"�E    �   Z   \   �      ;      expect(reply).to.have.been.calledWith(Boom.notFound()5�_�                   a   
    ����                                                                                                                                                                                                                                                                                                                            a   
       a   
       V   
    W"�x     �   a   b   �      *      const reply = { continue: () => {} }   .      const spy = sinon.spy(reply, 'continue')5�_�                    a       ����                                                                                                                                                                                                                                                                                                                            a   
       a   
       V   
    W"�z     �   `   b        5�_�                    k       ����                                                                                                                                                                                                                                                                                                                            a   
       a   
       V   
    W"��     �   k   l   �      &      apiVersionMiddleware(req, reply)   )      expect(spy).to.have.been.calledOnce5�_�      	              l       ����                                                                                                                                                                                                                                                                                                                            a   
       a   
       V   
    W"��     �   k   m        5�_�      
           	   m       ����                                                                                                                                                                                                                                                                                                                            a   
       a   
       V   
    W"��    �   l   n        5�_�   	               
   _       ����                                                                                                                                                                                                                                                                                                                            a   
       a   
       V   
    W"��    �   _   `   �    �   ^   `   �      E  describe('when a request have a wrong have accept headers', () => {5��