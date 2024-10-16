import { ContentType, ContentTypeEnum } from "@/constants/enums";
import { Item, Player, Position } from "@/game/types";
import Image from "next/image";

type GameBlockProps = {
  onMouseEnter?: () => void;
  contentType: ContentTypeEnum;
  coords: Position;
  highlight?: boolean;
  content?: Player | Item;
};

function GameBlock({
  onMouseEnter,
  contentType,
  coords,
  highlight,
  content,
}: GameBlockProps) {
  function renderBlockContent() {
    const { x, y } = coords;

    switch (contentType) {
      // --- Render Player Character ---
      case ContentType.PLAYER:
        return (
          <Image
            width="50"
            height="50"
            src="/images/characters/medieval/adventurer_03/adventurer_03_1.png"
            alt="adventurer"
          />
        );

      // --- Render Item  ---
      case ContentType.ITEM:
        break;

      // --- Render Empty Grid ---
      case ContentType.EMPTY:
        return;
      default:
        return;
    }
  }
  return (
    <div
      // TODO: Remove after testing
      className={`border border-customBorderGray w-[53px] h-[53px] ${highlight ? "bg-customBorderGray" : ""}`}
      // className={`w-[35px] h-[35px] ${highlight ? "bg-customBorderGray" : ""}`}
      onMouseEnter={onMouseEnter}
    >
      {renderBlockContent()}
    </div>
  );
}

export default GameBlock;
