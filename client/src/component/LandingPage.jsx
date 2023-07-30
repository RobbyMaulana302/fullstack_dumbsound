import React from "react";
import ImageLandingPage from "../assets/img/ImageLandingPage.png";
import Playlist from "./Playlist";
function LandingPage() {
  return (
    <>
      <div
        style={{ backgroundImage: `url(${ImageLandingPage})`, height: "512px" }}
        className="d-flex justify-content-center align-items-center"
      >
        <div>
          <div>
            <h1
              style={{
                color: "#FFF",
                fontFamily: "Sans-serif",
                fontSize: "48px",
                fontStyle: "normal",
                fontWeight: "400",
                lineHeight: "normal",
              }}
              className="text-center"
            >
              Connect on DumbSound
            </h1>
            <h6
              style={{
                color: "#FFF",
              }}
            >
              Discovery, Stream, and share a constantly expanding mix of music
              from emerging and major artists around the world
            </h6>
          </div>
        </div>
      </div>
      <Playlist />
    </>
  );
}

export default LandingPage;
