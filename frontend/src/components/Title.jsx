import React from 'react'

const Title = ({children, className = '', type = ''}) => {
  return (
    <h1 
    className={`text-center text-4xl pt-3 font-bold text-gray-800 dark:text-slate-100 ${className}`}>
        {children}
    </h1>
  )
}

export default Title