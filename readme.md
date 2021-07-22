   
   # Introduction
    This gist is for a sample httpServer api app written in GO, 
    that has two endpoints to save, emit data for Contact entity. 
    Contact is an entity with following fields base on initial requirement document.

    full_name string
    email     string 
    phone_numbers string[] 

    
   ## Solution
    This app will have following REST endpoints
   ### /contacts       GET to retrieve all contacts
   ### /contacts/{id}    GET specific contact

    /contacts       POST/PUT it will take JSON formatted request body to save.
    example 1 POST
   
    {
        "first_name":"john",
        "last_name":"doe",
        "email":"john.doe@example.com",
        "phone_numbers":[
            "+61XXXYYYZZZ",
            "+61YYYYYYYYY",
        ]
    }
    example 2 PUT
    {
        "contact_id":1,
        "first_name":"john",
        "last_name":"doe",
        "email":"john.doe@example.com",
        "phone_numbers":[
            "+61XXXYYYZZZ",
            "+61YYYYYYYYY",
        ]
    }
   
    Note: contact_id is generated on server side, and used for updating the existing contacts.


### Assumptions
* Validations ( base test cases too for input)
    * first_name can't be null and max chars allowed is 30
    * last_name max chars allowed is 30
    * email can't be null and has to be unique and in valid format
    * phone_numbers can be multiple (optional)
    * if one or more phone number are present, check for valid E164 format using regex
    * Only australian phone numbers allowed.

* Test cases ( In addition to the above validations, following edge cases can be added too.)
    * check all endpoints without schema and data
    * check all endpoints with shema but no data
    * post a larger payload (thousnnds of contacts) 
    * Check pagination on /contacts endpoint
    * check for sql injection, however parameterized procedures will block them
    * CORS check, if required.
  
* TODO / Suggestions  
    * ORM with repository pattern to handle db operations.
    * Logging fo requests/responses to capture errors, 
    * Stored procedure to miniming roundtrips, leverage query compilation & make db calls efficient 
    * if this API is only mean from fewer clients, it can have CORS headers match too
    
### To Run 
    1 Create a account on elephantSQL for or use or local postgres db.
    2 Modify db connection info in .env file
    3 Create schema and put some seed data in your db from schema.sql
    4 Run the commands below

    either use
    \>go run main.go app.go model.go
    or 
    \>go build -o app.exe
    \>app.exe


###### Dependencies
* github.com/gorilla/mux v1.8.0
* github.com/joho/godotenv v1.3.0
* github.com/lib/pq v1.10.2

###### App Structure
Main.go is the entry point of this api, and it hosts app.go, which handles the req/res followed by model.go, which interacts with db server.
main.go -> app.go -> models.go -> db 

It also a test file called main_test.go, which can be run like this.
\> go test -v



