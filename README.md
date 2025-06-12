# 🧪 User API - Backend CRUD de Usuarios

Esta es una API RESTful desarrollada en **Go** que permite gestionar usuarios (crear, obtener, actualizar y eliminar). La aplicación está estructurada con **arquitectura por capas**, siguiendo principios SOLID, y documentada con **Swagger**. La base de datos utilizada es **MongoDB**, y el despliegue local se hace mediante **Docker**.

---

## 🎯 Objetivo

Implementar una API que permita realizar operaciones CRUD sobre usuarios:

- Crear un usuario
- Listar todos los usuarios
- Obtener un usuario por ID
- Actualizar un usuario
- Eliminar un usuario

---

## 🚀 Tecnologías usadas

- [Go](https://golang.org/) — Lenguaje principal
- [Gin](https://github.com/gin-gonic/gin) — Framework HTTP
- [MongoDB](https://www.mongodb.com/) — Base de datos NoSQL
- [Swagger (swaggo)](https://github.com/swaggo/gin-swagger) — Documentación interactiva
- [Docker](https://www.docker.com/products/docker-desktop) — Contenedores para despliegue
- Arquitectura limpia y principios **SOLID**

---

## 🛠️ Instrucciones de instalación

### 1. Clonar el repositorio

```bash
git clone https://github.com/Josuehmz/user-api.git
cd user-api
```

### 2. Configurar el entorno

Crear un archivo `.env` en la raíz del proyecto con el siguiente contenido:

```env
MONGO_URI=mongodb+srv://<usuario>:<password>@<cluster>.mongodb.net/?retryWrites=true&w=majority
DB_NAME=userdb
DB_COLLECTION=users
```
*Nota: El archivo `.env` con los valores de las variables para la conexión, se comparten por correo*
### 3. Instalar Docker Desktop

Asegúrate de tener instalado Docker:  
📦 https://www.docker.com/products/docker-desktop

### 4. Construir y ejecutar el contenedor

```bash
docker build -t user-api .
docker run -p 8080:8080 --env-file .env user-api
```

---

## 📘 Documentación Swagger

Una vez el contenedor esté corriendo, puedes acceder a la documentación interactiva en:  
📄 [`http://localhost:8080/swagger/index.html`](http://localhost:8080/swagger/index.html)

---

## 📡 Endpoints disponibles

| Método | Ruta                | Descripción                          |
|--------|---------------------|--------------------------------------|
| `POST` | `/users`            | Crear un nuevo usuario               |
| `GET`  | `/users`            | Obtener todos los usuarios           |
| `GET`  | `/users/{id}`       | Obtener usuario por ID               |
| `PATCH`| `/users/{id}`       | Actualizar un usuario   |
| `DELETE`| `/users/{id}`      | Eliminar un usuario por ID           |

---

## 📦 Estructura del proyecto

```
internal/
├── domain         # Entidades del negocio e interfaces
├── dto            # Objetos de transferencia (entrada/salida)
├── handler        # Controladores HTTP
├── infrastructure # Conexión a MongoDB
├── router         # Configuración de rutas con Gin
├── usecase        # Lógica de negocio (servicios)
```

---

## 📚 Reflexión y aprendizajes

Durante el desarrollo de esta API se documentaron los aprendizajes clave y oportunidades de mejora en un archivo adicional. Puedes consultarlo en el archivo:

📄 `LEARNING.md`
---
