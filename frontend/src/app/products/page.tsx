import Image from "next/image";
import { Sidebar } from "../components/sidebar/sidebar";
import { Navbar } from "../components/navbar/navbar";
import Products from "../components/products/products";
import { GetProducts } from "./getProducts";
import { cookies } from "next/headers"; // 👈 NEW
import { sessionCookieName } from "../constant";

interface PageProps {
  params: any;
  searchParams: any;
}

export default async function Home({ params, searchParams }: PageProps) {
  const { page } = await searchParams;

  const cookieStore = cookies();
  const token: string | any = (await cookieStore).get(sessionCookieName)?.value;
  const data = await GetProducts(token, page);

  return (
    <div className="h-full w-full flex flex-col">
      <Navbar />
      <div className="flex h-full">
        <Sidebar SelectedName="Products" />
        <Products data={data?.data} page={page} totalProducts={data?.total_products} />
      </div>
    </div>
  );
}
