type GameBlockProps = {
  onMouseEnter?: () => void;
  highlight?: boolean;
};

function GameBlock({ highlight, onMouseEnter }: GameBlockProps) {
  return (
    <div
      // TODO: Remove after testing
      className={`border border-customBorderGray w-[35px] h-[35px] ${highlight ? "bg-customBorderGray" : ""}`}
      // className={`w-[35px] h-[35px] ${highlight ? "bg-customBorderGray" : ""}`}
      onMouseEnter={onMouseEnter}
    ></div>
  );
}

export default GameBlock;
