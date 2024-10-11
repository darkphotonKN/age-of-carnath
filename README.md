# Age of Carnath 1v1 - Turn-Based Online Multiplayer RPG

### Description

Age of Carnath is a 1v1 is a turn-based multiplayer RPG game where two players engage in player-vs-player (PvP) combat.

The backend server is written in Go, using WebSockets to facilitate real-time communication between the players. The server is designed to manage player connections, game state, and communication in a 1v1 format, spinning up unique instances for two 1v1 players while concurrently keeping the entire connected player base informed of the latest updates.

The client-side interface is developed in typescript reactjs with the nextjs framework, and involves in drawing the game's battle arena, showing the information of the players, and the UI which control the flow of the game.

### Features

- Real-time communication: The server uses WebSockets to enable real-time player interactions. Even though the game is turn based, the tracking of events and actual updates in the game happens real-time.

- Game state management: The server tracks each player's state, allowing actions like attacks, defenses, and movements to be broadcast to the other player instantly.

- Game loop: Players take turns in combat, with actions communicated through WebSocket messages.

- 1v1 player management and matchmaking: The server supports 1v1 matches., with both players connecting via WebSocket and Matches between two players are handled and creates a unique instance between two players when a match is found, whereby their actions are handled and broadcast to each other.

### Game Server

The protocol chosen was Websockets over TCP for the ease of use. However, there were still some design decisions that went into the game server's system design.

#### Using Unbuffered Channels

Since Gorilla Websocket was the most reliable websocket package in Go at the time of writing this project, it was used to create the entire websocket server for players to find matches and for the game's live action handling and tracking.

There was a big design decision to use **unbuffered** channels as a result, since there are analysis that show that the package only allows _one concurrent client to write to the server at once_. This means one client could potentially lock the entire server with message spam if unbuffered channels weren't used. Overall this prevented a huge number of writes to the server at the same time.

A simple read / write loop for websockets therefore turned into a slighty more complex read / write via websocket and with a communication layer via an unbuffered channel. Thankfully with go's channels and goroutines being performant the read / write remained incredibly smooth once handled appropriately.
