import React, { useState } from "react";
import { Alert, Button, Col, Container, Form, Row } from "react-bootstrap";
import ImageInputThumbnail from "../../assets/img/ImageInputThumbnail.png";
import { useMutation, useQuery } from "react-query";
import { API } from "../../config/Api";

function AdminAddMusic() {
  const { data: artistData } = useQuery("artistCache", async () => {
    const response = await API.get("/artists");
    return response.data.data.artist;
  });

  const [message, setMessage] = useState(null);

  const [form, setForm] = useState({
    title: "",
    thumbnail: "",
    year: "",
    artist_id: "",
    attache: "",
  });

  const { title, thumbnail, year, artist_id, attache } = form;

  const handleChange = (e) => {
    setForm({
      ...form,
      [e.target.name]:
        e.target.type === "file" ? e.target.files : e.target.value,
    });
  };

  const handleSubmit = useMutation(async (e) => {
    try {
      e.preventDefault();

      const config = {
        headers: {
          "Content-type": "multipart/form-data",
        },
      };

      const formData = new FormData();
      formData.set("title", form.title);
      formData.set("thumbnail", form.thumbnail[0], form.thumbnail[0].name);
      formData.set("year", form.year);
      formData.set("artist_id", form.artist_id);
      formData.set("attache", form.attache[0], form.attache[0].name);

      const response = await API.post("/music", formData, config);
      console.log("response music post", response);

      const alert = (
        <Alert variant="success" className="py-1">
          Add Music Success
        </Alert>
      );

      setMessage(alert);

      setForm({
        title: "",
        thumbnail: "",
        year: "",
        artist_id: "",
        attache: "",
      });
    } catch (error) {
      const alert = (
        <Alert variant="danger" className="py-1">
          Add Music Failed
        </Alert>
      );
      setMessage(alert);
      console.log("Add Music Failed : ", error);
    }
  });

  return (
    <div>
      <Container fluid style={{ background: "#161616", height: "100vh" }}>
        <Container style={{ background: "#161616", height: "100vh" }}>
          <div style={{ height: "100px" }}></div>
          <div className="text-white">
            <h1 className="mb-5">Add Music</h1>
            <Form onSubmit={(e) => handleSubmit.mutate(e)}>
              <Row>
                <Col>
                  <Row>
                    <Col xs={9} className="me-3 text-white">
                      <Form.Control
                        type="text"
                        placeholder="Title"
                        style={{
                          border: "2px solid #D2D2D2",
                          background: "rgba(210, 210, 210, 0.25)",
                          color: "white",
                        }}
                        className="mb-3"
                        id="title"
                        name="title"
                        value={title}
                        onChange={handleChange}
                      />
                    </Col>
                    <Col>
                      <Form.Group className="text-center">
                        <Form.Label
                          style={{
                            cursor: "pointer",
                            borderRadius: "5px",
                            border: "2px solid #D2D2D2",
                            background: "rgba(210, 210, 210, 0.25)",
                            color: "white",
                            width: "213px",
                          }}
                          for="thumbnail"
                        >
                          <span>Attache Thumbnail</span>{" "}
                          <img
                            src={ImageInputThumbnail}
                            alt="ImageInputThumbnail"
                          />
                          <Form.Control
                            type="file"
                            hidden
                            id="thumbnail"
                            name="thumbnail"
                            onChange={handleChange}
                          />
                        </Form.Label>
                      </Form.Group>
                    </Col>
                  </Row>
                  <Form.Control
                    type="text"
                    placeholder="Year"
                    className="mb-3"
                    style={{
                      border: "2px solid #D2D2D2",
                      background: "rgba(210, 210, 210, 0.25)",
                      color: "white",
                    }}
                    id="year"
                    name="year"
                    value={year}
                    onChange={handleChange}
                  />
                  <Form.Select
                    type="text"
                    placeholder="Singer"
                    className="mb-3"
                    style={{
                      border: "2px solid #D2D2D2",
                      background: "rgba(210, 210, 210, 0.25)",
                      color: "white",
                    }}
                    id="artist_id"
                    name="artist_id"
                    value={artist_id}
                    onChange={handleChange}
                  >
                    <option value="">-- artist --</option>\
                    {artistData?.map((data) => (
                      <option
                        value={data?.id}
                        className="bg-dark"
                        key={data?.id}
                      >
                        {data?.name}
                      </option>
                    ))}
                  </Form.Select>
                  <Form.Group className="mb-3">
                    <Form.Label
                      style={{
                        cursor: "pointer",
                        width: "119px",
                        height: "30px",
                        borderRadius: "5px",
                        border: "2px solid #D2D2D2",
                        background: "rgba(210, 210, 210, 0.25)",
                      }}
                      className="text-center"
                      for="attache"
                    >
                      Attache
                    </Form.Label>
                    <Form.Control
                      id="attache"
                      type="file"
                      name="attache"
                      hidden
                      onChange={handleChange}
                    />
                  </Form.Group>
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
                  Add Song
                </Button>
              </Container>
            </Form>
          </div>
        </Container>
      </Container>
    </div>
  );
}

export default AdminAddMusic;
