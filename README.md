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
