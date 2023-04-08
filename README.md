# Golang CRUD app:

## API:

### User:
- **GET** http://localhost:8080/user/1
- **GET** http://localhost:8080/user
- **POST** http://localhost:8080/user {"id":1,"first_name":"Jonathan","last_name":"Adams","dob":"1987-03-21T00:00:00Z","location":"USA","cellphone_number":"+16479250145","email":"Jonathan@gmail.com","password":"pw124567"}
  - curl -i -X POST -H "Content-Type: application/json" -d '{"id":1,"first_name":"Jonathan","last_name":"Adams","dob":"1987-03-21T00:00:00Z","location":"USA","cellphone_number":"+16479250145","email":"Jonathan@gmail.com","password":"pw124567"}' "127.0.0.1:8080/user"
- **PUT** http://localhost:8080/user {"id":1,"first_name":"UPD_Jonathan","last_name":"UPD_Adams","dob":"1987-03-21T00:00:00Z","location":"USA","cellphone_number":"+16479250145","email":"Jonathan@gmail.com","password":"pw124567"}
  - curl -i -X PUT -H "Content-Type: application/json" -d '{"id":1,"first_name":"UPD_Jonathan","last_name":"UPD_Adams","dob":"1987-03-21T00:00:00Z","location":"USA","cellphone_number":"+16479250145","email":"Jonathan@gmail.com","password":"pw124567"}' "127.0.0.1:8080/user"
- **DELETE** http://localhost:8080/user/1
  - curl -i -X DELETE "127.0.0.1:8080/user/1"

### Book:
- **GET** http://localhost:8080/book/1
- **GET** http://localhost:8080/book
- **POST** http://localhost:8080/book {"id" : 1,"tittle" : "Handbook of Steel Construction","author" : "CISC ICCA","pages" : 290,"quantity" : 10}
  - curl -i -X POST -H "Content-Type: application/json" -d '{"id" : 1,"tittle" : "Handbook of Steel Construction","author" : "CISC ICCA","pages" : 290,"quantity" : 10}' "127.0.0.1:8080/book"
- **PUT** http://localhost:8080/book {"id" : 1,"tittle" : "UPD_Handbook of Steel Construction","author" : "UPD_CISC ICCA","pages" : 290,"quantity" : 10}
  - curl -i -X PUT -H "Content-Type: application/json" -d '{"id" : 1,"tittle" : "UPD_Handbook of Steel Construction","author" : "UPD_CISC ICCA","pages" : 290,"quantity" : 10}' "127.0.0.1:8080/book"
- **DELETE** http://localhost:8080/book/1
  - curl -i -X DELETE "127.0.0.1:8080/book/1"

### Loan:
- **GET** http://localhost:8080/loan/borrow/1/1
- **GET** http://localhost:8080/loan/return/1/1