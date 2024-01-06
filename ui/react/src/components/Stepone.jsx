import { Link } from "react-router-dom";

export function Stepone() {
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
            <span style={{ display: "flex", gap: 15 }}>
                <Link to={"/"}>back</Link>
                <Link to={"/steptwo"}>next</Link>
            </span>
        </section>
    );
}
