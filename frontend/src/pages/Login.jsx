import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import api from '../services/api';
import ThemeToggle from '../components/ThemeToggle';

const Login = () => {
  const navigate = useNavigate();
  const [email, setEmail] = useState('');
  const [senha, setSenha] = useState('');
  const [error, setError] = useState('');

  const handleLogin = async (e) => {
    e.preventDefault();

    try {
      const response = await api.post('/user/login', {
        email,
        password: senha,
      });

      const token = response.headers['authorization'];

      if (token && response.data) {
        // Armazenar token e dados do usuário no localStorage
        localStorage.setItem('token', token);
        localStorage.setItem('user', JSON.stringify(response.data));

        // Navegar para o formulário após login
        navigate('/formulario');
      } else {
        setError('Token não recebido ou dados inválidos.');
      }

    } catch (err) {
      console.error('Erro ao fazer login:', err);
      setError('Credenciais inválidas. Tente novamente.');
    }
  };

  return (
    <div className="min-h-screen flex items-center justify-center bg-slate-300 dark:bg-gray-900">
      <ThemeToggle />
      <div className="bg-white p-8 rounded-md shadow-lg w-full max-w-sm space-y-6 dark:bg-slate-200">
        <div className="text-center">
          <img src="/img/Logo.png" alt="Logotipo da empresa" className="mx-auto w-24 h-24 object-contain" />
          <h1 className="text-2xl font-bold mt-4 text-gray-900">Fazer Login</h1>
        </div>

        {error && <p className="text-red-500 text-center">{error}</p>}

        <form className="space-y-4" onSubmit={handleLogin}>
          <div>
            <label htmlFor="email" className="block text-sm font-medium text-gray-700">E-mail</label>
            <input
              id="email"
              type="email"
              placeholder="Digite seu email"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
              className="mt-1 w-full px-4 py-2 border rounded-md focus:outline-none focus:ring-2 focus:ring-green-500"
              required
            />
          </div>

          <div>
            <label htmlFor="password" className="block text-sm font-medium text-gray-700">Senha</label>
            <input
              id="password"
              type="password"
              placeholder="Digite sua senha"
              value={senha}
              onChange={(e) => setSenha(e.target.value)}
              className="mt-1 w-full px-4 py-2 border rounded-md focus:outline-none focus:ring-2 focus:ring-green-500"
              required
            />
          </div>

          <button
            type="submit"
            className="w-full bg-green-600 text-white font-bold py-2 rounded-md hover:bg-green-500 transition"
          >
            Entrar
          </button>
        </form>
      </div>
    </div>
  );
};

export default Login;
