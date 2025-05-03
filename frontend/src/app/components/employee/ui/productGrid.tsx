import { Activity, Ellipsis, Send } from "lucide-react";
import React from "react";
import "../../../globals.css";
import { PorductType } from "@/app/products/types";
import Image from "next/image";
import { BASE_URL } from "@/app/constant";
import { EmployeeType } from "@/app/employee/types";
import EmployeeCard from "./employeeCard";
import NotFound from "../../ui/notFound";

type Props = {
  data: EmployeeType[];
};

export default function EmployeesGrid({ data }: Props) {
  const EmployeesList = () => {
    return (
      <>
        {data?.map((item: EmployeeType, index: number) => (
          <EmployeeCard key={index} data={item} />
        ))}
      </>
    );
  };
  return !data ? (
    <NotFound />
  ) : (
    <div className="h-full w-full grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4">
      <EmployeesList />
    </div>
  );
}
