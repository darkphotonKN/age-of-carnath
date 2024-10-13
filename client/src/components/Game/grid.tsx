import GameBlock from "./block";

function GameGrid() {
  const COL_SIZE = 24;
  const ROW_SIZE = 16;

  const cols = new Array(COL_SIZE).fill(0);
  const rows = new Array(ROW_SIZE).fill(cols);

  return (
    <div className="mt-4 flex flex-col justify-center items-center">
      {rows.map((row, rowIndex) => {
        return (
          <div key={rowIndex} className="flex">
            {row.map((col: number, colIndex: number) => (
              <GameBlock key={rowIndex + " " + colIndex} />
            ))}
          </div>
        );
      })}
    </div>
  );
}

export default GameGrid;
