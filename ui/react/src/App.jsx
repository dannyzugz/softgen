import "./App.css";
import { Routes, Route } from "react-router-dom";
import { NewProject } from "./components/NewProject";
import { Steptwo } from "./components/Steptwo";
import { Stepone } from "./components/Stepone";
import { Footer } from "./components/Footer";

function App() {
    return (
        <>
            <main>
                <Routes>
                    <Route path="" Component={NewProject}></Route>
                    <Route path="stepone" Component={Stepone}></Route>
                    <Route path="steptwo" Component={Steptwo}></Route>
                </Routes>
            </main>

            <Footer />
        </>
    );
}

export default App;
