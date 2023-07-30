import React, { useEffect, useState } from "react";
import { Alert, Button, Form, Modal } from "react-bootstrap";
import { useMutation, useQuery } from "react-query";
import { API } from "../../config/Api";

function Register({ show, showRegister }) {
  const [message, setMessage] = useState(null);
  const [emailUser, setEmailUser] = useState([]);
  const [form, setForm] = useState({
    email: "",
    password: "",
    fullname: "",
    gender: "",
    phone: "",
    address: "",
  });

  const { email, password, fullname, gender, phone, address } = form;

  const handleChange = (e) => {
    setForm({
      ...form,
      [e.target.name]: e.target.value,
    });
  };

  const handleSubmit = useMutation(async (e) => {
    try {
      e.preventDefault();

      const response = await API.post("/register", form);
      console.log(response);

      const alert = (
        <Alert variant="success" className="py-1">
          Register Success!
        </Alert>
      );
      setMessage(alert);

      setForm({
        email: "",
        password: "",
        fullname: "",
        gender: "",
        phone: "",
        address: "",
      });
    } catch (error) {
      const alert = (
        <Alert variant="danger" className="py-1">
          Failed to Register!
        </Alert>
      );
      setMessage(alert);
      console.log("Register failed :", error);
    }
  });

  // const { data: user } = useQuery(
  //   "userCache",
  //   async () => {
  //     const response = await API.get("/users");
  //     return response.data.data;
  //   }
  // );

  // console.log(user);
  // const plise = user.map((data) => data?.email)

  // setEmailUser(plise);
  // console.log(emailUser);

  return (
    <div>
      <Modal
        show={show}
        onHide={showRegister}
        size="md"
        className="d-flex"
        centered
      >
        <Modal.Body
          style={{ background: "#1F1F1F", borderRadius: "5px" }}
          className=""
        >
          {message && message}
          <h1 className="text-white mb-4">Register</h1>
          <Form onSubmit={(e) => handleSubmit.mutate(e)}>
            <Form.Control
              type="email"
              className="mb-4"
              style={{ width: "350px" }}
              placeholder="Email"
              id="email"
              name="email"
              value={email}
              onChange={handleChange}
            />
            {/* <div>
              {emailUser.includes(email) ? (
                <div></div>
              ) : (
                <span className="text-danger">email sudah digunakan</span>
              )}
            </div> */}
            <Form.Control
              type="password"
              className="mb-4"
              style={{ width: "350px" }}
              placeholder="Password"
              id="password"
              name="password"
              value={password}
              onChange={handleChange}
            />
            <Form.Control
              type="text"
              className="mb-4"
              style={{ width: "350px" }}
              placeholder="Full Name"
              id="fullname"
              name="fullname"
              value={fullname}
              onChange={handleChange}
            />
            <Form.Select
              className="mb-4"
              style={{ width: "350px" }}
              id="gender"
              name="gender"
              value={gender}
              onChange={handleChange}
            >
              <option value="" hidden>
                Gender
              </option>
              <option value="lanang">Lanang</option>
              <option value="wadon">Wadon</option>
            </Form.Select>
            <Form.Control
              type="number"
              className="mb-4"
              style={{ width: "350px" }}
              placeholder="Phone"
              id="phone"
              name="phone"
              value={phone}
              onChange={handleChange}
            />
            <Form.Control
              as="textarea"
              className="mb-4"
              style={{ width: "350px" }}
              placeholder="Address"
              id="address"
              name="address"
              value={address}
              onChange={handleChange}
            />
            {/* {emailUser.includes(email) ? (
              <Button
                style={{
                  width: "350px",
                  height: "50px",
                  background: "#EE4622",
                  border: "none",
                }}
                onClick={() => alert("Email Sudah digunakan")}
              >
                Register
              </Button>
            ) : (
              
            )} */}

            <Button
              style={{
                width: "350px",
                height: "50px",
                background: "#EE4622",
                border: "none",
              }}
              type="submit"
            >
              Register
            </Button>
          </Form>
        </Modal.Body>
      </Modal>
    </div>
  );
}

export default Register;
