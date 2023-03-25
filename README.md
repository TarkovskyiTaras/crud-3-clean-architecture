# Golang CRUD app:

## API:

- **GET** http://localhost:8080/person/1
- **GET** http://localhost:8080/person
- **POST** http://localhost:8080/person {"id" : 1, "first_name" : "Taras", "last_name" : "Tarkovskyi", "dob" : "1992-01-23T00:00:00Z", "home_address" : "Kyiv", "cellphone" : "0933115485"}
- **PUT** http://localhost:8080/person/1 {"id" : 1, "cellphone" : "000000000"}
- **DELETE** http://localhost:8080/person/1

## Docker:

- **$ docker container run --name crud-6-cont --publish 8080:8080 tarkovskyit/crud-docker-image-db-map**  - the app in the image is set to work with the go's map database (no postgres)