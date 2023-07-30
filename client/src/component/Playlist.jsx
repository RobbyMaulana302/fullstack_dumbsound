import React, { useContext, useState } from "react";
import { Card, Carousel, Col, Container, Row } from "react-bootstrap";
import ImageCard from "../assets/img/ImageCard.png";
import { UserContext } from "../context/UserContext";
import { useQuery } from "react-query";
import { API } from "../config/Api";
import Login from "./modal/Login";
import ModalPayment from "./modal/ModalPayment";
import AudioPlayer from "react-h5-audio-player";
import "react-h5-audio-player/lib/styles.css";

function Playlist() {
  const [state, _] = useContext(UserContext);
  const [showLogin, setShowLogin] = useState(false);
  const [showPayment, setShowPayment] = useState(false)
  const [song, setSong] = useState()


  let { data: music } = useQuery("musicCache", async () => {
    const response = await API.get("/musics");
    return response.data.data.music;
  });

  console.log(state.user);
  return (
    <>
      <Container fluid style={{ background: "#161616" }}>
        <Container className="p-5">
          <div className="text-center">
            <h5
              style={{
                color: "#EE4622",
                fontSize: "24px",
                fontStyle: "normal",
                fontWeight: 700,
                lineHeight: "normal",
              }}
              className="py-5"
            >
              Dengarkan Dan Rasakan
            </h5>
          </div>
          <Row>
            {music?.map((data) => (
              <Col xs={3} className="mb-3">
                <Card
                  style={{
                    width: "192px",
                    height: "240px",
                    borderRadius: "10px",
                    background: "#3A3A3A",
                    cursor: "pointer",
                  }}
                  className="p-2"
                  onClick={state.isLogin === true ? (state.user.listAs === true ? (() => setSong(data?.attache)) : (() => setShowPayment(true))) : () => setShowLogin(true)}
                >
                  <Card.Img src={data?.thumbnail} />
                  <div className="d-flex justify-content-between mt-2 text-white">
                    <h6>{data.title}</h6>
                    <h6>{data.year}</h6>
                  </div>
                  <div className="text-white">
                    <h6>{data.artist.name}</h6>
                  </div>
                </Card>
              </Col>
            ))}
          </Row>
        </Container>
      </Container>
      <AudioPlayer
        autoPlay
        src={song}
        onPlay={(e) => console.log("onPlay")}
      />
      <Login show={showLogin} showLogin={setShowLogin}/>
      <ModalPayment show={showPayment} showPayment={setShowPayment}/>

    </>
  );
}

export default Playlist;
