import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import { BrowserRouter, Route, Routes, Navigate } from "react-router-dom";
import "./index.css";
import SalesPage from "./pages/SalesPage";
import PurchaseOrderPage from "./pages/PurchaseOrderPage";

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <BrowserRouter>
      <Routes>
        <Route path="/sales-page" element={<SalesPage />} />
        <Route path="/purchase-order-page" element={<PurchaseOrderPage />} />
        <Route path="*" element={<Navigate to="/sales-page" replace />} />
      </Routes>
    </BrowserRouter>
  </StrictMode>
);
