type GameBlockProps = {
  onMouseEnter?: () => void;
  highlight?: boolean;
};

function GameBlock({ highlight, onMouseEnter }: GameBlockProps) {
  console.log("highlight:", highlight);
  return (
    <div
      className={`border border-customBorderGray w-[35px] h-[35px] ${highlight ? "bg-pink-500" : ""}`}
      onMouseEnter={onMouseEnter}
    ></div>
  );
}

export default GameBlock;
