import GameBlock from "./block";

function GameGrid() {
	return (
		<div className="grid grid-cols-5">
			{new Array(10).fill(0).map((block: number) => (
				<GameBlock key={block} />
			))}
		</div>
	);
}

export default GameGrid;
