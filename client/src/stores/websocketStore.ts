import { GameAction } from "@/constants/enums";
import { deduceGameAction } from "@/game/gameLogic";
import { GamePayload, GameState, Player } from "@/game/types";
import { create } from "zustand";

type WebSocketState = {
  ws: WebSocket | null;
  gameState: GameState | null;
  isConnected: boolean;
  findingMatch: boolean;
  matchInitiated: boolean;
  setupWebSocket: () => void;
  setWebSocket: (ws: WebSocket) => void;
  setGameState: (gameState: GameState) => void;
  setConnectionStatus: (status: boolean) => void;
  setFindingMatch: (finding: boolean) => void;
  sendMessage: <T>(payload: GamePayload<T>) => void;
  startMatchmaking: (player: Player) => void;
  setMatchInitiated: (initiated: boolean) => void;
  closeConnection: () => void;
};

/**
 * WebSocket State Management.
 * This store keeps track of the WebSocket instance, its connection state,
 * and provides methods to manage the connection and send messages.
 **/
export const useWebsocketStore = create<WebSocketState>((set, get) => ({
  // -- State Variables --
  ws: null,
  gameState: null,
  isConnected: false,
  findingMatch: false,
  matchInitiated: false,

  // -- State Methods --
  /**
   * Stores the WebSocket instance in the Zustand state.
   * This allows access to the WebSocket instance across the application.
   **/
  setWebSocket: (ws) => set({ ws }),

  /**
   * Initializes the WebSocket instance
   **/
  setupWebSocket: () => {
    // prevent re-setup
    if (get().ws) return;

    const socket = new WebSocket(`ws://localhost:4111/ws`);
    console.log("@WS creating websocket instance:", socket);

    socket.onopen = (event) => {
      console.log("Connection Server Info on Open:", event);
      get().setWebSocket(socket);
      get().setConnectionStatus(true);
    };

    socket.onerror = (error) => {
      console.log("@WS WebSocket error:", error);
    };

    socket.onmessage = (event) => {
      const serverMsgJson = JSON.parse(event.data);

      console.log("@WS JSON from server:", serverMsgJson);

      deduceGameAction(serverMsgJson);
    };

    socket.onclose = () => {
      console.log("@WS Disconnected from WebSocket server.");
    };
  },

  /**
   * Sets up gameState in zustand store as the single source of truth of the state of any match.
   **/

  setGameState: (gameState: GameState) => set({ gameState }),

  /**
   * Updates the WebSocket connection status.
   * Can be used to track the connection state across the app.
   **/
  setConnectionStatus: (status) => set({ isConnected: status }),

  /**
   * Updates the MatchInitiated status.
   * Used to determine whether or not a game has started for the client.
   **/
  setMatchInitiated: (initiated) => set({ matchInitiated: initiated }),

  /**
   * Sends a message to the WebSocket server.
   * Verifies the connection is open before sending the message.
   *
   * @param {GamePayload<T>} payload - The object with the type of action being * performed (e.g., 'find_match') and the payload of that action.
   **/
  sendMessage: <T>(payload: GamePayload<T>) => {
    const { ws } = get();
    if (ws && ws.readyState === WebSocket.OPEN) {
      ws.send(JSON.stringify(payload));
    }
  },

  /**
   * Sets the current state of the matchmaking process.
   **/
  setFindingMatch: (finding) => set({ findingMatch: finding }),

  // --- Helper Methods --

  /**
   * Helper method to start a match by sending a "join_match" action.
   * Wraps the logic of creating the message and calling sendMessage.
   **/
  startMatchmaking: (player: Player) => {
    console.log("@MATCHMAKING starting match player:", player);
    // let the client know matchmaking has started
    get().findingMatch = true;

    // initiate match making on the game server
    const payload: GamePayload<Player> = {
      action: GameAction.FIND_MATCH,
      payload: player,
    };
    get().sendMessage(payload);
  },

  /**
   * Performs a clean up for the websocket connection and the current websocket instance.
   **/
  closeConnection: () => {
    console.log("@MATCHMAKING @WS cleaning up");
    // close connection
    get().ws?.close();

    // reset ws instance
    set({ ws: null });

    // close connection status
    get().setConnectionStatus(false);

    // stop match
    get().setMatchInitiated(false);
  },
}));
