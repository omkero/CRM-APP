// lib/GetProducts.ts

import { RequestCookie } from "next/dist/compiled/@edge-runtime/cookies";
import { BASE_URL } from "../constant";

type Props = {
  token: string | any,
  page: number
}
export async function GetProducts(token: string | any , page = 1) {
  const path = BASE_URL + `/product/get_all_products?page=${page}`;

  try {
    const response = await fetch(path, {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${token}`,
      },
      cache: "no-store", // important for server-side fetching
    });

    if (!response.ok) {
        console.log(response.statusText);

    }
    const data = await response.json()
    
    return data; // <--- RETURN the data
  } catch (err) {
    console.error(err);
  }
}
