import React from "react";
import { Modal } from "react-bootstrap";
import { useMutation } from "react-query";
import { useNavigate } from "react-router-dom";
import { API } from "../../config/Api";

function ModalPayment({ show, showPayment }) {
  let navigate = useNavigate();

  const handleSubmit = useMutation(async (e) => {
    try {
      e.preventDefault;
      
      let form = new FormData()
      form.set("status", "pending")
      const response = await API.post("/transaction", form);
      console.log(response);
      navigate("/user/payment");
    } catch (error) {
      console.log(error);
    }
  });
  return (
    <div>
      <Modal
        show={show}
        onHide={showPayment}
        size="md"
        className="d-flex"
        centered
      >
        <Modal.Body
          style={{ background: "#1F1F1F", borderRadius: "5px" }}
          className="text-center"
        >
          <h2 className="text-white">
            Bayar sekarang dan nikmati streaming music yang kekinian dari
          </h2>
          <h2 className="text-danger">klik disini ⬇</h2>
          <h2
            className="text-white"
            style={{ cursor: "pointer" }}
            onClick={(e) => handleSubmit.mutate(e)}
          >
            <span style={{ color: "#EE4622" }}>DUMB</span>SOUND
          </h2>
        </Modal.Body>
      </Modal>
    </div>
  );
}

export default ModalPayment;
