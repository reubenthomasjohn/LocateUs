import SignInForm from "./components/SignInForm";
import { BrowserRouter, Routes, Route, useLocation } from "react-router-dom";
import Dashboard from "./components/Dashboard";
import { useEffect } from "react";

function App() {
  return (
    <BrowserRouter>
      <TitleUpdater /> {/* This will update the title based on route */}
      <Routes>
        <Route path="/" element={<SignInForm />} />
        <Route path="/dashboard" element={<Dashboard />} />
      </Routes>
    </BrowserRouter>
  );
}

// Component to update the tab title based on the route
function TitleUpdater() {
  const location = useLocation();

  useEffect(() => {
    const titles: Record<string, string> = {
      "/": "Sign In - CCC Heatmap",
      "/dashboard": "Dashboard - CCC Heatmap",
    };

    document.title = titles[location.pathname] || "CCC Heatmap";
  }, [location.pathname]);

  return null; // This component doesn't render anything
}

export default App;
