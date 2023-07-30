import React, { useContext, useEffect, useState } from "react";
import "bootstrap/dist/css/bootstrap.min.css";
import Header from "./component/Header";
import { Route, Routes, useNavigate } from "react-router-dom";
import LandingPage from "./component/LandingPage";
import Payment from "./pages/user/Payment";
import AdminListTransaction from "./pages/admin/AdminListTransaction";
import AdminAddMusic from "./pages/admin/AdminAddMusic";
import AdminAddArtist from "./pages/admin/AdminAddArtist";
import { UserContext } from "./context/UserContext";
import { API, setAuthToken } from "./config/Api";
import {
  PrivateRouteAdmin,
  PrivateRouteLogin,
  PrivateRouteUser,
} from "./private/PrivateRoute";

function App() {
  let navigate = useNavigate();

  const [state, dispatch] = useContext(UserContext);
  const [isLoading, setIsLoading] = useState(null);
  useEffect(() => {
    if (!isLoading) {
      if (state.isLogin === false) {
        navigate("/");
      }
    }
  }, [isLoading]);

  useEffect(() => {
    if (localStorage.token) {
      setAuthToken(localStorage.token);
      checkUser();
    } else {
      setIsLoading(false);
    }
  }, []);

  const checkUser = async () => {
    try {
      const response = await API.get("/check-auth");

      let payload = response.data.data.user;
      payload.token = localStorage.token;

      dispatch({
        type: "USER_SUCCESS",
        payload,
      });

      setIsLoading(false);
    } catch (error) {
      console.log(error);

      dispatch({
        type: "AUTH_ERROR",
      });

      setIsLoading(false);
    }
  };

  return (
    <>
      <Header />
      {isLoading ? null : (
        <Routes>
          <Route exact path="/" element={<LandingPage />} />
          <Route element={<PrivateRouteLogin />} >
          <Route element={<PrivateRouteUser />}>
            <Route path="/user" element={<LandingPage />} />
            <Route path="/user/payment" element={<Payment />} />
          </Route>
          <Route element={<PrivateRouteAdmin />}>
            <Route
              path="/admin/list-transaction"
              element={<AdminListTransaction />}
            />
            <Route path="/admin/add-music" element={<AdminAddMusic />} />
            <Route path="/admin/add-artist" element={<AdminAddArtist />} />
          </Route>
          </Route>
        </Routes>
      )}
    </>
  );
}

export default App;
