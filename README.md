# Simple API in Golang  

## Dependencies
The following dependencies have been used:
```
github.com/spf13/viper
github.com/gofiber/fiber/v2
gorm.io/gorm
gorm.io/driver/postgres
```

## Application starting
Docker must be installed on your PC. Clone this repository and run:
```
docker-compose up
```

## Endpoints
The following endpoints have been created:

| Method | Route         | Body                                           |
| ------ | ------------- | ---------------------------------------------- |
| GET    | /users     |                                                |
| GET    | /users/:id |                                                |
| POST   | /users     | `{"name": "John", "phone": "123456789", "email": "email@email.com" }`|
| DELETE | /users/:id |                                                |
| PUT    | /users/:id | `{"name": "John", "phone": "123456789", "email": "email@email.com" }`|