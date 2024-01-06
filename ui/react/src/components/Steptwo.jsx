import { Link } from "react-router-dom";

export function Steptwo() {
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
                <button>Generate Project</button>
            </section>

            <Link to={"/stepone"}>back</Link>
        </>
    );
}
