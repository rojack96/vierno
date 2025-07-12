import React from "react";
import { Navigate, Outlet } from "react-router-dom";

interface ProtectedRouteProps {
    isAuth: boolean;
}

const ProtectedRoute: React.FC<ProtectedRouteProps> = ({ isAuth }) => {
    return isAuth ? <Outlet /> : <Navigate to="/" replace />;
};

export default ProtectedRoute;
