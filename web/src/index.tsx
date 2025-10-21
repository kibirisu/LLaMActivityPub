import React from "react";
import { createRoot } from "react-dom/client";
import "./index.css";
import MainFeed from "./feed/MainFeed";

const container = document.getElementById("root");
const root = createRoot(container!);

root.render(
  <React.StrictMode>
    <div className="min-h-screen bg-gray-50 text-gray-900 flex justify-center">
      <div className="w-[20%]">
        <MainFeed />
      </div>
    </div>
  </React.StrictMode>
);

