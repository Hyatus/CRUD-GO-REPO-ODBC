# API de Usuarios (Go + Gin + ODBC)

Este proyecto es una API RESTful para la gestión de usuarios, desarrollada en Go utilizando el framework Gin y el driver ODBC para conectarse a MySQL (o cualquier base de datos compatible con ODBC).

## Características
- CRUD de usuarios (Crear, Leer, Actualizar, Eliminar)
- Arquitectura modular (controladores, servicios, modelos, rutas)
- Conexión a base de datos MySQL (o cualquier base compatible con ODBC)

## Estructura del Proyecto
```
Experimento/
├── cmd/           # Punto de entrada de la aplicación (main.go)
│   └── main.go
├── config/        # Configuración de la app y conexión a la base de datos
│   └── config.go
│   └── db.go
├── controllers/   # Lógica para manejar las peticiones HTTP (controladores)
│   └── user_controller.go
├── models/        # Definición de las estructuras de datos (modelos)
│   └── user.go
├── routes/        # Definición de rutas y asociación con controladores
│   └── routes.go
├── services/      # Lógica de negocio y acceso a datos
│   └── user_service.go
├── storage/       # Implementaciones de almacenamiento alternativo o temporal
│   └── memory.go
├── go.mod         # Dependencias del proyecto
├── go.sum         # Checksum de dependencias
```

## Instalación
1. Clona el repositorio:
   ```sh
   git clone <url-del-repo>
   cd Experimento
   ```
2. Instala las dependencias:
   ```sh
   go mod tidy
   ```
3. Configura la base de datos en `config/db.go` según tus credenciales de MySQL.
4. Ejecuta la migración de modelos si es necesario.

## Uso
Para iniciar el servidor:
```sh
go run cmd/main.go
```

## Endpoints Principales

- `GET /users` — Lista todos los usuarios
- `GET /users/:id` — Obtiene un usuario por ID
- `POST /users` — Crea un nuevo usuario
- `PUT /users/:id` — Actualiza un usuario existente
- `DELETE /users/:id` — Elimina un usuario


## Ejemplos de Peticiones en Postman

### Crear usuario (POST /users)
```
POST http://localhost:8080/users
Content-Type: application/json

{
   "name": "Juan Pérez",
   "age": 30,
   "email": "juan@example.com"
}
```

### Obtener todos los usuarios (GET /users)
```
GET http://localhost:8080/users
```

### Obtener usuario por ID (GET /users/:id)
```
GET http://localhost:8080/users/1
```

### Actualizar usuario (PUT /users/:id)
```
PUT http://localhost:8080/users/1
Content-Type: application/json

{
   "name": "Juan Actualizado",
   "age": 31,
   "email": "juan.actualizado@example.com"
}
```

### Eliminar usuario (DELETE /users/:id)
```
DELETE http://localhost:8080/users/1
```

## Dependencias principales
- [Gin](https://github.com/gin-gonic/gin) — Framework web para Go
- [ODBC Driver (alexbrainman/odbc)](https://github.com/alexbrainman/odbc) — Driver ODBC para Go

## Notas
- Asegúrate de tener una instancia de MySQL corriendo y accesible.
- Modifica las credenciales de la base de datos en `config/db.go` según tu entorno.

## Licencia
MIT
