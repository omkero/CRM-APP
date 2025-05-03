"use server";
import { headers } from "next/headers";
import { BASE_URL, sessionCookieName } from "../constant";
import { CreateProductResponseType, CreateProductType } from "./types";
import { cookies } from "next/headers"; // 👈 NEW
import { ReadonlyRequestCookies } from "next/dist/server/web/spec-extension/adapters/request-cookies";

export async function CreateProductAction(payload: CreateProductType): Promise<CreateProductResponseType | undefined > {
  try {
    const cookieStore: Promise<ReadonlyRequestCookies> = cookies();
    const token: string | any = (await cookieStore).get(
      sessionCookieName
    )?.value;
    const path = BASE_URL + "/product/create_product";

    const form = new FormData();
    form.append("product_title", payload.product_title);
    form.append("product_description", payload.product_description);
    form.append("product_price", payload.product_price);
    form.append("product_cover", payload.product_cover);

    const response = await fetch(path, {
      method: "POST",
      headers: {
        Authorization: `Bearer ${token}`,
      },
      body: form,
    });
    const result = await response.json();
    const replyPayload: CreateProductResponseType = {
        status: response.status,
        message: result.message
    }
    console.log(response)
    return replyPayload
  } catch (err) {
    console.log(err);
  }
}
