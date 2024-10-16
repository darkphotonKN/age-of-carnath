import { Position } from "@/game/types";

export type TooltipProps = {
  position: Position;
};

function Tooltip({ position }: TooltipProps) {
  const { x, y } = position;
  return (
    <div className="absolute top-0 left-0 border border-gray-700 w-[50px] h-[50px]">
      <div>x: {x}</div>

      <div>y: {y}</div>
    </div>
  );
}

export default Tooltip;
