"use client";
import Image from "next/image";
import { Sidebar } from "./components/sidebar/sidebar";
import { Navbar } from "./components/navbar/navbar";
import { ChangeEvent, FormEvent, useContext } from "react";
import { ConstantsContext } from "./providers/constantsProvider";

export default function Home() {
  const constatns = useContext(ConstantsContext);
  function widthDecider(): number {
    return constatns?.sideWidth == constatns?.minSideWidth
      ? constatns?.minSideWidth
      : constatns?.sideWidth;
  }
  return (
    <div className="max-h-screen w-full flex flex-col">
      <Navbar />
      <div className="flex h-screen">
        <Sidebar SelectedName="Dashboard" />
        <div
          style={{
            marginLeft:
              constatns?.sideWidth > constatns?.minSideWidth
                ? constatns?.sideWidth
                : constatns?.minSideWidth,
          }}
          className="bg-white w-[100%]  mt-16  h-auto flex flex-col justify-between p-5"
        >
          <h1>contetnt here</h1>
          <h1>contetnt here</h1>
          <input
            type="range"
            id="volume"
            onChange={(e: ChangeEvent<HTMLInputElement>) => {
              constatns?.setSideWidth(Number(e.target.value));
            }}
            value={constatns?.sideWidth}
            name="volume"
            min={constatns?.minSideWidth}
            max={500}
          />
        </div>
      </div>
    </div>
  );
}
