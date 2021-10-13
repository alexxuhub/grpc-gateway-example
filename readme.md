#### grpc gateway server

>grpc gateway server example 

>>start gateway server : 
> `make run`

>> http server
> >> get :
> `127.0.0.1:8080/api/v1/getCorpTicket?url=www.baidu.com&name=xxxx` <br>
> >> post: 127.0.0.1:8080/api/v1/postCorpTicket <br>
> you can see log in background <br>

>>grpc server
> >> see the proto file 

>> swagger file 
> >> use `make openapi` to generate swagger json file