"use client";
import { Button, Checkbox, Flex, Text } from "@radix-ui/themes";
import React, { FormEvent, useState } from "react";
import Image from "next/image";
import { Eye, EyeOff } from "lucide-react";
import Link from "next/link";
import { useRouter } from "next/navigation";
import { BASE_URL } from "@/app/constant";

type Props = {};

function Login({}: Props) {
  const [isPasswordHidden, setIsPasswordHidden] = useState<boolean>(true);
  const [email, setEmail] = useState<string>("");
  const [password, setPassword] = useState<string>("");

  const router = useRouter();
  async function LoginAction() {
    const path = BASE_URL + "/employee/signin";
    try {
      const response = await fetch(path, {
        method: "POST",
        credentials: "include", // <-- VERY IMPORTANT
        headers: {
          "Content-type": "application/json",
        },
        body: JSON.stringify({
          employee_email_address: email,
          employee_password: password,
        }),
      });
      const toJson = await response.json();
      if (response.status == 200) {
        router.push("/");
      }
    } catch (err: any) {
      console.log(err);
    }
  }

  return (
    <div className="flex items-center w-full justify-center rounded-md px-4 sm:px-10 lg:px-10">
      <div className="h-full w-full justify-center  flex ">
        <div className="flex w-full md:w-[580px] bg-white flex-col gap-5 p-5 sm:p-10 py-10 lg:py-16">
          <div className="flex flex-col gap-3">
            <h1 className="text-lg lg:text-2xl font-bold">Sign In</h1>
            <p className="text-xs lg:text-sm text-teal-600">
              Welcome & Join us by creating a free account !
            </p>
          </div>
          <div className="flex flex-col gap-6">
            <div className="flex flex-col gap-2">
              <label className="text-teal-800 text-sm lg:text-base">
                Username
              </label>
              <input
                type="text"
                className="w-full py-[10px] px-2 text-sm outline outline-neutral-400 rounded-sm hover:outline-2 hover:outline-blue-500"
                placeholder="Enter Email"
                onChange={(e: any) => {
                  setEmail(e.target?.value);
                }}
              />
            </div>
            <div className="flex flex-col gap-2">
              <div className="w-full flex items-center justify-between">
                <label className="text-teal-800 text-sm lg:text-base">
                  Password
                </label>
                <a className="text-xs lg:text-sm text-violet-700 cursor-pointer">
                  forgot password ?
                </a>
              </div>
              <div className="outline outline-neutral-400  flex items-center gap-2 pr-3 rounded-sm hover:outline-2 hover:outline-blue-500">
                <input
                  type={isPasswordHidden ? "password" : "text"}
                  className="w-full py-[12px] pl-2 text-xs outline-none placeholder:text-sm "
                  placeholder="Enter Password"
                  onChange={(e: any) => {
                    setPassword(e.target?.value);
                  }}
                />
                <button
                  className="cursor-pointer"
                  onClick={() => {
                    setIsPasswordHidden(!isPasswordHidden);
                  }}
                >
                  {isPasswordHidden ? <EyeOff /> : <Eye />}
                </button>
              </div>
              <Text as="label" size="2">
                <Flex gap="2">
                  <Checkbox defaultChecked />
                  Remember password ?
                </Flex>
              </Text>
            </div>
          </div>
          <div className="w-full">
            <Button
              style={{ width: "100%", padding: 20 }}
              className="w-full"
              color="indigo"
              variant="solid"
              onClick={async () => {
                await LoginAction();
              }}
            >
              Sign In
            </Button>
          </div>
        </div>
        <div className="h-auto w-[600px] bg-gray-700 hidden md:flex items-center justify-center">
          <div className="flex flex-col p-7 gap-2">
            <Image
              alt="image"
              height={60}
              width={60}
              src={require("@/app/assets/logo/v_log.jpg")}
            />
            <h1 className="text-white text-base lg:text-2xl font-bold">
              sign in
            </h1>
            <p className="text-gray-400 text-xs lg:text-sm">
              Lorem ipsum dolor sit amet, consectetur adipisicing elit. Ipsa
              eligendi expedita aliquam quaerat nulla voluptas facilis. Porro
              rem voluptates possimus, ad, autem quae culpa architecto, quam
              labore blanditiis at ratione.
            </p>
          </div>
        </div>
      </div>
    </div>
  );
}

export default Login;
