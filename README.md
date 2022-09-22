# GO-REST-API
A REST api based on CURD operation written in Go and MongoDB is used as Database.

## How to strat the server

Copy the source code 

Add a .env file like sample.env and provide your mongodb username and password.

run below command at the root folder of the source code.

```
go run main.go
```
It starts at 6000 port. Check port availabilty before run.


Feel free to tweak and build executable. 

---

## File Structure


![Alt text](projectFileStructure.png?raw=true "Title")

---

## Backend flow 


![Alt text](backendFlow.png?raw=true "Title")

---

# API Documentation

## Create a user
Call a POST request to **/user** with data as json in body like below.

```
  {
        
        "name": "test1",
        "dob": "220922",
        "address": "India",
        "description": "test descrption"
        
    }
```
__Note:__  id and createdAt will be created via server.

_Server response_

```
{
    "status": 201,
    "message": "success",
    "data": {
        "data": {
            "InsertedID": "632c8424e108bfa308bf2327"
        }
    }
}
```
__Note:__ InsertedID is not user id. To check user id run Get all Users request(see below). 

---
## Get a user
Call a GET method to **/user/{id}** 


_Server response_

```
{
    "status": 200,
    "message": "success",
    "data": {
        "data": {
            "id": "632c8424e108bfa308bf2326",
            "name": "test1",
            "dob": "220922",
            "address": "India",
            "description": "test descrption",
            "createdAt": "time.Date(2022, time.September, 22, 15, 49, 56, 97440329, time.UTC)"
        }
    }
}
```

## Update a user 
Call an UPDATE request to **/user/{id}** with user data as json in body like below.

```
  {
        
        "name": "testupdate",
        "dob": "020689",
        "address": "India",
        "description": "test descrption"
        
    }
```

__Note:__ id and createdAt will not change. You must provide remaining fields as partial update is not granted.


_Server response_

```
{
    "status": 200,
    "message": "success",
    "data": {
        "data": {
            "id": "632c8424e108bfa308bf2326",
            "name": "testupdate",
            "dob": "020689",
            "address": "India",
            "description": "test descrption",
            "createdAt": "time.Date(2022, time.September, 22, 15, 49, 56, 97440329, time.UTC)"
        }
    }
}
```

## Delete a user
Call a DELETE request to **/user/{id}**

_Server response_

```
{
    "status": 200,
    "message": "success",
    "data": {
        "data": "User successfully deleted!"
    }
}
```
---
## Get all user list

Call a GET request to **/users**. (it is "users" not "user" like before )

_Server response_

```
{
    "status": 200,
    "message": "success",
    "data": {
        "data": [
            {
                "id": "632c8917e108bfa308bf2328",
                "name": "test51",
                "dob": "220922",
                "address": "India",
                "description": "test descrption",
                "createdAt": "time.Date(2022, time.September, 22, 16, 11, 3, 896277860, time.UTC)"
            },
            {
                "id": "632c891ee108bfa308bf232a",
                "name": "test52",
                "dob": "220922",
                "address": "India",
                "description": "test descrption",
                "createdAt": "time.Date(2022, time.September, 22, 16, 11, 10, 198696356, time.UTC)"
            },
            {
                "id": "632c8924e108bfa308bf232c",
                "name": "test53",
                "dob": "220922",
                "address": "India",
                "description": "test descrption",
                "createdAt": "time.Date(2022, time.September, 22, 16, 11, 16, 446019938, time.UTC)"
            },
            {
                "id": "632c892ae108bfa308bf232e",
                "name": "test54",
                "dob": "220922",
                "address": "India",
                "description": "test descrption",
                "createdAt": "time.Date(2022, time.September, 22, 16, 11, 22, 104376518, time.UTC)"
            }
        ]
    }
}
```
