//import SkinCard from "../../components/SkinCard";

function Home() {
    return (
        <div className="flex justify-center align-middle">
            <div className="app-pannel container m-10 flex max-w-screen-xl flex-col gap-8">
                <div className="h-24 w-full rounded-t bg-slate-400 px-8"></div>
                <div className="max grid grid-flow-row grid-cols-1 justify-items-center gap-8 px-8 pb-8 md:grid-cols-4 xl:grid-cols-5"></div>
            </div>
        </div>
    );
}

export default Home;
