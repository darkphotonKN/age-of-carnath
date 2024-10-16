import { ContentType, ContentTypeEnum } from "@/constants/enums";
import { Item, Player, Position } from "@/game/types";
import Image from "next/image";
import Tooltip, { TooltipProps } from "../Tooltip";

type GameBlockProps = {
  onMouseEnter?: () => void;
  contentType: ContentTypeEnum;
  coords: Position;
  highlight?: boolean;
  content?: Player | Item;
  tooltipProps?: TooltipProps;
};

function GameBlock({
  onMouseEnter,
  contentType,
  coords,
  highlight,
  content,
  tooltipProps,
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
      className={`relative border border-customBorderGray w-[53px] h-[53px] ${highlight ? "bg-customBorderGray" : ""}`}
      // className={`w-[35px] h-[35px] ${highlight ? "bg-customBorderGray" : ""}`}
      onMouseEnter={onMouseEnter}
    >
      {tooltipProps &&
        tooltipProps.position.x === coords.x &&
        tooltipProps.position.x === coords.x && <Tooltip {...tooltipProps} />}
      {renderBlockContent()}
    </div>
  );
}

export default GameBlock;
