# Backend-Service
## Capstone Bangkit C242-PS001

If you want to use local db, set `APP_ENV` to `development`

## Endpoints

### POST /api/register
```
curl --location 'https://backend-service-282390196070.asia-southeast1.run.app/api/register' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "John Doe",
    "email": "john@doe.com",
    "password": "johndoehandsome"
}'
```

### POST /api/login
```
curl --location 'https://backend-service-282390196070.asia-southeast1.run.app/api/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "john@doe.com",
    "password": "johndoehandsome"
}'
```

### GET /api/ping
```
curl --location 'https://backend-service-282390196070.asia-southeast1.run.app/api/ping' \
--header 'Authorization: Bearer xxx'
```