import React from "react";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import Nav from "../src/etc/Nav";

import Register from "./pages/register/Register";

import Login from "./pages/login/Login";

function App() {
  return (
    <BrowserRouter>
      <div className="App">
        {/* <Nav /> */}

        <Routes>
          <Route path="/" element={<Register />} />
          <Route path="/login" element={<Login />} />
          <Route path="/register" element={<Register />} />
        </Routes>
      </div>
    </BrowserRouter>
  );
}

export default App;
