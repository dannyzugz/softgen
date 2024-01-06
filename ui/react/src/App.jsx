import "./App.css";
import { Routes, Route, NavLink } from "react-router-dom";
import { New } from "./components/New";
import { Steptwo } from "./components/Steptwo";
import { Stepone } from "./components/Stepone";

function Footer() {
    
    const STEPS = ["/", "/Stepone", "/Steptwo"];

    return STEPS.map((step, index) => (
        <span key={index} style={{padding: '3px'}}>
            <NavLink
                style={({ isActive }) => (
                isActive ? {
                    fontWeight: 'bold',
                    color: "white",
                    textDecorationLine: "underline",
                    fontSize: '110%'
                }:{})}
                to={step}
            >
                {index + 1}
            </NavLink>
        </span>
    ));
}

function App() {
    return (
        <>
            <main>
                <Routes>
                    <Route path="" Component={New}></Route>
                    <Route path="stepone" Component={Stepone}></Route>
                    <Route path="steptwo" Component={Steptwo}></Route>
                </Routes>
            </main>

            <Footer />
        </>
    );
}

export default App;
