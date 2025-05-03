import { Activity, Ellipsis, Send } from "lucide-react";
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

const GirdItem = ({ item }: ItemProps) => {
  return (
    <div className="flex flex-col outline shadow-light-hover  gap-4 bg-neutral-100 outline-neutral-300  p-4 w-full">
      <div className="flex flex-col items-center  w-full justify-between">
        <div className=" h-46 w-full  bg-neutral-300 relative">
          <Image
            alt="img"
            src={BASE_URL + "/" + item?.product_cover}
            height={50}
            width={50}
            className="object-cover h-full w-full bg-neutral-400"
          />
          <div className=" absolute top-2 right-2">
            <button className="hover:bg-neutral-200 p-2 outline outline-transparent rounded-full hover:outline-neutral-300 cursor-pointer ">
              <Ellipsis size={19} color="black" />
            </button>
          </div>
        </div>
      </div>
      <div className="flex flex-col gap-2  w-full">
        <span className="text-xl rounded-2xl">{item?.product_title}</span>
        <span className="text-sm rounded-2xl">{item?.product_description}</span>
      </div>
      <div className="flex flex-col gap-4">
        <div className="flex items-center gap-3">
          <h1>Created At: </h1>
          <span className="text-sm text-neutral-500">{item?.created_at}</span>
        </div>
        <div className="flex w-full gap-2 items-center justify-between">
          <button className=" bg-white hover:bg-neutral-100 cursor-pointer  flex items-center w-[50%] justify-center gap-2 shadow-2xs p-[6px] outline outline-neutral-300 rounded-3xl">
            <Send size={16} />
            <span className="text-sm">Contact</span>
          </button>
          <button className=" bg-white hover:bg-neutral-100 cursor-pointer flex items-center w-[50%] justify-center gap-2 shadow-2xs p-[6px] outline outline-neutral-300 rounded-3xl">
            <Activity size={16} />
            <span className="text-sm">Activity</span>
          </button>
        </div>
      </div>
    </div>
  );
};

export default function ProductsGrid({ data }: Props) {
  return (
    <div className="h-full w-full grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4">
      {data?.map((item: PorductType, index: number) => (
        <GirdItem item={item} />
      ))}
    </div>
  );
}
