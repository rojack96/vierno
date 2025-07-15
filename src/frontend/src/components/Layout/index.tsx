import { Outlet } from "react-router-dom";
import Navbar from "../Navbar";

const Layout = () => {
    return (
        <>
            <Navbar />
            <main style={{ padding: "2rem" }}>
                <Outlet />
            </main>
        </>
    );
};

export default Layout;
