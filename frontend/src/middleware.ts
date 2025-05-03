import { RequestCookie } from "next/dist/compiled/@edge-runtime/cookies";
import { NextResponse } from "next/server";
import type { NextRequest } from "next/server";
import jwt from "jsonwebtoken";
import * as jose from "jose";
import { url } from "inspector";
import { sessionCookieName } from "./app/constant";

const secreteKey = process.env.SIGN_IN_PRIVATE_KEY;

// This function can be marked `async` if using `await` inside
export async function middleware(request: NextRequest) {
  const tokenName: string = sessionCookieName
  const token: RequestCookie | undefined = request.cookies?.get(tokenName);
  const tokenString: any = token?.value?.toString();

  if (!token && !request.url.includes("/auth/signin")) {
    return NextResponse.redirect(new URL("/auth/signin", request.url));
  }

  if (token && !request.url.includes("/auth/signin")) {
    try {
      const result = await jose.jwtVerify(
        tokenString,
        new TextEncoder().encode(secreteKey),
        {
          algorithms: ["HS256"], // <-- explicitly say HS256!!
        }
      );
      return;
    } catch (err: any) {
      console.log("error verifying jwt token:", err.message);
      return NextResponse.redirect(new URL("/auth/signin", request.url));
    }
  }

  // Token exists → allow to continue
  return NextResponse.next();
}

// See "Matching Paths" below to learn more
export const config = {
  matcher: ["/", "/auth/signin"],
};
