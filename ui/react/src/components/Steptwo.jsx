import { useContext, useRef } from "react";
import { Link } from "react-router-dom";
import { SettingsContext } from "../context/settingsContext";

export function Steptwo() {
    const { customSettings, setCustomSettings } = useContext(SettingsContext);
    const settings = useRef();

    const Generate = () => {

        if (
            settings.current.databaseType.value &&
            settings.current.routerType.value
        ) {
            setCustomSettings({
                db: settings.current.databaseType.value,
                router: settings.current.routerType.value,
            });
        } else {
            console.log("empty");
        }
    };

    return (
        <>
            <section>
                <h2> STEP 2 Customize your application</h2>
                <strong>Choose a database:</strong>
                <form ref={settings}>
                    <div>
                        <label>
                            <input
                                type="radio"
                                name="databaseType"
                                value="none"
                                defaultChecked={
                                    customSettings.db === "none" ||
                                    customSettings.db === ""
                                }
                            />
                            None
                        </label>
                        <label>
                            <input
                                type="radio"
                                name="databaseType"
                                value="postgre"
                                defaultChecked={customSettings.db === "postgre"}
                            />
                            PostgreSQL
                        </label>
                        <label>
                            <input
                                type="radio"
                                name="databaseType"
                                value="mysql"
                                defaultChecked={customSettings.db === "mysql"}
                            />
                            MySQL
                        </label>
                    </div>

                    <strong>Pick your preferred router:</strong>

                    <div>
                        <label>
                            <input
                                type="radio"
                                name="routerType"
                                value="chi"
                                defaultChecked={
                                    customSettings.router === "chi" ||
                                    customSettings.router === ""
                                }
                            />
                            Chi
                        </label>
                        <label>
                            <input
                                type="radio"
                                name="routerType"
                                value="gorila"
                                defaultChecked={
                                    customSettings.router === "gorila"
                                }
                            />
                            Gorila Mux
                        </label>
                        <label>
                            <input
                                type="radio"
                                name="routerType"
                                value="httpRouter"
                                defaultChecked={
                                    customSettings.router === "httpRouter"
                                }
                            />
                            HttpRouter
                        </label>
                        <label>
                            <input
                                type="radio"
                                name="routerType"
                                value="gin"
                                defaultChecked={customSettings.router === "gin"}
                            />
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
                <button onClick={Generate}>Generate Project</button>
            </section>

            <Link to={"/stepone"}>back</Link>
        </>
    );
}
