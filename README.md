# SSO-using-JWT-authentication
This is a repository for the project 'Single Sign On using JWT authentication'.

Github link to the Project: https://github.com/bobysingh494/SSO-using-JWT-authentication.

The complete directory of the project is as follows:

SingleSignOn
  auth
    token.go
  config
    app.conf
    configDB.go
  controllers
    userdetails.go
    userlogin.go
  models
    Response.go
    user.go
  responses
    response.go
  routers
    router.go
  go.mod
  go.sum
  main.go

How to run the Project:

1. Open terminal & change directory, to the project directory. (In my case it is “C:\Users\Boby Singh\go\bin\SingleSignOn”)

2. Now, run “go run main.go” from the terminal. 

3. Now if it runs successfully without any errors, a popup screen would appear asking the permission to allow to build the current main.exe file.

4. Now click the allow button and open the Postman client app.

5. Now there are two APIs in the project. 

    a. The first one is a sign-in API which will take the Email_id and Password of the user as the login credential and if credentials match from the database, it will log in the user by generating a JWT for that user.
    Here we have to do a “GET” request from the postman client with request URL http://localhost:8000/signin.
    Also in the body of the request, we have to give the email_id and password of the user as follows:

    {
        "Email_id": "kewalshahcric@iitr.ac.in",
        "Password": "kshah"
    }

    If login credentials are wrong then the response will look like this:

    {
        "HTTP status code": 422,
        "message": "Invalid Credentials"
    }

    And if credentials are right then the response will look like this:

    {
        "HTTP status code": 200,
        "Message": "Sign-in Successful",
        "JWT": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2MjgxNTk1MDgsInVzZXJfaWQiOjl9.oUBqgOHVeHs-m8TEKurbx5lf5vCRS6tOjsaZjZIvFxw"
    }

    b. The second API is an API that will take JWT(JSON Web Token) as input and give the details of the corresponding user as an output if the Token is valid and corresponding to any particular user from the given database.
    Here also we have to do a “GET” request from the postman client with the request URL http://localhost:8000/getmydata. Here we need to paste the Token as input by going into the authorization field of the request and pasting the token in the bearer token field. If the Token is invalid or the request was not correctly made it will respond like this:

    {
        "HTTP status code": 422,
        "message": "Invalid Token"
    }
    
    And if everything went alright then the response would like this:

    {
        "HTTP status code": 200,
        "Message": "Valid JWT, Information is successfully extracted",
        "User Information": {
            "user_id": 9,
            "email_id": "kewalshahcric@iitr.ac.in",
            "first_name": "Kewal",
            "last_name": "Shah",
            "contact_no": "7552863695"
        }
    }

