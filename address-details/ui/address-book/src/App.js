import "./App.css";
import Login from "./components/login/Login";
import "bootstrap/dist/css/bootstrap.min.css";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import Contacts from "./components/contacts/contacts";
import Token from "./components/token/Token";
import { useEffect } from "react";
// import { useState } from "react";

// function setToken(userToken) {
//   sessionStorage.setItem("token", JSON.stringify(userToken));
// }
// function getToken() {
//   const tokenString = sessionStorage.getItem("token");
//   const userToken = JSON.parse(tokenString);
//   return userToken?.token;
// }

function App() {
  const { token, setToken } = Token();

  useEffect(() => {
    // setToken(sessionStorage.getItem("token"));
    console.log(sessionStorage.getItem("token"));
  }, []);
  // const [token, setToken] = useState();
  // const token = getToken();

  if (!token) {
    return (
      <BrowserRouter>
        <Routes>
          <Route path="/contacts">
            <Contacts />
          </Route>
          {/* <Route path="/preferences">
            <Preferences />
          </Route> */}
        </Routes>
      </BrowserRouter>
    );
  } else {
    return (
      // <div className="App">
      //   <BrowserRouter>
      //     <Routes>
      //       <Route path="/contacts" element={<Login />} />
      //     </Routes>
      //   </BrowserRouter>
      // </div>
      <div className="wrapper">
        <h1>Application</h1>
        <BrowserRouter>
          <Routes>
            <Route path="/contacts">
              <Contacts />
            </Route>
            {/* <Route path="/preferences">
            <Preferences />
          </Route> */}
          </Routes>
        </BrowserRouter>
      </div>
    );
  }
}

export default App;
