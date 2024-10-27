import { useWebsocketStore } from "@/stores/websocketStore";
import { Button } from "../Button";
import { useRouter } from "next/navigation";

function GameOptions() {
  // connect to websocket state store
  const { closeConnection } = useWebsocketStore();

  const router = useRouter();

  function handleCloseConnection() {
    // only close conneiton if ws is even instantiated
    console.log(
      "[@Game page] Cleaning up connection due to leaving the page or dismount",
    );
    closeConnection();

    router.push("/");
  }

  return (
    <div className="flex justify-center mt-5 gap-3">
      <Button variant="default">Settings</Button>
      <Button variant="default" onClick={handleCloseConnection}>
        Exit Match
      </Button>
    </div>
  );
}

export default GameOptions;
