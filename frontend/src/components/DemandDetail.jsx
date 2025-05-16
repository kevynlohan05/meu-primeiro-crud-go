import React from 'react'
import { useParams, useNavigate } from 'react-router-dom'

// Simulação do banco (ideal seria buscar de uma API real)
const mockData = [
  { id: 'DM-001', titulo: 'Erro ao acessar o sistema', status: 'Em andamento', data: '2025-04-23' },
  { id: 'DM-002', titulo: 'Solicitação de acesso ao servidor', status: 'Finalizado', data: '2025-04-22' },
  { id: 'DM-003', titulo: 'Troca de equipamento', status: 'Aguardando atendimento', data: '2025-04-21' }
]

const DemandDetail = () => {
  const { id } = useParams()
  const navigate = useNavigate()

  const demanda = mockData.find(d => d.id === id)

  if (!demanda) return <div className="p-10">Demanda não encontrada.</div>

  return (
    <div className="bg-slate-100 min-h-screen p-10">
      <button
        onClick={() => navigate(-1)}
        className="mb-4 px-4 py-2 bg-gray-300 hover:bg-gray-400 rounded"
      >
        Voltar
      </button>

      <div className="bg-white shadow-md p-8 rounded-md max-w-xl mx-auto">
        <h2 className="text-2xl font-bold mb-4">Detalhes da Demanda</h2>
        <p><strong>ID:</strong> {demanda.id}</p>
        <p><strong>Título:</strong> {demanda.titulo}</p>
        <p><strong>Status:</strong> {demanda.status}</p>
        <p><strong>Data:</strong> {demanda.data}</p>
        {/* Aqui você pode expandir com descrição, anexos, histórico, etc */}
      </div>
    </div>
  )
}

export default DemandDetail
