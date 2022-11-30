# backpack-bcgow6-lucas-soria

## Ejercicio Clase C2 - QUALITY: La importancia de documentar - P2 - PG

### Metodo POST http://localhost:8080/clientes/crear

**Request and Response Body:**

~~~json
{
    "username": "usuario",
    "password": "123456",
    "nombre": "minombre",
    "apellido": "miapellido",
    "email": "algunmail@mail.com"
}
~~~

~~~mermaid
sequenceDiagram
    actor C as client;
    participant H as handler;
    participant S as service;
    participant R as repository;
    participant D as DB;
    C->>H: POST /clientes/crear
    activate C
    activate H
    Note RIGHT of C: {<br/>"username": "usuario",<br/>"password": "123456",<br/>"nombre": "minombre",<br/>"apellido": "miapellido",<br/>"email": "algunmail@mail.com"<br/>}
    H->>S: Save(ctx, client)
    activate S
    S->>R: Store(ctx, client)
    activate R
    R->>D: INSERT clients VALUES(?,?,?,?,?)
    activate D
    D-->>R: SQL result
    deactivate D
    R-->>S: (id, error)
    deactivate R
    S-->>H: (client, error)
    deactivate S
    alt OK
    H-->>C: (201, response)
    Note RIGHT of C: {<br/>"username": "usuario",<br/>"password": "123456",<br/>"nombre": "minombre",<br/>"apellido": "miapellido",<br/>"email": "algunmail@mail.com"<br/>}
    else conflict
    H-->>C: (409, response)
    else unprocesable entity
    H-->>C: (422, response)
    end
    deactivate H
    deactivate C
~~~

### Metodo DELETE http://localhost:8080/clientes/eliminar?username=xxxx

~~~mermaid
sequenceDiagram
    actor C as client;
    participant H as handler;
    participant S as service;
    participant R as repository;
    participant D as DB;
    C->>H: DELETE /clientes/eliminar?username=xxxx
    activate C
    activate H
    H->>S: Delete(ctx, username)
    activate S
    S->>R: Remove(ctx, username)
    activate R
    R->>D: DELETE FROM clients WHERE username=?;
    activate D
    D-->>R: SQL result
    deactivate D
    R-->>S: error
    deactivate R
    S-->>H: error
    deactivate S
    alt No content
    H-->>C: 204
    else Not Found
    H-->>C: 404
    end
    deactivate H
    deactivate C
~~~

### Metodo PUT http://localhost:8080/clientes/actualizar

**Request and Response Body:**
~~~json
{
    "username": "usuario",
    "password": "123456",
    "nombre": "minombre",
    "apellido": "miapellido",
    "email": "algunmail@mail.com"
}
~~~

~~~mermaid
sequenceDiagram
    actor C as client;
    participant H as handler;
    participant S as service;
    participant R as repository;
    participant D as DB;
    C->>H: PUT /clientes/actualizar
    activate C
    activate H
    Note RIGHT of C: {<br/>"username": "usuario",<br/>"password": "123456",<br/>"nombre": "minombre",<br/>"apellido": "miapellido",<br/>"email": "algunmail@mail.com"<br/>}
    H->>S: Update(ctx, client)
    activate S
    S->>R: Update(ctx, client)
    activate R
    R->>D: UPDATE clients SET nombre=?, password=?, ... WHERE username=?
    activate D
    D-->>R: SQL result
    deactivate D
    R-->>S: error
    deactivate R
    S-->>H: (client, error)
    deactivate S
    alt No content
    H-->>C: (200, response)
    Note RIGHT of C: {<br/>"username": "usuario",<br/>"password": "123456",<br/>"nombre": "minombre",<br/>"apellido": "miapellido",<br/>"email": "algunmail@mail.com"<br/>}
    else Not Found
    H-->>C: 404
    else Conflict
    H-->>C: 409
    else Unprocessable entity
    H-->>C: 422
    end
    deactivate H
    deactivate C
~~~

### Metodo GET http://localhost:8080/clientes/buscar?username=xxxx

**Response Body:**
~~~json
{
    "username": "usuario",
    "password": "123456",
    "nombre": "minombre",
    "apellido": "miapellido",
    "email": "algunmail@mail.com"
}
~~~

~~~mermaid
sequenceDiagram
    actor C as client;
    participant H as handler;
    participant S as service;
    participant R as repository;
    participant D as DB;
    C->>H: GET /clientes/buscar?username=xxxx
    activate C
    activate H
    H->>S: GetByUsername(ctx, username)
    activate S
    S->>R: Find(ctx, username)
    activate R
    R->>D: SELECT * FROM clients WHERE username=?
    activate D
    D-->>R: SQL result
    deactivate D
    R-->>S: (client, error)
    deactivate R
    S-->>H: (client, error)
    deactivate S
    alt No content
    H-->>C: (200, response)
    Note RIGHT of C: {<br/>"username": "usuario",<br/>"password": "123456",<br/>"nombre": "minombre",<br/>"apellido": "miapellido",<br/>"email": "algunmail@mail.com"<br/>}
    else Not Found
    H-->>C: 404
    end
    deactivate H
    deactivate C
~~~


### Metodo GET http://localhost:8080/clientes/listar

**Response Body:**
~~~json
[
    {
        "username": "usuario",
        "password": "123456",
        "nombre": "minombre",
        "apellido": "miapellido",
        "email": "algunmail@mail.com"
    },
    ...
]
~~~

~~~mermaid
sequenceDiagram
%%{init: { 'sequence': {'noteAlign': 'left'} }}%%
    actor C as client;
    participant H as handler;
    participant S as service;
    participant R as repository;
    participant D as DB;
    C->>H: GET /clientes/listar
    activate C
    activate H
    H->>S: GetAll(ctx)
    activate S
    S->>R: FindAll(ctx)
    activate R
    R->>D: SELECT * FROM clients
    activate D
    D-->>R: SQL result
    deactivate D
    R-->>S: ([]client, error)
    deactivate R
    S-->>H: ([]client, error)
    deactivate S
    H-->>C: (200, response)
    Note RIGHT of C: [<br/>{<br/>"username": "usuario",<br/>"password": "123456",<br/>"nombre": "minombre",<br/>"apellido": "miapellido",<br/>"email": "algunmail@mail.com"<br/>},<br/>...<br/>]
    deactivate H
    deactivate C
~~~
