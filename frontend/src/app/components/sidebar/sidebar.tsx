"use client";
import React from "react";
import { SideBarItem, SideBarDropItem, SideBarSubItem } from "./ui/item";
import SideBarGroup from "./ui/group";
import {
  Landmark,
  LayoutDashboard,
  ChartColumnIncreasing,
  ClipboardList,
  LogOut,
  ChevronUp,
  ChevronDown,
  Scan,
  UsersRound,
  ChartBar,
  SquareKanban,
  Megaphone,
  Settings,
  Handshake,
  Bug,
  Bell,
  Calendar1,
  MessageSquare,
  Layers,
  KeyRound,
  Package,
} from "lucide-react";
import { useContext } from "react";
import { ConstantsContext } from "@/app/providers/constantsProvider";

type Props = {
  SelectedName: string;
};

export const Sidebar = ({ SelectedName }: Props) => {
  const constatns = useContext(ConstantsContext);
  function IsSelected(name: string): boolean {
    return SelectedName == name;
  }
  return (
    <div
      style={{
        minWidth:
          constatns?.sideWidth > constatns?.minSideWidth
            ? constatns?.sideWidth
            : constatns?.minSideWidth,
      }}
      className=" bg-neutral-200 fixed h-[100%] border-r-2 border-r-neutral-300 flex flex-col "
    >
      <div className="px-4 py-4 flex flex-col gap-7 h-[100%] justify-start  mt-20  overflow-y-scroll">
        <SideBarGroup GroupTitle="Main">
          <SideBarItem
            ItemTitle="Dashboard"
            Href={"/"}
            IsSelected={IsSelected("Dashboard")}
          >
            <LayoutDashboard size={19} />
          </SideBarItem>

          <SideBarItem
            ItemTitle="Notification"
            Href={"/"}
            IsSelected={IsSelected("Awe")}
          >
            <Bell size={19} />
          </SideBarItem>
        </SideBarGroup>

        <SideBarGroup GroupTitle="Management">
          <SideBarDropItem ItemTitle="Tasks" LucideIcon={ChartColumnIncreasing}>
            <div className="flex  px-3 py-2 flex-col gap-1 bg-white rounded-lg w-full">
              <SideBarSubItem
                ItemTitle="Create Task"
                Href={"/scan_task"}
                IsSelected={IsSelected("scaning tasks")}
              ></SideBarSubItem>
              <SideBarSubItem
                ItemTitle="My Tasks"
                Href={"/scan_task"}
                IsSelected={IsSelected("scaning tasks")}
              ></SideBarSubItem>
            </div>
          </SideBarDropItem>

          <SideBarItem
            ItemTitle="Products"
            Href={"/products"}
            IsSelected={IsSelected("Products")}
          >
            <Package size={19} color="black" />
          </SideBarItem>

          <SideBarDropItem ItemTitle="Deals" LucideIcon={Handshake}>
            <div className="flex  px-3 py-2 flex-col gap-1 bg-white rounded-lg w-full">
              <SideBarSubItem
                ItemTitle="Add Deal"
                Href={"/scan_task"}
                IsSelected={IsSelected("scaning tasks")}
              ></SideBarSubItem>
              <SideBarSubItem
                ItemTitle="Current Deals"
                Href={"/scan_task"}
                IsSelected={IsSelected("scaning tasks")}
              ></SideBarSubItem>
            </div>
          </SideBarDropItem>

          <SideBarItem
            ItemTitle="Employees"
            Href={"/employee"}
            IsSelected={IsSelected("Employee")}
          >
            <UsersRound size={19} color="black" />
          </SideBarItem>

          <SideBarItem
            ItemTitle="Reports"
            Href={"/"}
            IsSelected={IsSelected("Awe")}
          >
            <Megaphone size={19} />
          </SideBarItem>

          <SideBarItem
            ItemTitle="Customers"
            Href={"/customers"}
            IsSelected={IsSelected("Customers")}
          >
            <Layers size={19} />
          </SideBarItem>

          <SideBarItem
            ItemTitle="Roles - Permissions"
            Href={"/"}
            IsSelected={IsSelected("Awe")}
          >
            <KeyRound size={19} />
          </SideBarItem>
        </SideBarGroup>

        <SideBarGroup GroupTitle="Analytics">
          <SideBarItem
            ItemTitle="Global Overview"
            Href={"/customers"}
            IsSelected={IsSelected("Players")}
          >
            <ChartBar size={19} color="black" />
          </SideBarItem>

          <SideBarDropItem ItemTitle="More" LucideIcon={SquareKanban}>
            <div className="flex  px-3 py-2 flex-col gap-1 bg-white rounded-lg w-full">
              <SideBarSubItem
                ItemTitle="Customers Charts"
                Href={"/scan_task"}
                IsSelected={IsSelected("scaning tasks")}
              ></SideBarSubItem>
              <SideBarSubItem
                ItemTitle="Products Charts"
                Href={"/scan_task"}
                IsSelected={IsSelected("scaning tasks")}
              ></SideBarSubItem>
              <SideBarSubItem
                ItemTitle="Tasks Charts"
                Href={"/scan_task"}
                IsSelected={IsSelected("scaning tasks")}
              ></SideBarSubItem>
            </div>
          </SideBarDropItem>
        </SideBarGroup>
        <SideBarGroup GroupTitle="Other">
          <SideBarItem
            ItemTitle="Settings"
            Href={"/"}
            IsSelected={IsSelected("Awe")}
          >
            <Settings size={19} />
          </SideBarItem>
          <SideBarItem
            ItemTitle="Bug Reports"
            Href={"/"}
            IsSelected={IsSelected("Awe")}
          >
            <Bug size={19} />
          </SideBarItem>
        </SideBarGroup>
      </div>

      <div className="px-7 py-4 border-t-2 border-t-neutral-300">
        <button className="border-2 flex items-center justify-center gap-3 border-neutral-300 bg-neutral-200 w-full p-3 cursor-pointer hover:bg-neutral-300">
          <LogOut size={19} />
          <p>sign out</p>
        </button>
      </div>
    </div>
  );
};
