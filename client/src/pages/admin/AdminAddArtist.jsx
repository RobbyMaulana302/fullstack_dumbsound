import React, { useState } from "react";
import { Alert, Button, Col, Container, Form, Row } from "react-bootstrap";
import { useMutation } from "react-query";
import { API } from "../../config/Api";

function AdminAddArtist() {
  const [message, setMessage] = useState(null);

  const [form, setForm] = useState({
    name: "",
    old: "",
    type: "",
    start_career: "",
  });

  const { name, old, type, start_career } = form;
  console.log(type);

  const handleChange = (e) => {
    setForm({
      ...form,
      [e.target.name]: e.target.value,
    });
  };

  const handleSubmit = useMutation(async (e) => {
    try {
      e.preventDefault();

      const response = await API.post("/artist", form);
      console.log(response);

      const alert = (
        <Alert variant="success" className="py-1">
          Add Ticket Success
        </Alert>
      );
      setMessage(alert);

      setForm({
        name: "",
        old: "",
        type: "",
        start_career: "",
      });
    } catch (error) {
      const alert = (
        <Alert variant="danger" className="py-1">
          Add Artist Failed
        </Alert>
      );
      setMessage(alert);
      console.log("Add ticket failed : ", error);
    }
  });

  return (
    <div>
      <Container fluid style={{ background: "#161616", height: "100vh" }}>
        <Container style={{ background: "#161616", height: "100vh" }}>
          {message && message}
          <div style={{ height: "100px" }}></div>
          <div className="text-white">
            <h1 className="mb-5">Add Artist</h1>
            <Form onSubmit={(e) => handleSubmit.mutate(e)}>
              <Row>
                <Col>
                  <Form.Control
                    type="text"
                    placeholder="Name"
                    style={{
                      border: "2px solid #D2D2D2",
                      background: "rgba(210, 210, 210, 0.25)",
                      color: "white",
                    }}
                    className="mb-3"
                    id="name"
                    name="name"
                    value={name}
                    onChange={handleChange}
                  />
                  <Form.Control
                    type="text"
                    placeholder="Old"
                    className="mb-3"
                    style={{
                      border: "2px solid #D2D2D2",
                      background: "rgba(210, 210, 210, 0.25)",
                      color: "white",
                    }}
                    id="old"
                    name="old"
                    value={old}
                    onChange={handleChange}
                  />
                  <Form.Select
                    type="text"
                    placeholder="Type"
                    className="mb-3"
                    style={{
                      border: "2px solid #D2D2D2",
                      background: "rgba(210, 210, 210, 0.25)",
                      color: "white",
                    }}
                    id="type"
                    name="type"
                    value={type}
                    onChange={handleChange}
                  >
                    <option value="">-- select --</option>
                    <option value="solo" onChange={handleChange} selected>
                      Solo
                    </option>
                    <option value="band" onChange={handleChange} selected>
                      Band
                    </option>
                  </Form.Select>
                  <Form.Control
                    type="date"
                    placeholder="date"
                    className="mb-3"
                    style={{
                      border: "2px solid #D2D2D2",
                      background: "rgba(210, 210, 210, 0.25)",
                      color: "white",
                    }}
                    id="start_career"
                    name="start_career"
                    value={start_career}
                    onChange={handleChange}
                  />
                </Col>
              </Row>
              <Container className="d-flex justify-content-center">
                <Button
                  style={{
                    width: "350px",
                    height: "40px",
                    borderRadius: "5px",
                    background: "#F58033",
                    border: "none",
                  }}
                  className="text-center"
                  type="submit"
                >
                  Add Artist
                </Button>
              </Container>
            </Form>
          </div>
        </Container>
      </Container>
    </div>
  );
}

export default AdminAddArtist;
