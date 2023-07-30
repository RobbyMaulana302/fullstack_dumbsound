import React, { useEffect } from "react";
import { Button, Container } from "react-bootstrap";
import BrandName from "../../assets/img/BrandName.png";
import { useNavigate } from "react-router-dom";
import { useMutation, useQuery } from "react-query";
import { API } from "../../config/Api";

function Payment() {
  const navigate = useNavigate();

  let { data: transaction } = useQuery("transactionCache", async () => {
    const response = await API.get("/transaction");
    return response.data.data.transaction;
  });
  console.log(transaction);

  console.log(new Date(transaction?.due_date));
 
  const countDuration = (dueDate) => {
    let start_date = Date.now()
    let due_date = new Date (dueDate)
    let distance = due_date - start_date
    let milisecond = 1000
    let secondInHour = 3600
    let hourInDay = 24

    let distanceDay = Math.floor(distance / (milisecond * secondInHour * hourInDay))

    return `${distanceDay}/hari`
  }

  useEffect(()=> {
    const midtransScriptUrl = "https://app.sandbox.midtrans.com/snap/snap.js";
    const myMidtransClientKey = import.meta.env.VITE_MIDTRANS_CLIENT_KEY;
    let scriptTag = document.createElement("script");

    scriptTag.src = midtransScriptUrl;
    scriptTag.setAttribute("data-client-key", myMidtransClientKey);

    document.body.appendChild(scriptTag);

    return () => {
      document.body.removeChild(scriptTag);
    };
  }, [])

  let id = transaction?.id
  console.log("Transaction ID", id);

  const handleBuy = useMutation(async () => {
    try {
      const response = await API.get(`/getpayment/${id}`)
      
      const token = response.data.data.token
      console.log("ini token payment", token)

      window.snap.pay(token, {
        onSuccess: function (result) {
          console.log(result);
          navigate("/user");
        },
        onPending: function (result) {
          console.log(result);
          navigate("/user");
        },
        onError: function (result) {
          console.log(result);
          navigate("/user");
        },
        onClose: function () {
          alert("tutup")
        },
      });
    } catch (error) {
      console.log(error);
    }
  })
  return (
    <div>
      {transaction?.status === "success" ? (
        <Container
          fluid
          style={{ background: "#161616", height: "100vh" }}
          className="justify-content-center d-flex align-items-center"
        >
          <Container className="justify-content-center d-flex align-items-center">
            <div className="text-center text-white p-5">
              <h1 className="my-3">Premium anda tersisa</h1>
              <h2 className="my-5">
                {countDuration(transaction?.due_date)}
              </h2>
            </div>
          </Container>
        </Container>
      ) : (<Container
        fluid
        style={{ background: "#161616", height: "100vh" }}
        className="justify-content-center d-flex align-items-center"
      >
        <Container className="justify-content-center d-flex align-items-center">
          <div className="text-center text-white p-5">
            <h1 className="my-3">Premium</h1>
            <h2 className="my-5">
              Bayar sekarang dan nikmati streaming music yang kekinian dari
              <img src={BrandName} alt="BrandName" />
            </h2>
            <Button
              style={{
                width: "350px",
                height: "40px",
                borderRadius: "5px",
                border: "none",
                background: "#F58033",
              }}

              onClick={() => handleBuy.mutate(id)}
            >
              Bayar
            </Button>
          </div>
        </Container>
      </Container>)}
    </div>
  );
}

export default Payment;
