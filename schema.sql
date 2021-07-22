CREATE TABLE contacts
(
    id SERIAL PRIMARY KEY,
    first_name varchar(30) NOT NULL,
    last_name varchar(30),
    email varchar(30) NOT NULL,
    UNIQUE(email) 
)

CREATE TABLE contact_phones
(
    id SERIAL,
    contact_id int NOT NULL REFERENCES contacts,
    phone_number varchar(20) NOT NULL 
    /*
    TODO,
    Good to have fields like country_code, and phone_type (for eg. Home, Office) for better tagging.

    */
)
 

/*
//TODO
//Better way to handle the multi table inserts is a stored procedure.
//Plan was to save in primary table first (contacts)
// then take the newly generated contact.id and remove its child data (contact_phones)
// so that new phones on contact_phones have no cases for duplicates.
// not complete yet. need to fix the delete and insert part to child table.

create or replace procedure saveContact(
    first_name varchar(30),
    last_name varchar(30),   
    email varchar(30),
    phones varchar(500),
    c_id INOUT int
)
language plpgsql    
as $$
begin
     with c as 
     (
        insert into contacts (first_name, last_name,email)
        values( first_name ,last_name,email)
        RETURNING id
    )
    select id into contact_var from c; 
    delete from contact_phones cp join contact_var cv on  where cp.contact_id = cv.id;
 
    with phones as
    (
        select id, regexp_split_to_table(phones , E',') as ph from contact_var 
    )   
    insert into contact_phones(contact_id,phone_number)
    select id,ph from phones; 
    commit;
end;$$

----------------------
sample insert queries for data generation


insert into contacts (first_name,last_name,email) values ('John','Doe','john.doe@example.com');
insert into contacts (first_name,last_name,email) values('Karl','M','k.m@example.com');
insert into contacts (first_name,last_name,email) values ('Jessie','D','j.doe@example.com');

insert into contact_phones (phone_number,contact_id) values ('+61XXXYYYZZZ',1);


*/