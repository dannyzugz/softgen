import { useNavigate } from "react-router-dom";
import { useContext, useRef } from "react";
import { SettingsContext } from "../context/settingsContext";

export function NewProject() {
    const data = useRef();
    const navigate = useNavigate();
    const { projectData, setProjectData } = useContext(SettingsContext);

    const handleClick = () => {

        if (data.current.projectName.value && data.current.modulePath.value) {
            setProjectData({
                projectName: data.current.projectName.value,
                modulePath: data.current.modulePath.value,
            });
            navigate("/stepone");
        } else {
            console.log("Empty Data");
        }
    };

    return (
        <section>
            <h1>Code Generator</h1>
            {/* Temporal styles change later */}
            <form
                ref={data}
                style={{ display: "flex", flexDirection: "column" }}
            >
                <label>
                    Project Name
                    <input
                        type="text"
                        name="projectName"
                        defaultValue={projectData.projectName}
                        autoComplete="off"
                    />
                </label>

                <label>
                    Module Path
                    <input
                        type="text"
                        name="modulePath"
                        defaultValue={projectData.modulePath}
                        autoComplete="off"
                    />
                </label>

                <button type="button" onClick={handleClick}>
                    Next
                </button>
            </form>
        </section>
    );
}
