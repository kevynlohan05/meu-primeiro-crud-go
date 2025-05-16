import React from 'react'
import { Link } from 'react-router-dom'

const Linked = ({ href, icon, children, className = ''}) => {
  return (
    <Link
      to={href}
      className="text-gray-800 dark:text-gray-100 hover:text-green-600  dark:hover:text-green-400 border-green-400 font-medium px-3 py-2 rounded-lg hover:bg-gray-300 dark:hover:bg-slate-600 flex items-center gap-2"
    >
     {icon} {children}
    </Link>
  )
}

export default Linked