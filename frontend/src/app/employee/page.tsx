import Image from "next/image";
import { Sidebar } from "../components/sidebar/sidebar";
import { Navbar } from "../components/navbar/navbar";
import Products from "../components/products/products";
import { GetEmployees } from "./getProducts";
import { cookies } from "next/headers"; // 👈 NEW
import { sessionCookieName } from "../constant";
import Employees from "../components/employee/employees";


interface PageProps {
  params: any;
  searchParams: any;
}

export default async function Home({ params, searchParams }: PageProps) {
  const { page } = await searchParams;
  console.log(page)

  const cookieStore = cookies();
  const token: string | any = (await cookieStore).get(sessionCookieName)?.value;
  const data = await GetEmployees(token, page);

  return (
    <div className="h-full w-full flex flex-col">
      <Navbar />
      <div className="flex h-full">
        <Sidebar SelectedName="Employee" />
        <Employees data={data?.data} page={page} totalEmployees={data?.total_employees} />
      </div>
    </div>
  );
}
