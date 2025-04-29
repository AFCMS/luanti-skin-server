import { BrowserRouter, Route, Routes } from "react-router-dom";

// import Header from "./components/Header";
import Home from "./pages/Home";
import Login from "./pages/Login";
import Register from "./pages/Register";
import About from "./pages/About";
import NotFound from "./pages/NotFound";
// import { AppContextProvider } from "./services/AppContext.tsx";

function App() {
    return (
        <div className="min-h-screen bg-blue-200">
            {/* <AppContextProvider> */}
                <BrowserRouter>
                    {/* <Header /> */}
                    <Routes>
                        <Route path="/" element={<Home />} />
                        <Route path="/login" element={<Login />} />
                        <Route path="/register" element={<Register />} />
                        <Route path="/about" element={<About />} />
                        <Route path="/*" element={<NotFound />} />
                    </Routes>
                </BrowserRouter>
            {/* </AppContextProvider> */}
        </div>
    );
}

export default App;
