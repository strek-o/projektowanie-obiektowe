import React, { useState } from "react";

const API_URL =
  window.__API_URL__ ||
  process.env.REACT_APP_API_URL ||
  "http://localhost:8080";

function App() {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [message, setMessage] = useState("");
  const [loggedUser, setLoggedUser] = useState(null);

  const handleSubmit = async (e) => {
    e.preventDefault();
    setMessage("");
    try {
      const res = await fetch(`${API_URL}/login`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ username, password }),
      });
      const data = await res.json();
      if (res.ok && data.status === "ok") {
        setLoggedUser(data.username);
        setMessage("Success!");
      } else {
        setMessage(data.message || "Failure!");
      }
    } catch (err) {
      setMessage("ERROR: " + err.message);
    }
  };

  const handleLogout = () => {
    setLoggedUser(null);
    setUsername("");
    setPassword("");
    setMessage("");
  };

  if (loggedUser) {
    return (
      <div style={{ padding: 20 }}>
        <h2>Welcome, {loggedUser}!</h2>
        <button onClick={handleLogout}>Log Out</button>
      </div>
    );
  }

  return (
    <div style={{ padding: 20 }}>
      <h2>Log In</h2>
      <div>
        <p>Available accounts:</p>
        <ul>
          <li>user / user</li>
          <li>admin / admin</li>
        </ul>
      </div>
      <form onSubmit={handleSubmit}>
        <div>
          <input
            type="text"
            placeholder="Username"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
          />
        </div>
        <div>
          <input
            type="password"
            placeholder="Password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
          />
        </div>
        <button type="submit">Submit</button>
      </form>
      {message && <p>{message}</p>}
    </div>
  );
}

export default App;
