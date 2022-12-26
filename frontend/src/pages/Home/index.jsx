import SkinCard from "../../components/SkinCard";

function Home() {
	return (
		<div className="justify-center align-middle flex">
			<div className="max-w-screen-xl  bg-slate-50 m-10 rounded shadow-md flex flex-col gap-8">
				<div className="w-full h-24 bg-slate-400 rounded-t px-8"></div>
				<div className="grid grid-flow-col gap-8 px-8 pb-8">
					{(() => {
						var e = [];

						for (var i = 0; i < 10; i++) {
							e.push(
								<SkinCard
									description={`A skin (${i})`}
									key={i}
								/>
							);
						}

						return e;
					})()}
				</div>
			</div>
		</div>
	);
}

export default Home;
