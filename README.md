# minesweeper-API

## The Game
Develop the classic game of [Minesweeper](https://en.wikipedia.org/wiki/Minesweeper_(video_game))

## Features
- Start a new game with variable number of rows and columns
- Tracking the time of the game
- Tracking the game status (in progress, won or lost)
- Persist games into a database
- Use uuid for identifier for games instead of ids
- Fire clicks and flag events into the board
- Fire recursevely the click reveal event when clicking on a empty cell
- Get Board Status
- Detection of game ending and winning

## Endpoints
- Start: For starting a new game
- Click: For firing a click event on the board
- Flag: For firing a flag event on the board
- Status: Just for knowing that the service is available
- Game Status: To get the game status
- Draw: For testing. This endpoint returns a raw text with the board drawn

## Important consideration
- For the time being I didnt implement user authentication or session. I just didnt had that time.
- Overall the API covers all of the cases for the game.
- I also left to be done that I believe that some fields in the responses could be hidden so in a 
- client-server connection the developer-players dont just look at the data and won the game.
- I decided to use Heroku as a Host and mysql with ClearDB in heroku because I felt confortable with it
and dont have time to invest. Deployent in Heroku is really fast.
- There is a JS client in `client` folder you can take a look to see how to implement a client
to connect to this api
- I decided to use GIN as I think its a very good and straight forward library to create microservices in Golang
- I decided to use GORM as the ORM because it is very easy to automigrate and manage tables as structs
- There is a test.http file for testing endpoints directly in the IDE
- I didnt created tests because of the time I had to do this.