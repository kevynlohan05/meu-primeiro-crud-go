import React, { useState } from "react";
import Sidebar from "../components/Sidebar";
import ThemeToggle from "../components/ThemeToggle";
import { UserRegister } from "../components/UserRegister";
import { Outlet } from "react-router-dom";

export const Layout = () => {
  const [collapsed, setCollapsed] = useState(false)

  return (
    <>  
    <ThemeToggle />
    <Sidebar collapsed={collapsed} setCollapsed={setCollapsed} />
    
      <div className={`${collapsed ? "ml-20" : "ml-72"} transition-all duration-300 min-h-screen bg-slate-200 dark:bg-gray-900 text-gray-900 dark:text-black px-8`}>
       
       <Outlet />
       

      </div>
    </>
  );
};
