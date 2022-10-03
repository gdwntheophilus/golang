import React, { useState } from "react";
import "./Login.css";
import PropTypes from "prop-types";
import axios from "axios";
import { useNavigate } from "react-router-dom";

function Login({ setToken }) {
  const navigate = useNavigate();
  const baseUrl = "http://127.0.0.1:8080/";
  const [username, setUserName] = useState();
  const [password, setPassword] = useState();
  const [LoginStatus, setLoginStatus] = useState("");

  const handleSubmit = (event) => {
    console.log(username);
    console.log(password);
    const payload = { password: password, emailid: username };
    console.log(payload);
    console.log(baseUrl + "user/getRegisteredUser");
    axios.post(baseUrl + "user/getRegisteredUser", payload).then((response) => {
      console.log(response.data.status);
      console.log(response.data.message);
      if (response.data.status === "error") {
        setLoginStatus(response.data.message);
        // setUserName("");
        // setPassword("");
      } else {
        console.log("--------setting the Token---------------");
        setToken(response.data.status);
        navigate("/contacts");
        // console.log(Token.getToken());
        // console.log(Token.getToken());
      }
    });

    event.preventDefault();
  };

  return (
    <div className="Auth-form-container">
      <form className="Auth-form" onSubmit={handleSubmit}>
        <div className="Auth-form-content">
          <h3 className="Auth-form-title">Sign In</h3>
          <div className="form-group mt-3">
            <label>Email address</label>
            <input
              type="email"
              className="form-control mt-1"
              placeholder="Enter email"
              onChange={(e) => setUserName(e.target.value)}
            />
          </div>
          <div className="form-group mt-3">
            <label>Password</label>
            <input
              type="password"
              className="form-control mt-1"
              placeholder="Enter password"
              onChange={(e) => setPassword(e.target.value)}
            />
          </div>
          <div className="d-grid gap-2 mt-3">
            <button type="submit" className="btn btn-primary">
              Submit
            </button>
          </div>
          <p className="forgot-password text-right mt-2 responseStatus">
            {LoginStatus}
            {/* <a>password?</a> */}
          </p>
        </div>
      </form>
    </div>
  );
}

export default Login;

Login.propTypes = {
  setToken: PropTypes.func.isRequired,
};
