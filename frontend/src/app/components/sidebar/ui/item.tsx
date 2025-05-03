"use client";
import { LucideProps } from "lucide-react";
import React from "react";
import { useState } from "react";
import { ChevronUp, ChevronDown } from "lucide-react";
import Link from "next/link";

type ItemProps = {
  ItemTitle: string;
  Href?: any;
  IsSelected?: boolean;
  children?: React.ReactNode;
};
type GroupItemProps = {
  ItemTitle: string;
  IsSelected?: boolean;
  LucideIcon: React.ForwardRefExoticComponent<
    Omit<LucideProps, "ref"> & React.RefAttributes<SVGSVGElement>
  >;
  children: React.ReactNode;
};

export function SideBarItem({
  ItemTitle,
  Href,
  IsSelected,
  children,
}: ItemProps) {
  return (
    <Link
      href={Href}
      className={`flex items-center gap-4 w-full cursor-pointer ${
        IsSelected ? "bg-neutral-300" : ""
      } hover:bg-neutral-300 px-3 py-2 rounded-sm`}
    >
      {children}
      <h1 className="text-black">{ItemTitle}</h1>
    </Link>
  );
}

export function SideBarSubItem({
  ItemTitle,
  Href,
  IsSelected,
  children,
}: ItemProps) {
  return (
    <Link
      href={Href}
      className={`flex  items-center gap-4 cursor-pointer ${
        IsSelected ? "bg-neutral-300" : ""
      }  hover:bg-neutral-200  duration-300 px-3 py-2 rounded-sm`}
    >
      {children}
      <h1 className="text-black">{ItemTitle}</h1>
    </Link>
  );
}

export function SideBarDropItem({
  ItemTitle,
  LucideIcon,
  children,
}: GroupItemProps) {
  const [isOpen, setIsOpen] = useState<boolean>(false);
  return (
    <div className="flex  items-start gap-3 w-full rounded-sm flex-col">
      <button
        className="flex px-3  items-center justify-between w-full py-2 hover:bg-neutral-300 rounded-sm"
        onClick={() => {
          setIsOpen(!isOpen);
        }}
      >
        <div className="cursor-pointer flex items-center gap-4 w-full">
          <LucideIcon size={19} color="black" />
          <h1 className="text-black">{ItemTitle}</h1>
        </div>
        <div className="cursor-pointer">
          {isOpen ? (
            <ChevronUp size={19} color="black" />
          ) : (
            <ChevronDown size={19} color="black" />
          )}
        </div>
      </button>
      {isOpen && (
        <div
          className={`overflow-hidden  w-[100%]  ${
            isOpen ? "max-h-[1000px]" : "max-h-0"
          }`}
        >
          {children} {/* Always render */}
        </div>
      )}
    </div>
  );
}
