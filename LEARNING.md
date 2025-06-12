# 🚀 Aprendizajes del Desarrollo de la API User-API

Durante el desarrollo de esta API basada en un CRUD de usuarios, he consolidado y ampliado mis conocimientos en:

---

## ✅ Conceptos y Tecnologías Aprendidas

### 🧱 Arquitectura por capas
- Separación clara de responsabilidades: `domain`, `dto`, `handler`, `usecase`, `infrastructure`, y `router`.
### 🧱 ¿Por qué usar arquitectura por capas?

Opté por una arquitectura por capas porque permite **separar responsabilidades** en distintos niveles del sistema. Esto aporta varios beneficios importantes:

- ✅ **Mantenibilidad:** El código es más fácil de entender, modificar y extender sin afectar otras capas.
- ✅ **Escalabilidad:** Es posible agregar nuevas funcionalidades o tecnologías (por ejemplo, cambiar la base de datos) sin reescribir toda la aplicación.
- ✅ **Claridad y organización:** La estructura del proyecto es más limpia, clara y facilita la incorporación de nuevos desarrolladores.

Esta separación favorece además el cumplimiento de los principios **SOLID**, mejorando la calidad general del software.


### 🔄 Principios SOLID
- **Responsabilidad Única (SRP):** cada capa tiene su propia responsabilidad clara.
- **Inversión de Dependencias (DIP):** uso de interfaces para desacoplar la lógica de negocio de la capa de infraestructura.

### 🧩 DTOs
- Uso de `UserDTO` y `UserPatchDTO` para definir qué datos se exponen y reciben por la API, desacoplando el modelo de dominio del transporte de datos.

### 🌐 MongoDB
- Integración de una base de datos no relacional.
- Conversión entre `ObjectID` y strings para la interoperabilidad entre la base de datos y la API.

### 🧪 Swagger (OpenAPI)
- Generación automática de documentación de endpoints con comentarios en los handlers usando `swaggo`.

### 🐳 Docker
- Empaquetado de la aplicación con Docker para facilitar su despliegue.
- Uso de variables de entorno con `.env` para separar configuración del código.

---

## 🛠️ ¿Qué mejoraría?

### 🔐 Seguridad
- Añadir validaciones más robustas y control de entrada de datos.
- Incorporar autenticación (JWT/API Keys).

### 📦 Versionado de API
- Implementar versionado como `/api/v1/users` para garantizar compatibilidad futura.

### 🧪 Tests automatizados
- Agregar pruebas unitarias para la lógica de negocio.
- Uso de mocks para aislar pruebas de la infraestructura (Mongo).

### 📊 Mejor manejo de errores y respuestas estándar

- Se pueden utilizar funciones reutilizables para retornar errores de forma consistente (`respondWithError()`).
- Los errores personalizados (`var ErrUserNotFound = errors.New("user not found")`) permiten mejorar la claridad y controlar los flujos.
- Usar tipos como `AppError` con `código + mensaje` para tener respuestas uniformes en la API.


---

## 🙌 Conclusión

Este proyecto me permitió aplicar buenas prácticas de diseño de software con Go, fortalecer la estructura de una API RESTful, y afianzar conceptos esenciales como el uso de interfaces, manejo de errores, desacoplamiento y empaquetado de aplicaciones modernas con herramientas como Swagger y Docker.
