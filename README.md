# Yofio
Yofio Prueba Tecnica

Instalar dependencias requeridas para GO 

```sh
go get -u github.com/gorilla/mux
go get -u github.com/go-sql-driver/mysql
```

Configurar el ambiente de la base de datos, usando los scripts que estan en la carperta /db

- db.sql
- user.sql

Luego de esto hacer ejecutable el script de bash usando el comando

```
chmod +x run.sh
```

Este script agrega las variables de entorno que contienen los datos de conexión a la base de datos y el puerto en el cual el api.

Luego ejecutar el script run.sh

```
./run.sh
```

Mostrara un mensaje notificando si el servidor esta abierto escuchando peticiones en el puerto configurado.

# Pruebas
Luego de esto probar peticiones al servidor usando curl o postman.

## /
```sh
curl --location --request GET 'localhost:8000/'
```

## /credit-assignment
```sh
curl --location --request GET 'localhost:8000/' \
--header 'Content-Type: text/plain' \
--data-raw '{"investment": 3000}'
```

## /statistics
```sh
curl --location --request POST 'localhost:8000/statistics'
```
