import React, { useState } from 'react'
import { useNavigate } from 'react-router-dom';

const UserManager = () => {
  const navigate = useNavigate();

  const [usuarios, setUsuarios] = useState([
    { id: 1, nome: "David flor", email: "david@email.com", bloqueado: false },
  ]);

  const toggleBloqueio = (id) => {
    setUsuarios(
      usuarios.map((u) => (u.id === id ? { ...u, bloqueado: !u.bloqueado } : u))
    );
  };

  const excluirUsuario = (id) => {
    setUsuarios(usuarios.filter((u) => u.id !== id));
  };

  const handleCadastro = (e) => {
    e.preventDefault();
    const form = new FormData(e.target);
    const novoUsuario = {
      id: usuarios.length + 1,
      nome: form.get("nome"),
      email: form.get("email"),
      bloqueado: false,
    };
    setUsuarios([...usuarios, novoUsuario]);
    e.target.reset();
  };
  


  return (
    <div className="overflow-x-auto">
            <table className="w-full mt-4 border-collapse bg-white dark:bg-slate-200 rounded-md">
              <thead>
                <tr>
                  <th className="border p-2">Nome</th>
                  <th className="border p-2">E-mail</th>
                  <th className="border p-2">Status</th>
                  <th className="border p-2">Ações</th>
                </tr>
              </thead>
              <tbody>
                {usuarios.map((user) => (
                  <tr key={user.id}>
                    <td className="border p-2">{user.nome}</td>
                    <td className="border p-2">{user.email}</td>
                    <td className="border p-2">
                      {user.bloqueado ? "Bloqueado" : "Ativo"}
                    </td>
                    <td className="border p-2 flex gap-2">
                      <button
                        className="text-white bg-yellow-500 px-2 py-1 rounded"
                        onClick={() => toggleBloqueio(user.id)}
                      >
                        {user.bloqueado ? "Desbloquear" : "Bloquear"}
                      </button>
                      <button
                        className="text-white bg-red-500 px-2 py-1 rounded"
                        onClick={() => excluirUsuario(user.id)}
                      >
                        Excluir
                      </button>
                    </td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
  )
}

export default UserManager