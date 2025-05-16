import { useState } from 'react';
import { useNavigate } from 'react-router-dom';


export function UserRegister() {
  const navigate = useNavigate();

  const [form, setForm] = useState({
    nome: '',
    email: '',
    cargo: '',
    perfil: '',
    senha: '',
    confirmacaoSenha: ''
  });

  const handleChange = (e) => {
    const { name, value } = e.target;
    setForm((prev) => ({ ...prev, [name]: value }));
  };

  const handleCadastro = (e) => {
    e.preventDefault();
    console.log('Dados do cadastro:', form);
    // Validação e envio à API aqui
  };


  return (
    

      <div className="bg-white p-6 shadow-lg rounded-md w-full max-w-lg dark:bg-slate-200">
        <h2 className="text-2xl font-bold mb-4">Cadastro de Usuário</h2>

        <form onSubmit={handleCadastro} className="space-y-4">
          <div>
            <label className="block text-sm font-medium text-gray-700">Nome:</label>
            <input
              type="text"
              name="nome"
              value={form.nome}
              onChange={handleChange}
              required
              className="w-full p-2 border rounded-md"
            />
          </div>

          <div>
            <label className="block text-sm font-medium text-gray-700">E-mail:</label>
            <input
              type="email"
              name="email"
              value={form.email}
              onChange={handleChange}
              required
              className="w-full p-2 border rounded-md"
            />
          </div>

          <div>
            <label className="block text-sm font-medium text-gray-700">Cargo:</label>
            <input
              type="text"
              name="cargo"
              value={form.cargo}
              onChange={handleChange}
              required
              className="mt-1 w-full px-4 py-2 border rounded-md focus:outline-none focus:ring-2 focus:ring-green-500"
            />
          </div>

          <div>
            <label className="block text-sm font-medium text-gray-700">Perfil:</label>
            <select
              name="perfil"
              value={form.perfil}
              onChange={handleChange}
              required
              className="mt-1 w-full px-4 py-2 border rounded-md bg-white focus:outline-none focus:ring-2 focus:ring-green-500"
            >
              <option value="">Selecione</option>
              <option value="admin">Administrador</option>
              <option value="user">Usuário</option>
            </select>
          </div>

           
    
          <div>
            <label className="block text-sm font-medium text-gray-700">Senha:</label>
            <input
              type="password"
              name="senha"
              value={form.senha}
              onChange={handleChange}
              required
              className="w-full p-2 border rounded-md"
            />
          </div>

          <div>
            <label className="block text-sm font-medium text-gray-700">Confirmação de senha:</label>
            <input
              type="password"
              name="confirmacaoSenha"
              value={form.confirmacaoSenha}
              onChange={handleChange}
              required
              className="w-full p-2 border rounded-md"
            />
          </div>

          <button
            type="submit"
            className="bg-green-600 text-white px-4 py-2 rounded-md"
          >
            Cadastrar
          </button>
        </form>
      </div>

  );
}
