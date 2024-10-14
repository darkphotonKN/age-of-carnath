type GameBlockProps = {
  onMouseEnter?: () => void;
};

function GameBlock(props: GameBlockProps) {
  return (
    <div className={`border border-customBorderGray w-[35px] h-[35px]`}></div>
  );
}

export default GameBlock;
