import { GameAction } from "@/constants/enums";
import { GamePayload, Player } from "@/game/types";
import { create } from "zustand";

type WebSocketState = {
  ws: WebSocket | null;
  isConnected: boolean;
  setupWebSocket: () => void;
  setWebSocket: (ws: WebSocket) => void;
  setConnectionStatus: (status: boolean) => void;
  sendMessage: <T>(payload: GamePayload<T>) => void;
  startMatchmaking: (player: Player) => void;
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
  isConnected: false,

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

    socket.onopen = (event) => {
      console.log("Connection Server Info on Open:", event);
      get().setWebSocket(socket);
    };

    socket.onerror = (error) => {
      console.log("WebSocket error:", error);
    };

    socket.onmessage = (event) => {
      console.log("Received from Server:", event.data);

      const serverMsgJson = JSON.parse(event.data);

      console.log("JSON from server:", serverMsgJson);

      if (typeof event.data === "string") {
        try {
          const message = JSON.parse(event.data);
          console.log("message received:", message);
        } catch (error) {
          console.log("Failed to parse incoming message:", error);
        }
      }
    };

    socket.onclose = () => {
      console.log("Disconnected from WebSocket server.");
    };
  },

  /**
   * Updates the WebSocket connection status.
   * Can be used to track the connection state across the app.
   **/
  setConnectionStatus: (status) => set({ isConnected: status }),

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

  // --- Helper Methods --

  /**
   * Helper method to start a match by sending a "join_match" action.
   * Wraps the logic of creating the message and calling sendMessage.
   **/
  startMatchmaking: (player: Player) => {
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
    console.log("cleaning up");
    // close connetion
    get().ws?.close();

    // reset ws instance
    set({ ws: null });
  },
}));
