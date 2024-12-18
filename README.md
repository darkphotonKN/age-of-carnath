# Age of Carnath 1v1 - Turn-Based Online Multiplayer RPG

### Description

Age of Carnath is a 1v1 is a turn-based multiplayer RPG game where two players engage in player-vs-player (PvP) combat.

The backend server is written in Go, using WebSockets to facilitate real-time communication between the players. The server is designed to manage player connections, game state, and communication in a 1v1 format, spinning up unique instances for two 1v1 players while concurrently keeping the entire connected player base informed of the latest updates. Go was chosen for performance and simplicity.

The client-side interface is developed in typescript reactjs with the nextjs framework, and focuses rendering the game's battle arena, showing the information of the players, and the UI which control the flow of the game. Despite react performance not being optimate for rendering many updates on a game-like grid these tools were chosen due to the familiarity of the framework.

### Features

- Real-time communication: The server uses WebSockets to handle real-time player interactions. Even though the game is turn based, the tracking of events and actual updates in the game happens real-time.

- Game state management: The server tracks each player's state, allowing actions like attacks, defenses, and movements to be broadcast to the other player instantly.

- Game loop: Players take turns in combat, with actions communicated through WebSocket messages, and the game loop tracks vital information to keep matches in sync.

- Matchmaking: Matchmaking alogithms to help players find open matches to play.

- 1v1 player management: Hanldes events of 1v1 matches, with both players connecting via WebSocket connections. The matches between two players are handled by creating a unique instance between two players when a match is found, whereby their actions are handled and broadcast to each other.

### Game Server

The protocol chosen was Websockets over TCP for the ease of use. However, there were still some design decisions that went into the game server's system design.

#### Using Unbuffered Channels

Since Gorilla Websocket was the most reliable websocket package in Go at the time of writing this project, it was used to create the entire websocket server for players to find matches and for the game's live action handling and tracking.

There was a big design decision to use **unbuffered** channels as a result, since there are analysis that show that the package only allows _one concurrent client to write to the server at once_. This means one client could potentially lock the entire server with message spam if unbuffered channels weren't used. Overall this prevented a huge number of writes to the server at the same time.

A simple read / write loop for websockets therefore turned into a slighty more complex read / write via websocket and with a communication layer via an unbuffered channel. Thankfully with go's channels and goroutines being performant the read / write remained incredibly smooth once handled appropriately.

Race conditions were the other challenge that made multiple connections a challenge, but more details on this in the game logic section.

### Client Game State & Logic

The frontend client was coded in reactjs, using a mixture of zustand for global statemanagement and react's built-in state hooks for game interactions.

#### GameGrid Component

The core component for rendering the game's visuals and controlling the syncing between game state of the client and server.

#### Game State Handling

The game's state is managed by both the game server providing infromation over the persistent websocket connection as well as react's local component state and zustand providing global state.

The local state inside the GameGrid component is used to render things for temporary visual purposes, such as highlighting the game view with actions or seeing temporary information.

The global state uses zustand over redux for reduced boilerplate and simplicity of code. There is a single useWebsocketStore to interact with the global state source. This source provides two things:

- A centralized access to the websocket instance connection for the application.
- A synced game-state that originated from the websocket server that multiple components require.

This global source is used for controlling game interactions and providing real-time updates for each match.
