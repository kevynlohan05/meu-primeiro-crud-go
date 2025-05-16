import React, { useState } from "react";
import { UserRegister } from "../components/UserRegister";
import UserManager from "../components/UserManager";

const SettingsPage = () => {
  const [activeTab, setActiveTab] = useState("cadastro");

  return (
    <>
      <h1 className="text-3xl font-bold mb-6">Configurações</h1>

      <div className="flex gap-4 mb-6">
        <button
          className={`px-4 py-2 rounded-md ${
            activeTab === "cadastro"
              ? "bg-green-600 text-white"
              : "bg-gray-300 dark:bg-slate-700"
          }`}
          onClick={() => setActiveTab("cadastro")}
        >
          Cadastrar Usuário
        </button>
        <button
          className={`px-4 py-2 rounded-md ${
            activeTab === "gerenciar"
              ? "bg-green-600 text-white"
              : "bg-gray-300 dark:bg-slate-700"
          }`}
          onClick={() => setActiveTab("gerenciar")}
        >
          Gerenciar Usuários
        </button>
      </div>

      {/* Abas */}
      {activeTab === "cadastro" && <UserRegister />}

      {activeTab === "gerenciar" && <UserManager />}
    </>
  );
};

export default SettingsPage;
