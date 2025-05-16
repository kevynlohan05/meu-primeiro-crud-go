import React, { useEffect, useState } from 'react'
import { useNavigate } from 'react-router-dom'

// Simulação do fetch da API do Asana
const mockData = [
  {
    id: 'DM-001',
    titulo: 'Erro ao acessar o sistema',
    status: 'Em andamento',
    data: '2025-04-23'
  },
  {
    id: 'DM-002',
    titulo: 'Solicitação de acesso ao servidor',
    status: 'Finalizado',
    data: '2025-04-22'
  },
  {
    id: 'DM-003',
    titulo: 'Troca de equipamento',
    status: 'Aguardando atendimento',
    data: '2025-04-21'
  }
]

const MyDemand = () => {
  const [demandas, setDemandas] = useState([])
  const navigate = useNavigate()

  

  useEffect(() => {
    // Em produção: fetch de uma API real
    setTimeout(() => {
      setDemandas(mockData)
    }, 1000)
  }, [])

  return (
    <>
      <div className="bg-slate-200  dark:bg-gray-900 min-h-screen pl-72">
        <div className="px-6 py-10">

          <div className='bg-white p-8 rounded-md shadow-lg mx-auto space-y-6 mt-8 dark:bg-slate-200'>
          <h1 className="text-3xl font-bold text-gray-900 dark:mb-6">
            Meus Chamados
          </h1>

          <div className="overflow-x-auto">
            <table className="min-w-full bg-white dark:bg-slate-200 rounded-md shadow-md">
              <thead className="bg-green-600 text-white">
                <tr>
                  <th className="py-3 px-4 text-left">ID</th>
                  <th className="py-3 px-4 text-left">Título da Demanda</th>
                  <th className="py-3 px-4 text-left">Status</th>
                  <th className="py-3 px-4 text-left">Data</th>
                </tr>
              </thead>
              <tbody>
                {demandas.map((demanda, index) => (
                  <tr
                  key={index}
                  className="border-t border-slate-300 hover:bg-slate-100 dark:hover:bg-slate-300 cursor-pointer"
                  onClick={() => navigate(`/demand/${demanda.id}`)}
                >
                    <td className="py-3 px-4">{demanda.id}</td>
                    <td className="py-3 px-4">{demanda.titulo}</td>
                    <td className="py-3 px-4">{demanda.status}</td>
                    <td className="py-3 px-4">{demanda.data}</td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
          </div>
          
          
        </div>
      </div>
    </>
  )
}

export default MyDemand
