"use client";
import { DropdownMenu } from "@radix-ui/themes";
import { Ellipsis, EllipsisVertical } from "lucide-react";
import React from "react";
import "../../../globals.css";
import { PorductType } from "@/app/products/types";
import Image from "next/image";
import { BASE_URL } from "@/app/constant";

type Props = {
  data: PorductType[];
};

type ItemProps = {
  item: PorductType;
};

const ItemList = ({ item }: ItemProps) => {
  return (
    <div className=" bg-neutral-100 shadow-light-hover outline h-auto flex items-center justify-between outline-neutral-200    p-3 px-5 w-full">
      <div className="flex gap-4">
        <Image
          alt="img"
          src={BASE_URL + "/" + item?.product_cover}
          height={50}
          width={50}
          className="h-26  w-26 rounded-2xl bg-neutral-400"
        />
        <div className="flex flex-col  w-[620px] gap-2">
          <h1 className="text-xl font-medium">{item?.product_title}</h1>
          <p className="text-xs">{item?.product_description}</p>
          <div className="flex items-center">
            <span className="text-sm">{item?.created_at}</span>
          </div>
        </div>
      </div>
      <div className="flex flex-col items-center  h-full justify-between gap-4">
        <DropdownMenu.Root>
          <DropdownMenu.Trigger>
            <button className="hover:bg-neutral-300 p-2 cursor-pointer duration-200 rounded-full">
              <Ellipsis />
            </button>
          </DropdownMenu.Trigger>
          <DropdownMenu.Content>
            <DropdownMenu.Item shortcut="⌘ E">Edit</DropdownMenu.Item>
            <DropdownMenu.Item shortcut="⌘ D">Duplicate</DropdownMenu.Item>
            <DropdownMenu.Separator />
            <DropdownMenu.Item shortcut="⌘ N">Archive</DropdownMenu.Item>

            <DropdownMenu.Sub>
              <DropdownMenu.SubTrigger>More</DropdownMenu.SubTrigger>
              <DropdownMenu.SubContent>
                <DropdownMenu.Item>Move to project…</DropdownMenu.Item>
                <DropdownMenu.Item>Move to folder…</DropdownMenu.Item>

                <DropdownMenu.Separator />
                <DropdownMenu.Item>Advanced options…</DropdownMenu.Item>
              </DropdownMenu.SubContent>
            </DropdownMenu.Sub>

            <DropdownMenu.Separator />
            <DropdownMenu.Item>Share</DropdownMenu.Item>
            <DropdownMenu.Item>Add to favorites</DropdownMenu.Item>
            <DropdownMenu.Separator />
            <DropdownMenu.Item shortcut="⌘ ⌫" color="red">
              Delete
            </DropdownMenu.Item>
          </DropdownMenu.Content>
        </DropdownMenu.Root>
        <span className="text-2xl font-medium">{item?.product_price}$</span>
        <span className="text-3xl font-medium"></span>
      </div>
    </div>
  );
};

export default function ProductLists({ data }: Props) {
  return (
    <div className="h-full w-full  flex flex-col gap-3">
      {data?.map((item: PorductType, index: number) => (
        <ItemList item={item} />
      ))}
    </div>
  );
}
