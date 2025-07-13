import { useState } from "react";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import Login from "./pages/Login";
import Dashboard from "./pages/Dashboard";
import Layout from "./components/Layout";
import ProtectedRoute from "./components/ProtectedRoute";
import FilePage from "./pages/FilePage";
import FileViewer from "./pages/FileViewer";

function App() {
  const [isAuth, setIsAuth] = useState(false);

  return (
    <BrowserRouter>
      <Routes>
        <Route path="/login" element={<Login setIsAuth={setIsAuth} />} />

        <Route element={<ProtectedRoute isAuth={isAuth} />}>
          <Route path="/" element={<Layout />}>
            <Route path="/dashboard" element={<Dashboard />} />
            <Route path="/folders/:folder" element={<FilePage />} />
            <Route path="/folders/:folder/:file" element={<FileViewer />} />
          </Route>
        </Route>
      </Routes>
    </BrowserRouter>
  );
}

export default App;
