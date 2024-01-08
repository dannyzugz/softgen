import { createContext, useState, useRef } from "react";

export const SettingsContext = createContext();

export function SettingsProvider({ children }) {

    const [projectData, setProjectData] = useState({ projectName: "", modulePath: "" });
    const [projectType, setProjectType] = useState('')

    // Customization Settings
    const [customSettings, setCustomSettings] = useState({db:'', router: ''}) 

    return (
        <SettingsContext.Provider
            value={{
                projectData: projectData,
                setProjectData: setProjectData,
                projectType: projectType,
                setProjectType: setProjectType,
                customSettings: customSettings,
                setCustomSettings: setCustomSettings
            }}
        >
            {children}
        </SettingsContext.Provider>
    );
}
