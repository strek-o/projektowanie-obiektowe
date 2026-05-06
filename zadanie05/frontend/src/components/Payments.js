import React, { useState } from "react";

const Payments = () => {
  const [amount, setAmount] = useState("");

  const handlePayment = (e) => {
    e.preventDefault();

    fetch("http://localhost:8080/payments", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ amount: parseFloat(amount) }),
    })
      .then((response) => {
        if (response.ok) {
          alert("Payment successful!");
          setAmount("");
        }
      })
      .catch((error) => console.error("Sending payment failed:", error));
  };

  return (
    <div>
      <h2>Payments</h2>
      <form onSubmit={handlePayment}>
        <label>
          To pay (PLN):
          <input
            type="number"
            step="0.01"
            value={amount}
            onChange={(e) => setAmount(e.target.value)}
            required
          />
        </label>
        <button type="submit" style={{ marginLeft: "10px" }}>
          Pay
        </button>
      </form>
    </div>
  );
};

export default Payments;
