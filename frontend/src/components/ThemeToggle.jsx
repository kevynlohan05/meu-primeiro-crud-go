import React, { useContext, useEffect, useState } from 'react'
import { MdSunny } from "react-icons/md";
import { IoIosMoon } from "react-icons/io";
import { ThemeContext } from '../contexts/ThemeContext';

const ThemeToggle = () => {
  const {theme, toggleTheme} = useContext(ThemeContext)

  return (
    <>
    <button
      onClick={toggleTheme}
      className="absolute top-4 right-4 p-2 bg-gray-300 dark:bg-gray-700 text-gray-900 dark:text-white rounded-md shadow hover:scale-105 transition"
    >
      {theme === 'dark' ? <MdSunny /> : <IoIosMoon />}
    </button>

    
    </>
    

    

  )

}

export default ThemeToggle
