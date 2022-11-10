## README.md

Primero necesitamos el archivo docker-compose.yml 

```yml
services:
  dynamodb-local:
    image: amazon/dynamodb-local:latest
    container_name: dynamodb-local
    ports:
      - "8000:8000"
    command: "-jar DynamoDBLocal.jar -sharedDb -dbPath ."
  dynamodb-admin:
    image: aaronshaf/dynamodb-admin
    ports:
      - "8001:8001"
    environment:
      DYNAMO_ENDPOINT: "http://dynamodb-local:8000"
    depends_on:
      - dynamodb-local
```

Luego necesitamos agregar la liblioteca que nos permita trabajar desde go con dynamodb

```go
go get github.com/aws/aws-sdk-go/aws
```

Es probable que necesite realizar ```go mod tidy``` para las dependencias que no se hayan descargado.


Json para el armado de la requests [POST/PATCH]:

```json
{
	"firstname":"gregory",
	"lastname":"house",
	"username":"dr. house",
	"password":"asd%Sara#sadf",
	"email":"gregory@digitalhouse.com",
	"ip":"134.532.234.12",
	"macAddress":"AD:DF:FD:fDF:DF:FD",
	"website":"drhouse.com",
	"image":"imagerandom"
}
```