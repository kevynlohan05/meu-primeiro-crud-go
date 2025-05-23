import React, { useState, useEffect } from 'react';
import axios from 'axios';

const DemandForm = () => {
  const [user, setUser] = useState(null);
  const [projects, setProjects] = useState([]);
  const [form, setForm] = useState({
    title: '',
    project: '',
    description: '',
    request_type: '',
    priority: '',
    attachment_urls: [],
  });
  const [message, setMessage] = useState('');
  const [error, setError] = useState('');

  useEffect(() => {
    try {
      const storedUser = JSON.parse(localStorage.getItem('user'));
      if (storedUser) {
        setUser(storedUser);
        setProjects(storedUser.projects || []);
      } else {
        setError('Erro ao carregar informações do usuário.');
      }
    } catch (err) {
      setError('Erro ao carregar informações do usuário.');
    }
  }, []);

  const handleChange = (e) => {
    setForm({
      ...form,
      [e.target.name]: e.target.value,
    });
  };

  const handleFileChange = (e) => {
    const files = Array.from(e.target.files);
    setForm({
      ...form,
      attachment_urls: [...form.attachment_urls, ...files],
    });
  };

  const handleRemoveFile = (index) => {
    const updatedFiles = [...form.attachment_urls];
    updatedFiles.splice(index, 1);
    setForm({
      ...form,
      attachment_urls: updatedFiles,
    });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    const token = localStorage.getItem('token');
    if (!token || !user) {
      setError('Usuário não autenticado.');
      return;
    }

    const formData = new FormData();
    formData.append('title', form.title);
    formData.append('project', form.project);
    formData.append('description', form.description);
    formData.append('request_type', form.request_type);
    formData.append('priority', form.priority);

    form.attachment_urls.forEach((file) => {
      formData.append('attachment_urls', file);
    });

    try {
      await axios.post('http://localhost:8080/ticket/createTicket', formData, {
        headers: {
          Authorization: token,
          'Content-Type': 'multipart/form-data',
        },
      });

      setMessage('Demanda enviada com sucesso!');
      setError('');
      setForm({
        title: '',
        project: '',
        description: '',
        request_type: '',
        priority: '',
        attachment_urls: [],
      });
    } catch (err) {
      setError('Erro ao enviar demanda. Verifique os dados e tente novamente.');
      setMessage('');
    }
  };

  return (
    <div className="p-6 max-w-2xl mx-auto bg-gray-100 rounded-xl shadow-md">
      <h2 className="text-2xl font-bold mb-4 text-center">Solicitação de Demanda</h2>
      <p className="text-gray-600 mb-4 text-center">
        Preencha o formulário abaixo para enviar sua solicitação.
      </p>

      {error && <p className="text-red-500 text-center mb-4">{error}</p>}
      {message && <p className="text-green-600 text-center mb-4">{message}</p>}

      <form onSubmit={handleSubmit} className="space-y-4">
        <div>
          <label className="block font-medium">Título da solicitação*</label>
          <input
            type="text"
            name="title"
            value={form.title}
            onChange={handleChange}
            required
            className="w-full border border-gray-300 rounded px-3 py-2"
          />
        </div>

        <div>
          <label className="block font-medium">Projeto*</label>
          <select
            name="project"
            value={form.project}
            onChange={handleChange}
            required
            className="w-full border border-gray-300 rounded px-3 py-2"
          >
            <option value="">Selecione um projeto</option>
            {projects.map((proj, index) => (
              <option key={index} value={proj}>
                {proj}
              </option>
            ))}
          </select>
        </div>

        <div>
          <label className="block font-medium">Detalhes da solicitação*</label>
          <textarea
            name="description"
            value={form.description}
            onChange={handleChange}
            required
            className="w-full border border-gray-300 rounded px-3 py-2"
          />
        </div>

        <div>
          <label className="block font-medium">Tipo de Solicitação*</label>
          <select
            name="request_type"
            value={form.request_type}
            onChange={handleChange}
            required
            className="w-full border border-gray-300 rounded px-3 py-2"
          >
            <option value="">Selecione</option>
            <option value="Bugs">Bugs</option>
            <option value="Dúvidas">Dúvidas</option>
            <option value="Nova implementação">Nova implementação</option>
            <option value="Redefinir a senha">Redefinir a senha</option>
            <option value="Solicitação de acesso">Solicitação de acesso</option>
            <option value="Verificação de apostas">Verificação de apostas</option>
            <option value="Outros">Outros</option>
          </select>
        </div>

        <div>
          <label className="block font-medium">Prioridade*</label>
          <select
            name="priority"
            value={form.priority}
            onChange={handleChange}
            required
            className="w-full border border-gray-300 rounded px-3 py-2"
          >
            <option value="">Selecione</option>
            <option value="baixa">Baixa</option>
            <option value="media">Média</option>
            <option value="alta">Alta</option>
          </select>
        </div>

        <div>
          <label className="block font-medium mb-1">Anexos (imagens)*</label>
          <input
            type="file"
            name="attachment_urls"
            accept="image/*"
            multiple
            onChange={handleFileChange}
            className="w-full border border-gray-300 rounded px-3 py-2"
          />
          {form.attachment_urls.length > 0 && (
            <ul className="mt-2 space-y-1">
              {form.attachment_urls.map((file, index) => (
                <li
                  key={index}
                  className="flex items-center justify-between text-sm text-gray-700 bg-gray-200 px-2 py-1 rounded"
                >
                  <span className="truncate max-w-xs">{file.name}</span>
                  <button
                    type="button"
                    onClick={() => handleRemoveFile(index)}
                    className="text-red-500 hover:underline ml-4"
                  >
                    Remover
                  </button>
                </li>
              ))}
            </ul>
          )}
        </div>

        <div className="text-center">
          <button
            type="submit"
            className="bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700"
          >
            Enviar Demanda
          </button>
        </div>
      </form>
    </div>
  );
};

export default DemandForm;
