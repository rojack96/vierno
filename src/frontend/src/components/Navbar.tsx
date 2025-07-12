import { NavLink } from "react-router-dom";
import "./Navbar.css"; // (opzionale)

const Navbar = () => {
    return (
        <nav className="navbar">
            <div className="nav-logo">Vierno</div>
            <ul className="nav-links">
                <li>
                    <NavLink to="/dashboard" end className={({ isActive }) => isActive ? "active" : ""}>
                        Home
                    </NavLink>
                </li>
                <li>
                    <NavLink to="/settings" className={({ isActive }) => isActive ? "active" : ""}>
                        Settings
                    </NavLink>
                </li>
            </ul>
        </nav>
    );
};

export default Navbar;
