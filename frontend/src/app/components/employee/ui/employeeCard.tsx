import { EmployeeType } from "@/app/employee/types";
import { ParseDate } from "@/lib/parseData";
import { Activity, Ellipsis, Send, SquareActivity } from "lucide-react";
import React from "react";

type Props = {
  data: EmployeeType;
};

function EmployeeCard({ data }: Props) {
  return (
    <div className="flex flex-col outline outline-neutral-100  gap-4 bg-neutral-100 hover:outline-neutral-300  p-4 w-full">
      <div className="flex items-center gap-4 w-full justify-between">
        <div className="flex items-center gap-4">
          <span className=" h-16 w-16 rounded-full bg-neutral-300"></span>
          <div className="flex flex-col gap-1">
            <h1 className="text-lg">{data.employee_full_name}</h1>
            <p className="text-sm text-neutral-500">{data.employee_position}</p>
          </div>
        </div>
        <div className="h-full">
          <button className="hover:bg-white p-2 outline outline-transparent rounded-full hover:outline-neutral-300 cursor-pointer ">
            <Ellipsis size={19} color="black" />
          </button>
        </div>
      </div>
      <div className="flex items-center gap-2">
        {data?.employee_role?.map?.((item: string, index: number) => (
          <span
            key={index}
            className="bg-violet-200 shadow-3xl text-violet-700 py-1 px-2 text-xs rounded-2xl"
          >
            {item}
          </span>
        ))}
      </div>
      <div className="flex flex-col gap-4">
        <div className="flex items-center gap-3">
          <h1>Assinged at: </h1>
          <span className="text-sm text-neutral-500">
            {ParseDate(data.created_at)}
          </span>
        </div>
        <div className="flex w-full gap-2 items-center justify-between">
          <button className=" bg-white hover:bg-neutral-100 cursor-pointer  flex items-center w-[50%] justify-center gap-2 shadow-2xs p-[6px] outline outline-neutral-300 rounded-3xl">
            <Send size={16} />
            <span className="text-sm">Contact</span>
          </button>
          <button className=" bg-white hover:bg-neutral-100 cursor-pointer flex items-center w-[50%] justify-center gap-2 shadow-2xs p-[6px] outline outline-neutral-300 rounded-3xl">
            <Activity size={16} />
            <span className="text-sm">Assign Task</span>
          </button>
        </div>
      </div>
    </div>
  );
}

export default EmployeeCard;
