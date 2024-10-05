function GameBlock() {
	const WIDTH = 100;
	const HEIGHT = 100;

	return (
		<div
			className={`border border-customBorderGray w-[${WIDTH}px] h-[${HEIGHT}px]`}
		></div>
	);
}

export default GameBlock;
