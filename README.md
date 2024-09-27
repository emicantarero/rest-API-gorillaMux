
# REST API con Gorilla Mux

Este proyecto es una API REST básica construida con Go, utilizando el paquete Gorilla Mux para enrutar. La API permite realizar operaciones CRUD (Crear, Leer, Eliminar) sobre una base de datos de elementos en Postgres.

## Rutas disponibles

- **GET /users**: Obtiene todos los usuarios.
- **GET /users/{id}**: Obtiene un user por su ID.
- **POST /users**: Crea un nuevo usuario.
- **DELETE /users/{id}**: Elimina un user por ID.
- **GET /tasks**: Obtiene todos las tareas.
- **GET /tasks/{id}**: Obtiene una tarea por su ID.
- **POST /tasks**: Crea una nueva tarea.
- **DELETE /tasks/{id}**: Elimina una tarea por ID.

## Estructura del proyecto

- **main.go**: Punto de entrada que configura el servidor y las rutas.
- **routes/**: Manejadores para las rutas HTTP.
- **models/**: Definiciones de los datos utilizados en la API.
- **db/**: Conexión y gestión de la base de datos.

## Instalación

1. Clona el repositorio:
   ```bash
   git clone https://github.com/emicantarero/rest-API-gorillaMux.git
   ```
2. Dirígete al directorio del proyecto:
   ```bash
   cd rest-API-gorillaMux
   ```
3. Instala las dependencias necesarias:
   ```bash
   go mod download
   ```
4. Ejecuta la aplicación:
   ```bash
   go run main.go
   ```

La API correrá en `http://localhost:8080`.

## Tecnologías usadas

- **Gorilla Mux**: Para el manejo de rutas.
- **GORM**: ORM utilizado para interactuar con la base de datos.
- **Postgres**: Base de datos usada en este proyecto.

## Autor
Emilio Cantarero
