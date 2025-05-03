"use client";
import { ChevronDown, CircleHelp } from "lucide-react";
import React, { useEffect, useState, useRef } from "react";
import Image from "next/image";
import { getSupportedBrowsers } from "next/dist/build/utils";

type Props = {};

function Profile({}: Props) {
  const [isHover, setIsHover] = useState(false);
  const modalRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    function Handler(event: any) {
      if (modalRef?.current && !modalRef?.current.contains(event.target)) {
        setIsHover(false);
      }
    }
    document.addEventListener("click", Handler);
    return () => {
      document.removeEventListener("click", Handler);
    };
  }, [modalRef]);
  return (
    <div className="flex items-center gap-9">
      {/* Help Center */}
      <div className="flex items-center gap-3 hover:bg-neutral-200 cursor-pointer p-3 rounded-xl transition-colors duration-300">
        <CircleHelp size={18} color="black" />
        <p className="text-xs">Help Center</p>
      </div>

      {/* Profile + Dropdown */}
      <div className="relative flex items-center gap-3">
        {/* Profile Section */}
        <div
          className="flex items-center gap-3 cursor-pointer"
          onMouseEnter={() => setIsHover(true)}
        >
          <Image
            alt="img"
            height={40}
            width={40}
            src="https://images.unsplash.com/photo-1502823403499-6ccfcf4fb453?&w=256&h=256&q=70&crop=focalpoint&fp-x=0.5&fp-y=0.3&fp-z=1&fit=crop"
            className="rounded-full"
          />
          <p className="text-sm font-normal">John Doe .F</p>
          <ChevronDown size={19} color="black" />
        </div>

        {/* Dropdown Modal - always rendered */}
        <div
          ref={modalRef}
          className={`absolute top-0 z-50 right-0 mt-16 bg-neutral-200 h-60 w-60 rounded-sm transition-all duration-300 ease-in-out ${
            isHover
              ? "opacity-100 translate-y-0"
              : "opacity-0 -translate-y-2 invisible"
          }`}
          onMouseLeave={() => setIsHover(false)}
        >
          <p className="p-4">Smooth Modal Content</p>
        </div>
      </div>
    </div>
  );
}
export default Profile;
