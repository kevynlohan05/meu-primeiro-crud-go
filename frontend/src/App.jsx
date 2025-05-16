import React from 'react'
import { BrowserRouter as Router, Routes, Route, Outlet } from 'react-router-dom'
import Login from './pages/Login'
import Home from './pages/Home'
import { ThemeProvider } from './contexts/ThemeContext'
import Desk from './pages/Desk'
import DemandForm from './pages/DemandForm'
import MyDemand from './pages/myDemand'
import SettingsPage from './pages/SettingsPage'
import Demand from './pages/Demand'
import { Layout } from './Layout/Layout'


const App = () => {
  return (
    <>
      <ThemeProvider>
      <Router>
        <Routes>
          <Route path="/login" element={<Login />} />
          <Route path="/" element={<Layout />} >
          <Route path="/home" index element={<Home />} />
          <Route path="/desk" element={<Desk />} />
          <Route path="/formulario" element={<DemandForm />} />
          <Route path="/demandas" element={<MyDemand />} />
          <Route path="/configuração" element={<SettingsPage />} />
          <Route path="/demand/:id" element={<Demand />} />
          </Route>
        </Routes>
      </Router>
      </ThemeProvider>
    </>
  )
}

export default App
