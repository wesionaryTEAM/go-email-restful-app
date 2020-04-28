### Create env

Create .env file in root of the project and add your credentials like that of .env.example

### API Endpoints
GET user-api/user → Retrieves all the user data
POST user-api/user → Add new user data
GET user-api/user/{id} → Retrieve the single user data
PUT user-api/user/{id} → Update the user data
DELETE user-api/user → Delete the user data
POST email/send-emal → Sends email (send email_to on body of the request like {"email_to": "email@email.com"})