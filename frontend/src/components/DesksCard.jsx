

const DesksCard = ({children, title, icon, href}) => {
  return (
    <a href={href} className="bg-white rounded-md text-black shadow-lg hover:shadow-2xl border-4 border-white hover:border-green-600 hover: p-6 w-full max-w-sm mx-auto dark: bg-slate-200 border-slate-300">
      <h3 className="text-xl font-semibold text-center flex items-center gap-2">{icon} {title}</h3>
      <hr className="mb-4 border-slate-300 dark:border-slate-500" />
      <p>{children}</p>
    </a>
  )
}

export default DesksCard
