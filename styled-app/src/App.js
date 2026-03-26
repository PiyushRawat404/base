import { BrowserRouter, Routes, Route } from "react-router-dom";
import Register from "./pages/Register/Register";
import StyledPage from "./pages/StyledPage/StyledPage";

function App() {
  return (
    <>
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<Register />} />
          <Route path="/styled" element={<StyledPage />} />
        </Routes>
      </BrowserRouter>
    </>
  );
}

export default App;