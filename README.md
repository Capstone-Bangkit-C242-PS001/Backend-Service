# SkillPath

SkillPath is an application designed to recommend online courses based on the skills users want to learn. The app consolidates courses from multiple platforms, such as Coursera, Udemy, and MIT OpenCourseWare, into a single, user-friendly platform. The backend is built using Gin and integrates seamlessly with cloud services for scalability.

## Tech Stack
### Backend
- **Framework:** Gin (Go, version 1.23)
- **API Documentation:** Swagger
- **Dependencies:**
  - Docker
  - Golang-migrate

### Cloud
- **GCP Dependencies:**
  - Cloud Run for deploying backend services
  - Cloud Storage for storing data assets

## API Documentation
The API is documented using Swagger. To view the API documentation, follow these steps:

1. Run the application.
2. Navigate to the `/swagger` endpoint in your browser.

Example:
```
http://localhost:9000/swagger/index.html
```

## File Structure
The project follows a modular structure:

```
SkillPath/
|   .dockerignore
|   .env
|   .env.example
|   .gitignore
|   docker-compose.yml
|   Dockerfile
|   go.mod
|   go.sum
|   Makefile
|   prod.Dockerfile
|   README.md
|
+---cmd
|       main.go
|
+---config
|       config.go
|
+---controller
|   +---auth
|   |       auth.controller.go
|   |
|   +---course
|   |       course.controller.go
|   |
|   +---interest_mapping
|   |       interest_mapping.controller.go
|   |
|   +---prediction
|   |       prediction.controller.go
|   |
|   +---rating
|   |       rating.controller.go
|   |
|   +---user
|   |       user.controller.go
|   |
|   \---user_interest
|           user_interest.controller.go
|
+---credentials
|       credentials.json
|
+---database
|   |   db.go
|   |
|   \---migration
|           000001_create_table_user.down.sql
|           000001_create_table_user.up.sql
|           000002_create_table_course.down.sql
|           000002_create_table_course.up.sql
|           000003_create_table_user_course_rating.down.sql
|           000003_create_table_user_course_rating.up.sql
|
+---docs
|       docs.go
|       swagger.json
|       swagger.yaml
|
+---dto
|   +---auth
|   |       auth.dto.go
|   |
|   +---course
|   |       course.dto.go
|   |
|   +---external
|   |       ml_service.dto.go
|   |
|   +---interest_mapping
|   |       interest_mapping.dto.go
|   |
|   +---prediction
|   |       prediction.dto.go
|   |
|   +---rating
|   |       rating.dto.go
|   |
|   +---response
|   |       paginate.dto.go
|   |       param.dto.go
|   |       response.dto.go
|   |
|   +---user
|   |       user.dto.go
|   |
|   \---user_interest
|           user_interest.dto.go
|
+---errorhandler
|       handler.go
|       types.go
|
+---external
|       ml_service.go
|
+---middleware
|       middleware.go
|
+---model
|   +---course
|   |       course.go
|   |
|   +---interest_mapping
|   |       interest_mapping.go
|   |
|   +---rating
|   |       rating.go
|   |
|   +---user
|   |       user.go
|   |
|   \---user_interest
|           user_interest.go
|
+---repository
|   +---course
|   |       course.repository.go
|   |
|   +---interest_mapping
|   |       interest_mapping.repository.go
|   |
|   +---rating
|   |       rating.repository.go
|   |
|   +---user
|   |       user.repository.go
|   |
|   \---user_interest
|           user_interest.repository.go
|
+---routes
|       auth.routes.go
|       course.routes.go
|       interest_mapping.routes.go
|       prediction.routes.go
|       rating.routes.go
|       user.routes.go
|       user_interest.routes.go
|
+---service
|   +---auth
|   |       auth.service.go
|   |
|   +---course
|   |       course.service.go
|   |       
|   +---interest_mapping
|   |       interest_mapping.service.go
|   |
|   +---prediction
|   |       prediction.go
|   |
|   +---rating
|   |       rating.service.go
|   |
|   +---user
|   |       user.service.go
|   |
|   \---user_interest
|           user_interest.service.go
|
\---utils
        gcs.go
        hashed.go
        token.go
        upload.go
```

## Known Issues
- **Cold Start Latency:** The server may experience high latency on the first hit due to cold start.

## How to Run the Project
1. Clone the repository.
2. Set up the `.env` file with the required environment variables. Refer to the `.env.example` file for the format and required fields.
3. Run the backend service:
    ```bash
    make start-build
    ```
4. Stop the backend service:
    ```bash
    make stop
    ```

## How to Run the Database Migration
1. Create migration
    ```bash
    make create-migration name={name}
    ```
2. Create migration up
    ```bash
    make migration-up
    ```
3. Create migration down
    ```bash
    make migration-down
    ```
4. Create migration force
    ```bash
    make migration-force version={version}
    ```
