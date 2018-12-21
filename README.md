# ExemploTcpGolang
Exemplo básico de Sockets TCP em Golang
# melhorias
-Criar API client
-Criar API Server  
-A partir do protocolo SMQP bolamos  TMQServer (Trivial Management Query Server)
-terá um conndispatcher para associar solitições de clientes com as gorotines 
-A aplicação vai implementar um handshake em 3 vias;
    -Cliente: envia "TMQP ver01\n"
    -Server: responde "Start\n"
    Cliente: responde "Start OK\n"
-implementação de handshake 