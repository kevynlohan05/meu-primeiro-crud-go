import React from 'react'

const Paragraph = ({children, classname = ''}) => {
  return (
    <p className= {`text-center text-gray-900 dark:text-slate-100 ${classname}`}>
        {children}
    </p>
  )
}

export default Paragraph