import React from "react";
import ImageLandingPage from "../assets/img/ImageLandingPage.png";
import Playlist from "./Playlist";
// import ReactJkMusicPlayer from 'react-jinke-music-player'
// import 'react-jinke-music-player/assets/index.css'
import AudioPlayer from "react-h5-audio-player";
import "react-h5-audio-player/lib/styles.css";
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
      {/* <ReactJkMusicPlayer
        mode="full"
        autoPlay={false}
        showDownload={false}
        showThemeSwitch={true}
        toggleMode={false}
        responsive={false}
        glassBg={true}
        audioLists={[
          {
            name: "Song 1",
            singer: "Artist 1",
            cover: "path/to/cover1.jpg",
            musicSrc: "path/to/song1.mp3",
          },
          {
            name: "Song 2",
            singer: "Artist 2",
            cover: "path/to/cover2.jpg",
            musicSrc: "path/to/song2.mp3",
          },
          // Tambahkan lebih banyak lagu ke dalam audioLists sesuai kebutuhan
        ]}
      /> */}


    </>
  );
}

export default LandingPage;
