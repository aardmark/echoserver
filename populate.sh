#!/bin/sh
curl -X POST http://localhost:8080/users -H "Content-Type: application/json" -d '{"firstname" : "Fred", "lastname" : "Flintstone", "isAdmin" : true, "password" : "password", "email" : "fred@bedrock.gov" }'
curl -X POST http://localhost:8080/users -H "Content-Type: application/json" -d '{"firstname" : "Wilma", "lastname" : "Flintstone", "isAdmin" : false, "password" : "password", "email" : "wilma@bedrock.gov" }'
curl -X POST http://localhost:8080/users -H "Content-Type: application/json" -d '{"firstname" : "Barney", "lastname" : "Rubble", "isAdmin" : false, "password" : "password", "email" : "barney@bedrock.gov" }'
curl -X POST http://localhost:8080/users -H "Content-Type: application/json" -d '{"firstname" : "Betty", "lastname" : "Rubble", "isAdmin" : false, "password" : "password", "email" : "betty@bedrock.gov" }'
