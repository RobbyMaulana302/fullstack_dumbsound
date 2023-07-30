import React, { useContext, useEffect, useState } from "react";
import { Button, Container, Navbar, Row, Col, Dropdown } from "react-bootstrap";
import BrandLogo from "../assets/img/BrandLogo.png";
import BrandName from "../assets/img/BrandName.png";
import Register from "./modal/Register";
import Login from "./modal/Login";
import { API, setAuthToken } from "../config/Api";
import { useNavigate } from "react-router-dom";
import { UserContext } from "../context/UserContext";
import Profile from "../assets/img/Profile.png";
import ImageAddMusic from "../assets/img/ImageAddMusic.png";
import ImageAddArtist from "../assets/img/ImageAddArtist.png";
import ImageLogout from "../assets/img/ImageLogout.png";
import ImagePayment from "../assets/img/ImagePayment.png";

function Header() {
  let navigate = useNavigate();
  const [showLogin, setShowLogin] = useState(false);
  const [showRegister, setShowRegister] = useState(false);

  const [state, dispatch] = useContext(UserContext);

  const logout = () => {
    dispatch({
      type: "LOGOUT",
    });
  };

  return (
    <>
      <div className="position-absolute container-fluid">
        <Navbar expand="lg" style={{ backgroundColor: "transparent" }}>
          <Container>
            <Navbar.Brand onClick={()=> navigate('/')} style={{cursor: "pointer"}}>
              <img src={BrandLogo} alt="BrandLogo" className="me-3" />
              <img src={BrandName} alt="BrandName" />
            </Navbar.Brand>
            <Navbar.Toggle aria-controls="basic-navbar-nav" />
            <Navbar.Collapse
              id="basic-navbar-nav"
              className="justify-content-end"
            >
              <Row>
                {state.isLogin === true ? (
                  state.user.role === "admin" ? (
                    <>
                      <Dropdown className="d-flex">
                        <Dropdown.Toggle
                          style={{ border: "none", background: "transparent" }}
                        >
                          <img src={Profile} alt="Profile" />
                        </Dropdown.Toggle>
                        <Dropdown.Menu style={{ background: "#3A3A3A" }}>
                          <Dropdown.Item
                            onClick={() => navigate("/admin/add-music")}
                          >
                            <img src={ImageAddMusic} alt="ImageAddMusic" />{" "}
                            <span className="text-white">Add Music</span>
                          </Dropdown.Item>
                          <Dropdown.Item
                            onClick={() => navigate("/admin/add-Artist")}
                          >
                            <img src={ImageAddArtist} alt="ImageAddArtist" />{" "}
                            <span className="text-white">Add Artist</span>
                          </Dropdown.Item>
                          <Dropdown.Item
                            onClick={() => navigate("/admin/list-Transaction")}
                          >
                            <img src={ImageAddMusic} alt="ImageAddMusic" />{" "}
                            <span className="text-white">List Transaction</span>
                          </Dropdown.Item>
                          <hr style={{ border: "3px solid white" }} />
                          <Dropdown.Item onClick={logout}>
                            <img src={ImageLogout} alt="ImageLogout" />{" "}
                            <span className="text-white">Logout</span>
                          </Dropdown.Item>
                        </Dropdown.Menu>
                      </Dropdown>
                    </>
                  ) : (
                    <>
                      <Dropdown>
                        <Dropdown.Toggle
                          style={{
                            border: "none",
                            background: "transparent",
                          }}
                        >
                          <img src={Profile} alt="Profile" />
                        </Dropdown.Toggle>
                        <Dropdown.Menu className="bg-dark">
                          <Dropdown.Item onClick={() => navigate("/user/payment")}>
                            <img src={ImagePayment} alt="ImagePayment" />{" "}
                            <span className="text-white">Pay</span>
                          </Dropdown.Item>
                          <hr style={{ border: "3px solid white" }} />
                          <Dropdown.Item onClick={logout}>
                            <img src={ImageLogout} alt="ImageLogout" />{" "}
                            <span className="text-white">Logout</span>
                          </Dropdown.Item>
                        </Dropdown.Menu>
                      </Dropdown>
                    </>
                  )
                ) : (
                  <>
                    <Col>
                      <Button
                        style={{
                          color: "#FFF",
                          textAlign: "center",
                          fontFamily: "Product Sans",
                          fontSize: "14px",
                          fontWeight: "700",
                          lineHeight: "normal",
                          width: "100px",
                          height: "30px",
                          borderRadius: "5px",
                          border: "1px solid #FFF",
                          backgroundColor: "transparent",
                        }}
                        onClick={() => setShowLogin(true)}
                      >
                        Login
                      </Button>
                    </Col>
                    <Col>
                      <Button
                        style={{
                          color: "#FFF",
                          textAlign: "center",
                          fontFamily: "Product Sans",
                          fontSize: "14px",
                          fontWeight: "700",
                          lineHeight: "normal",
                          width: "100px",
                          height: "30px",
                          borderRadius: "5px",
                          border: "1px solid #EE4622",
                          backgroundColor: "#EE4622",
                        }}
                        onClick={() => setShowRegister(true)}
                      >
                        Register
                      </Button>
                    </Col>
                  </>
                )}
              </Row>
            </Navbar.Collapse>
          </Container>
        </Navbar>
      </div>
      <Login show={showLogin} showLogin={setShowLogin} />
      <Register show={showRegister} showRegister={setShowRegister} />
    </>
  );
}

export default Header;
