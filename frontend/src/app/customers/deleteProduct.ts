"use server";
import { headers } from "next/headers";
import { BASE_URL, sessionCookieName } from "../constant";
import { CreateProductResponseType, CreateProductType } from "./types";
import { cookies } from "next/headers";
import { ReadonlyRequestCookies } from "next/dist/server/web/spec-extension/adapters/request-cookies";

export async function DeleteProductAction(product_id: number): Promise<CreateProductResponseType | undefined > {
  try {
    const cookieStore: Promise<ReadonlyRequestCookies> = cookies();
    const token: string | any = (await cookieStore).get(
      sessionCookieName
    )?.value;
    const path = BASE_URL + "/product/delete_product";

    const response = await fetch(path, {
      method: "POST",
      headers: {
        "Content-type": "application/json",
        Authorization: `Bearer ${token}`,
      },
      body: JSON.stringify({
        product_id: product_id
      }),
    });
    const result = await response.json();
    const replyPayload: CreateProductResponseType = {
        status: response.status,
        message: result.message
    }
    return replyPayload
  } catch (err) {
    console.log(err);
  }
}
