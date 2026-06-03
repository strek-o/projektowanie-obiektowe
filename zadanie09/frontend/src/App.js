import React from "react";
import Products from "./components/Products";
import Payments from "./components/Payments";

function App() {
  return (
    <div style={{ padding: "20px", fontFamily: "Arial" }}>
      <h1>Exercise 9</h1>
      <hr />
      <Products />
      <hr />
      <Payments />
    </div>
  );
}

export default App;
