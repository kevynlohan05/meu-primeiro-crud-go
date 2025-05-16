import React from "react";
import DesksCard from "../components/DesksCard";
import { MdOutlineWorkHistory } from "react-icons/md";
import { MdOutlineSupportAgent } from "react-icons/md";
import { Outlet, useNavigate } from "react-router-dom";


const Desk = () => {
  const Navigate = useNavigate()

  return (
    <div className="bg-slate-200 dark:bg-gray-900 min-h-screen">
    
      <div>
      <div className="">
        <div className="flex justify-center items-start min-h-screen py-36">
          <div className="flex flex-row gap-4 flex-wrap justify-center">
            <DesksCard href={'/formulario'} title={'Nova chamada'} icon={<MdOutlineSupportAgent />}>
              Solicitar abertura de um chamado com a equipe de suporte do FDS para atendimento.
            </DesksCard>
            <DesksCard href={'/demandas'} title={'Minhas Chamadas'} icon={<MdOutlineWorkHistory />}>
            Acompanhe o status de suas solicitações, sejam elas em aberto ou já finalizadas pelo atendimento.
            </DesksCard>
          </div>
         
        </div>
      </div>
      </div>
      
    </div>
  );
};

export default Desk;
