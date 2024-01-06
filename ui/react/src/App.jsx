import "./App.css";
import { Routes, Route, Link } from "react-router-dom";

function Stepone() {
    return (
        <section>
            <h2>STEP 1 Choose your application type</h2>
            <form action="">
                <div>
                    <label>
                        <input type="radio" name="app-type" />
                        Traditional web application: Generate the scaffold for a
                        'traditional' web application which uses server-side
                        HTML rendering.
                    </label>
                </div>
                <div>
                    <label>
                        <input type="radio" name="app-type" />
                        JSON API: Generate the scaffold for an API which
                        processes JSON requests and sends JSON responses.
                    </label>
                </div>
                <div>
                    <label>
                        <input type="radio" name="app-type" />
                        CLI Project: Generate the scaffold for a CLI application
                    </label>
                </div>
            </form>

            {/* Temporal style */}
            <span style={{display:"flex", gap:15}}>
                <Link to={"/"}>back</Link> 
                <Link to={"/steptwo"}>next</Link>
            </span>
        </section>
    );
}

function Steptwo() {
    return (
        <>
            <section>
                <h2> STEP 2 Customize your application</h2>
                <strong>Choose a database:</strong>
                <form action="">
                    <div>
                        <label>
                            <input type="radio" name="database-type" />
                            None
                        </label>
                        <label>
                            <input type="radio" name="database-type" />
                            PostgreSQL
                        </label>
                        <label>
                            <input type="radio" name="database-type" />
                            MySQL
                        </label>
                    </div>

                    <strong>Pick your preferred router:</strong>

                    <div>
                        <label>
                            <input type="radio" name="router-type" />
                            Chi
                        </label>
                        <label>
                            <input type="radio" name="router-type" />
                            Gorila Mux
                        </label>
                        <label>
                            <input type="radio" name="router-type" />
                            HttpRouter
                        </label>
                        <label>
                            <input type="radio" name="router-type" />
                            Gin
                        </label>
                    </div>
                </form>
            </section>

            <section>
                <h2>Generate</h2>
                <div>
                    Click the button below to generate a ZIP file containing
                    your project code
                </div>
                <input type="button" value={"Generate project!"} />
            </section>

            <Link to={"/stepone"}>back</Link>
        </>
    );
}

function New() {
    return (
        <section>
            <h1>Code Generator</h1>
            {/* Temporal styles change later */}
            <div style={{ display: "flex", flexDirection: "column" }}>
                <label>
                    Project Name
                    <input type="text" name="project-name" />
                </label>

                <label>
                    Module Path
                    <input type="text" name="project-name" />
                </label>
            </div>

            <Link to={"/stepone"}>next</Link>
        </section>
    );
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

            <footer>1 2 3</footer>
        </>
    );
}

export default App;
