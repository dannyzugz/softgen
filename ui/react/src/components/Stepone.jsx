import { Link, useNavigate } from "react-router-dom";
import { useContext, useRef } from "react";
import { SettingsContext } from "../context/settingsContext";

export function Stepone() {
    const appTypeRef = useRef();
    const navigate = useNavigate();
    const { projectType, setProjectType } = useContext(SettingsContext);

    const handleSubmit = (event) => {
        event.preventDefault();

        setProjectType(appTypeRef.current.appType.value);

        // replace with a switch
        if (appTypeRef.current.appType.value === "cli") {
            navigate("/new");
        } else if (
            (appTypeRef.current.appType.value === "ssr") |
            (appTypeRef.current.appType.value === "api")
        ) {
            navigate("/steptwo");
        }
    };

    return (
        <section>
            <h2>STEP 1 Choose your application type</h2>

            <form onSubmit={handleSubmit} ref={appTypeRef}>
                <div>
                    <label>
                        <input
                            type="radio"
                            name="appType"
                            value="ssr"
                            defaultChecked={
                                (projectType === "ssr") | (projectType === "")
                            }
                        />
                        Traditional web application: Generate the scaffold for a
                        'traditional' web application which uses server-side
                        HTML rendering.
                    </label>
                </div>
                <div>
                    <label>
                        <input
                            type="radio"
                            name="appType"
                            value="api"
                            defaultChecked={projectType === "api"}
                        />
                        JSON API: Generate the scaffold for an API which
                        processes JSON requests and sends JSON responses.
                    </label>
                </div>
                <div>
                    <label>
                        <input
                            type="radio"
                            name="appType"
                            value="cli"
                            defaultChecked={projectType === "cli"}
                        />
                        CLI Project: Generate the scaffold for a CLI application
                    </label>
                </div>
                {/* This is wrong */}
                <Link to={"/"}>back</Link>
                <button type="submit">next</button>
            </form>
        </section>
    );
}
