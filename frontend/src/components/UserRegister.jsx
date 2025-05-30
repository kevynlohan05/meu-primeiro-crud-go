import { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import axios from 'axios';
import Select from 'react-select';

export function UserRegister() {
  const navigate = useNavigate();
  const [projectsList, setProjectsList] = useState([]);
  const [selectedProjects, setSelectedProjects] = useState([]);

  const [form, setForm] = useState({
    name: '',
    email: '',
    password: '',
    confirmPassword: '',
    phone: '',
    department: '',
    enterprise: '',
    role: '',
  });

  const [error, setError] = useState('');

  useEffect(() => {
    const fetchProjects = async () => {
      try {
        const token = localStorage.getItem('token');
        const response = await axios.get('http://localhost:8080/project/getAllProjects', {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        });

        const formattedProjects = response.data.map(project => ({
          value: project.name,
          label: project.name,
        }));

        setProjectsList(formattedProjects);
      } catch (err) {
        console.error('Erro ao buscar projetos:', err);
      }
    };

    fetchProjects();
  }, []);

  const handleChange = (e) => {
    const { name, value } = e.target;
    setForm(prev => ({ ...prev, [name]: value }));
  };

  const handleCadastro = async (e) => {
    e.preventDefault();

    if (form.password !== form.confirmPassword) {
      setError('As senhas não coincidem.');
      return;
    }

    try {
      const token = localStorage.getItem('token');

      const response = await axios.post(
        'http://localhost:8080/user/createUser',
        {
          name: form.name,
          email: form.email,
          password: form.password,
          phone: form.phone,
          department: form.department,
          projects: selectedProjects.map(p => p.value),
          enterprise: form.enterprise,
          role: form.role,
        },
        {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        }
      );

      console.log('Usuário cadastrado com sucesso:', response.data);
      navigate('/login');
    } catch (err) {
      console.error('Erro ao cadastrar usuário:', err);
      setError(err?.response?.data?.message || 'Erro desconhecido.');
    }
  };

  return (
    <div className="bg-white p-6 shadow-lg rounded-md w-full max-w-lg dark:bg-slate-200">
      <h2 className="text-2xl font-bold mb-4">Cadastro de Usuário</h2>

      {error && <p className="text-red-500 text-sm">{error}</p>}

      <form onSubmit={handleCadastro} className="space-y-4">
        <Input label="Nome" name="name" value={form.name} onChange={handleChange} />
        <Input label="E-mail" type="email" name="email" value={form.email} onChange={handleChange} />
        <Input label="Telefone (apenas números)" name="phone" value={form.phone} onChange={handleChange} />
        <Input label="Departamento" name="department" value={form.department} onChange={handleChange} />

        <div>
          <label className="block text-sm font-medium text-gray-700">Projetos (selecione um ou mais):</label>
          <Select
            isMulti
            name="projects"
            options={projectsList}
            className="basic-multi-select"
            classNamePrefix="select"
            onChange={setSelectedProjects}
            value={selectedProjects}
          />
        </div>

        <Input label="Empresa" name="enterprise" value={form.enterprise} onChange={handleChange} />

        <div>
          <label className="block text-sm font-medium text-gray-700">Perfil:</label>
          <select
            name="role"
            value={form.role}
            onChange={handleChange}
            required
            className="w-full p-2 border rounded-md"
          >
            <option value="">Selecione</option>
            <option value="admin">Administrador</option>
            <option value="user">Usuário</option>
          </select>
        </div>

        <Input label="Senha" type="password" name="password" value={form.password} onChange={handleChange} />
        <Input label="Confirmação de senha" type="password" name="confirmPassword" value={form.confirmPassword} onChange={handleChange} />

        <button type="submit" className="bg-green-600 text-white px-4 py-2 rounded-md">
          Cadastrar
        </button>
      </form>
    </div>
  );
}

function Input({ label, name, value, onChange, type = 'text', placeholder = '' }) {
  return (
    <div>
      <label className="block text-sm font-medium text-gray-700">{label}:</label>
      <input
        type={type}
        name={name}
        value={value}
        onChange={onChange}
        placeholder={placeholder}
        required
        className="w-full p-2 border rounded-md"
      />
    </div>
  );
}
