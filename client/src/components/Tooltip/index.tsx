import { Position } from "@/game/types";

export type TooltipProps = {
  position: Position;
};

function Tooltip({ position }: TooltipProps) {
  const { x, y } = position;
  return (
    <div className="absolute bottom-[-80px] left-[-70px] z-10 w-[70px] h-[80px] flex justify-center items-center border rounded-sm bg-gray-500 text-white ">
      <div>
        <div>move to:</div>
        <div>
          {x} , {y}
        </div>
      </div>
    </div>
  );
}

export default Tooltip;
