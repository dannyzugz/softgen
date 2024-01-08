import { NavLink } from "react-router-dom";

export function Footer() {
    const STEPS = ["/", "/Stepone", "/Steptwo"];

    return STEPS.map((step, index) => (
        <span key={index} style={{ padding: "3px" }}>
            <NavLink
                style={({ isActive }) => isActive
                    ? {
                        fontWeight: "bold",
                        color: "white",
                        textDecorationLine: "underline",
                        fontSize: "110%",
                    }
                    : null}
                to={step}
            >
                {index + 1}
            </NavLink>
        </span>
    ));
}
