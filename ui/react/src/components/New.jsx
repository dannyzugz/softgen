import { Link } from "react-router-dom";

export function New() {
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
