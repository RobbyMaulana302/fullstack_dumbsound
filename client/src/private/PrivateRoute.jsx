import { useContext } from "react";
import { UserContext } from "../context/UserContext";
import { Navigate, Outlet } from "react-router-dom";

export function PrivateRouteLogin() {
  const [state] = useContext(UserContext);

  if (!state.isLogin) {
    return <Navigate to={"/"} />;
  }

  return <Outlet />;
}

export function PrivateRouteUser() {
  const [state] = useContext(UserContext);

  console.log(state.user);
  if (state.user.role === "admin") {
    return <Navigate to="/admin/list-transaction" />;
  }
  return <Outlet />;
}

export function PrivateRouteAdmin() {
  const [state] = useContext(UserContext);

  if (state.user.role !== "user") {
    return <Navigate to="/" />;
  }
  return <Outlet />;
}
