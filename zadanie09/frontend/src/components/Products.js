import React, { useState, useEffect } from "react";

const Products = () => {
  const [products, setProducts] = useState([]);

  useEffect(() => {
    fetch("https://zadanie09-backend-strek.azurewebsites.net/products")
      .then((response) => response.json())
      .then((data) => setProducts(data))
      .catch((error) => console.error("Fetching products failed:", error));
  }, []);

  return (
    <div>
      <h2>Products</h2>
      <ul>
        {products.map((product) => (
          <li key={product.id}>
            {product.name} - {product.price.toFixed(2)} PLN
          </li>
        ))}
      </ul>
    </div>
  );
};

export default Products;
