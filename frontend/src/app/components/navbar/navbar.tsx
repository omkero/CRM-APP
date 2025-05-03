"use client"
import React, { useState } from "react";
import Logo from "../sidebar/ui/logo";
import { useContext } from "react";
import { ConstantsContext } from "@/app/providers/constantsProvider";
import Search from "./ui/search";
import Profile from "./ui/profile";
type Props = {};

export const Navbar = (props: Props) => {
  const [isHover, setIsHover] = useState<boolean>(false);
  const constatns = useContext(ConstantsContext);
  return (
    <div 
    className="w-full  h-16 flex bg-white border-b-1  z-50 fixed border-b-neutral-300"
    >
      <div
        style={{
          minWidth:
            constatns?.sideWidth > constatns?.minSideWidth
              ? constatns?.sideWidth
              : constatns?.minSideWidth,
        }}
        className="flex flex-col gap-3 "
      >
        <Logo LogoTitle="Venture CRM" />
      </div>
      <div className="w-full p-5 px-9 flex items-center justify-between">
        <Search />
        <Profile  />
      </div>
    </div>
  );
};
