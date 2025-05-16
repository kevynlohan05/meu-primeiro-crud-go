import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import api from '../services/api'; // usa instância do axios

const DemandForm = () => {
  const navigate = useNavigate();

  const [formData, setFormData] = useState({
    title: '',
    request_user: '',
    sector: '',
    description: '',
    request_type: '',
    priority: '',
    attachment_url: ''
  });

  const [file, setFile] = useState(null); // opcional, se quiser enviar o anexo depois
  const [error, setError] = useState('');
  const [success, setSuccess] = useState('');

  const handleChange = (e) => {
    const { name, value, files } = e.target;
    if (files) {
      setFile(files[0]); // Aqui futuramente você poderá enviar o arquivo.
    } else {
      setFormData((prev) => ({
        ...prev,
        [name]: value
      }));
    }
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    try {
      const token = localStorage.getItem('token');

      await api.post('/ticket/createTicket', formData, {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });

      const payload = {
        ...formData,
        attachment_url: '' // ou envie o caminho do arquivo quando for usar upload
      };

      console.log('Token enviado:', token);
      console.log('Config headers:', config.headers);

      await api.post('/ticket/createTicket', payload, config);

      setSuccess('Solicitação enviada com sucesso!');
      setError('');
      setTimeout(() => navigate('/demandas'), 1500);
    } catch (err) {
      console.error(err);
      setError('Erro ao enviar a solicitação. Tente novamente.');
      setSuccess('');
    }
  };

  return (
    <div className="bg-slate-300 dark:bg-gray-900 min-h-screen flex items-center justify-center px-4">
      <div className="bg-white h-[calc(100vh-40px)] overflow-auto p-8 shadow-lg w-full max-w-2xl space-y-6 dark:bg-slate-200">
        <h1 className="text-2xl font-bold text-center text-gray-900">Solicitação de Demanda</h1>
        <p className="text-gray-600 text-center">Preencha o formulário abaixo para enviar sua solicitação.</p>

        {success && <p className="text-green-600 text-center">{success}</p>}
        {error && <p className="text-red-600 text-center">{error}</p>}

        <form className="space-y-8" onSubmit={handleSubmit}>
          <div>
            <label className="block text-sm font-medium text-gray-700">Título da solicitação<span className='text-red-500'>*</span></label>
            <input
              type="text"
              name="title"
              required
              onChange={handleChange}
              className="mt-1 w-full px-4 py-2 border rounded-md focus:outline-none focus:ring-2 focus:ring-green-500"
              placeholder="Digite a sua resposta"
            />
          </div>

          <div>
            <label className="block text-sm font-medium text-gray-700">Nome do solicitante<span className='text-red-500'>*</span></label>
            <input
              type="text"
              name="request_user"
              required
              onChange={handleChange}
              className="mt-1 w-full px-4 py-2 border rounded-md focus:outline-none focus:ring-2 focus:ring-green-500"
              placeholder="Digite o seu nome"
            />
          </div>

          <div>
            <label className="block text-sm font-medium text-gray-700">Setor do solicitante<span className='text-red-500'>*</span></label>
            <select
              name="sector"
              required
              onChange={handleChange}
              className="mt-1 w-full px-4 py-2 border rounded-md bg-white focus:outline-none focus:ring-2 focus:ring-green-500"
            >
              <option value="">Selecione</option>
              <option value="Atendimento">Atendimento</option>
              <option value="Segurança">Segurança</option>
              <option value="Gestão do Afiliado">Gestão do Afiliado</option>
              <option value="Marketing">Marketing</option>
              <option value="Outro">Outro</option>
            </select>
          </div>

          <div>
            <label className="block text-sm font-medium text-gray-700">Detalhes da solicitação<span className='text-red-500'>*</span></label>
            <textarea
              name="description"
              required
              rows={4}
              onChange={handleChange}
              className="mt-1 w-full px-4 py-2 border rounded-md focus:outline-none focus:ring-2 focus:ring-green-500"
              placeholder="Digite os detalhes da sua solicitação"
            />
          </div>

          <div>
            <label className="block text-sm font-medium text-gray-700">Tipo de Solicitação<span className='text-red-500'>*</span></label>
            <select
              name="request_type"
              required
              onChange={handleChange}
              className="mt-1 w-full px-4 py-2 border rounded-md bg-white focus:outline-none focus:ring-2 focus:ring-green-500"
            >
              <option value="">Selecione</option>
              <option value="Bugs">Bugs</option>
              <option value="Dúvidas">Dúvidas</option>
              <option value="Nova implementação">Nova implementação</option>
              <option value="Verificação de apostas">Verificação de apostas</option>
              <option value="Solicitação de acesso">Solicitação de acesso</option>
              <option value="Outros">Outros</option>
            </select>
          </div>

          <div>
            <label className="block text-sm font-medium text-gray-700">Prioridade<span className='text-red-500'>*</span></label>
            <select
              name="priority"
              required
              onChange={handleChange}
              className="mt-1 w-full px-4 py-2 border rounded-md bg-white focus:outline-none focus:ring-2 focus:ring-green-500"
            >
              <option value="">Selecione</option>
              <option value="Baixa">Baixa</option>
              <option value="Média">Média</option>
              <option value="Alta">Alta</option>
            </select>
          </div>

          <div>
            <label className="block text-sm font-medium text-gray-700">Anexos</label>
            <input
              type="file"
              name="anexos"
              onChange={handleChange}
              className="mt-1 block w-full text-sm text-gray-500 file:mr-4 file:py-2 file:px-4 file:rounded file:border-0 file:text-sm file:font-semibold file:bg-green-50 file:text-green-700 hover:file:bg-green-100"
            />
            <p className="text-xs text-gray-500 mt-1">Caso haja necessidade de um print, envie-nos.</p>
          </div>

          <button
            type="submit"
            className="w-full bg-green-600 text-white font-bold py-2 rounded-md hover:bg-green-500 transition"
          >
            Enviar Solicitação
          </button>
        </form>
      </div>
    </div>
  );
};

export default DemandForm;
