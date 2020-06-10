# ID-Service
Manages authentication and authorization. For authentication this application uses a very 
simple approach. So it is the users responsibility not to insert any invalid user in the 
database.

# Run the app
```
./run.sh
```
## If connection refused error occur on id_service container then
> This error occur because mysql need some time to initialize but go request to connect 
> before mysql connection stabilize 
```
docker-compose down
# then
docker-compose up -d mysql
docker-compose up -d id_service
``` 

## How it works?
For a user get authenticated the system must know the user first. There are two types of 
users: *admin* and regular *user*. An admin cannot be created by a user sign-up process.
Admin account is created on company sign-up process. There can be only one admin per company.

*Signing up a company is restricted to Realtime Location Service authority.* 

During a company sign-up process an admin account is created with an AppKey to use with the
application for all kinds of interaction. For example signing up users and getting user 
locations and metadata, etc.

# ENV file
Rename the .env.example file to .env

## API endpoints

1. **POST** /api/v1/company/signup
   
   **Description:** Endpoint to register a company. This is a restricted endpoint. This will not be provided to any personnel
   outside of the Realtime Location Service. All request must have a secret header.
   
   **Payload** 
```        
        {
            "name": "example",
            "domain": "example.com"
        }
``` 
2. **POST** /api/v1/user/signup
   
   **Description:** Endpoint to register users for the companies.
   
   **Payload** AppKey is a secret which will be provided while registering the company.
   UserId is a unique id for the company to identify their users. Any sign-up request 
   having the already existing userId against an AppKey will not create any new user.
        
``` 
        {
            "appKey": "9Of0u659BiyVAmJ/58yzrw==",
            "userId": "test2"
        }
```
**Response:**
```
        {
            "id": 13,
            "company_id": 7,
            "user_id": "1",
            "role": "user",
            "created_at": "2020-03-19 06:08:56",
            "updated_at": "2020-03-19 06:08:56"
        }
```

3. **GET** /api/v1/user/resolve
   
   **Description:** Endpoint to retrieve all users for a company/AppKey.
   
   **Header:** `AppKey=9Of0u659BiyVAmJ%2F58yzrw%3D%3D`
        
   **ResponseBody:**
```   
       {
            "company_name": "Rezwanul",
            "company_id": 7,
            "domain": "rezwanul.com",
            "user_id": "861973be-0509-49c3-ad5b-e1c62153926f",
            "role": "admin",
            "subordinates": [
                {
                    "id": 13,
                    "company_id": 7,
                    "user_id": "1",
                    "role": "user",
                    "created_at": "2020-03-19 06:08:56",
                    "updated_at": "2020-03-19 06:08:56"
                }
            ]
        }
```

## Release Notes:
#### 1.0.2: 
> #### ID-Service converted to Golang
