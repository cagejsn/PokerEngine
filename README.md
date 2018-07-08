
# PokerEngine

This is a Go App that listens on port 8080 and exposes three endpoints:i

/register <- POST with new user in JSON

/login <- POST with user credentials in JSON

/play <- accepts a websocket request and opens a websocket connection. It will validate that there is a valid UserSession that was created with a successful login using a cookie

the logic of the poker game is inside GameController

##Important Notes

the cookies which enable the usersessions to persist across network calls will not work unless you hit the loopback ip 127.0.0.1 rather than localhost


## Getting Started


In order to use this you'll probably want to go grab the stack.yml from here https://github.com/cagejsn/PokerStack

deploying this stack will make the database that this app depends on and it will also make the frontend on port 4200 so you can interact with it more easily. Deploying the stack will make an app with this same source on port 8090.

for development you can either:

a) deploy the stack , clone this repo and run it on 8090, and then clone the frontend repo and run it using the angular CLI on port 4201

or 

b) run the db container on it's own, clone this repo and run it on 8090, clone the front end repo and run it using the angular CLI on 4201

either way the trick is to get the Database up and running which I find annoying so I just use the use the docker stack with the db service configured

