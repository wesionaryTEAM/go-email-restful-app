### Create env

Create .env file in root of the project and add your credentials like that of .env.example

### API Endpoints
GET user-api/user → Retrieves all the user data <br />
POST user-api/user → Add new user data <br />
GET user-api/user/{id} → Retrieve the single user data <br />
PUT user-api/user/{id} → Update the user data <br />
DELETE user-api/user → Delete the user data <br /><br />
POST email/send-emal → Sends email (send email_to on body of the request like {"email_to": "email@email.com"})<br />