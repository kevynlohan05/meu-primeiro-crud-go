import React, { useContext, useState } from 'react'
import { useNavigate } from 'react-router-dom'
import Linked from './Linked'
import { MdDashboard, MdCall } from "react-icons/md"
import { BsFillGearFill } from "react-icons/bs"
import { BiSolidExit } from "react-icons/bi"
import { ThemeContext } from '../contexts/ThemeContext'
import { GoSidebarExpand } from "react-icons/go";
import { GoSidebarCollapse } from "react-icons/go";

const Sidebar = ({collapsed, setCollapsed}) => {
  const navigate = useNavigate()
  const { theme } = useContext(ThemeContext)
  

  return (
    <aside className={`flex flex-col min-h-screen ${collapsed ? 'w-20' : 'w-72'} transition-all duration-300 fixed bg-slate-100 dark:bg-gray-800 p-4 shadow-xl z-10`}>
      {/* Botão de recolher/expandir */}
      <button
        onClick={() => setCollapsed(!collapsed)}
        className="self-end text-gray-500 dark:text-white mb-4"
      >
        {collapsed ? <GoSidebarCollapse size={20} /> : <GoSidebarExpand size={20} />}
      </button>

      {/* Logo */}
      {!collapsed && (
        <img className='p-4' src={theme === 'dark' ? './img/logo_white.png' : './img/logo_black.png'} alt="logo" />
      )}

      <hr className="w-full border-gray-400 dark:border-white mb-6" />

      {/* Links */}
      <nav className="flex flex-col gap-4 w-full">
        <Linked href={'/Home'} icon={<MdDashboard />}>
          {!collapsed && 'Dashboard'}
        </Linked>
        <Linked href={'/Desk'} icon={<MdCall />}>
          {!collapsed && 'Chamado'}
        </Linked>
        <Linked href={'/Configuração'} icon={<BsFillGearFill />}>
          {!collapsed && 'Configurações'}
        </Linked>
        <Linked href={'/login'} icon={<BiSolidExit />}>
          {!collapsed && 'Sair'}
        </Linked>
      </nav>
    </aside>
  )
}

export default Sidebar
