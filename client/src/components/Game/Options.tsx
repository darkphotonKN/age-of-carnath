import { useWebsocketStore } from "@/stores/websocketStore";
import { Button } from "../Button";
import { useRouter } from "next/navigation";

function GameOptions() {
  // connect to websocket state store
  const { ws, closeConnection } = useWebsocketStore();

  const router = useRouter();

  function handleCloseConnection() {
    // only close conneiton if ws is even instantiated
    if (ws) {
      console.log(
        "[@Game page] Cleaning up connection due to leaving the page or dismount",
      );
      closeConnection();

      router.push("/");
    }
  }

  return (
    <div>
      <Button variant="default" onClick={handleCloseConnection}>
        Exit Match
      </Button>
    </div>
  );
}

export default GameOptions;
