import { useState } from "react";
import Home from "./pages/Home";
import Cart from "./pages/Cart";

function App() {
  const [page, setPage] = useState("home");

  return (
    <div>
      <button onClick={() => setPage("home")}>Home</button>
      <button onClick={() => setPage("cart")}>Cart</button>

      {page === "home" ? <Home /> : <Cart />}
    </div>
  );
}

export default App;