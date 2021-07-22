## Introduction
This gist is for a sample httpServer api app in GO, 
that has two endpoints to save, emit data for Contact entity. 
Contact is an entity with following fields base on initial requirement document.

full_name string
email     string 
phone_numbers string[] 

## Solution
this app will have following REST endpoints
### /contacts       GET to retrieve all contacts
### /contacts/id    GET specific contact

/contacts       POST/PUT it will take JSON formatted request body to save.
example 1 POST
```JSON
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
```
Note: contact_id is generated on server side, and used for updating the existing contacts.


### Assumptions, Validations
* Validations ( base test cases too for input)

 first_name can't be null and max chars allowed is 30
 last_name max chars allowed is 30
 email can't be null and has to be unique and in valid format
** phone_numbers can be multiple (optional)
** if one or phone number is present, check for valid E164 format using regex
** Only australian phone numbers allowed.

* test cases
In addition to the above validations, following edge cases can be added too.





* Suggestion 1 
full_name can be split into first name and last name, as this makes it easier to index, filter & search.
so the input data looks like
```json
{
    "first_name":"john",
    "last_name":"doe",
    "email":"john.doe@example.com",
    "phone_numbers":[
        "+61XXXYYYZZZ",
        "+61YYYYYYYYY",
    ]
}
```

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


* Dependencies
*** github.com/gorilla/mux v1.8.0
***	github.com/joho/godotenv v1.3.0
***	github.com/lib/pq v1.10.2

