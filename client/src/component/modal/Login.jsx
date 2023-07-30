import React, { useContext, useState } from "react";
import { Alert, Button, Form, Modal } from "react-bootstrap";
import { useNavigate } from "react-router-dom";
import { UserContext } from "../../context/UserContext";
import { useMutation, useQuery } from "react-query";
import { API, setAuthToken } from "../../config/Api";

export default function Login({ show, showLogin }) {
  let navigate = useNavigate();

  const [_, dispatch] = useContext(UserContext);

  const [message, setMessage] = useState(null);
  const [form, setForm] = useState({
    email: "",
    password: "",
  });

  const { email, password } = form;

  const handleChange = (e) => {
    setForm({
      ...form,
      [e.target.name]: e.target.value,
    });
  };

  const handleSubmit = useMutation(async (e) => {
    try {
      e.preventDefault();

      const response = await API.post("/login", form);
      console.log("login response", response.data.data.user.token);

      dispatch({
        type: "LOGIN_SUCCESS",
        payload: response.data.data.user,
      });

      setAuthToken(localStorage.token);

      if (response.data.data.user.role === "admin") {
        navigate("/admin/list-transaction");
        alert("admin :login success");
      } else {
        navigate("/user");
        alert("user :login success");
      }
      setForm({
        email: "",
        password: "",
      });
    } catch (error) {
      const alert = (
        <Alert variant="danger" className="py-1">
          Login failed
        </Alert>
      );
      setMessage(alert);
      console.log("login failed :", error);
    }
  });

  


  return (
    <div>
      <Modal
        show={show}
        onHide={showLogin}
        size="md"
        className="d-flex"
        centered
      >
        <Modal.Body
          style={{ background: "#1F1F1F", borderRadius: "5px" }}
          className=""
        >
          {message && message}
          <h1 className="text-white mb-4">Login</h1>
          <Form onSubmit={(e) => handleSubmit.mutate(e)}>
            <Form.Control
              type="email"
              className="mb-4"
              style={{ width: "350px" }}
              placeholder="Email ..."
              name="email"
              id="email"
              value={email}
              onChange={handleChange}
            />
            
            <Form.Control
              type="password"
              className="mb-4"
              style={{ width: "350px" }}
              placeholder="Password ..."
              name="password"
              id="password"
              value={password}
              onChange={handleChange}
            />
            <Button
              style={{
                width: "350px",
                height: "50px",
                background: "#EE4622",
                border: "none",
              }}
              type="submit"
            >
              Login
            </Button>
          </Form>
        </Modal.Body>
      </Modal>
    </div>
  );
}
