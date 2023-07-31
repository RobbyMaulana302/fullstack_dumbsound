import { useContext } from "react";
import { UserContext } from "../context/UserContext";
import { Navigate, Outlet } from "react-router-dom";


export function PrivateRouteAdmin() {
  const [state] = useContext(UserContext);

  if (state.user.role === "admin") {
    return <Outlet />;
  }
  return <Navigate to="/" />;
}

export function PrivateRouteUser() {
  const [state] = useContext(UserContext);

  if (state.user.role === "user") {
    return <Outlet />;
  }
  return <Navigate to="/" />;
}
