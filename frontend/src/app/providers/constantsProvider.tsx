"use client";
import React from "react";
import {
  useContext,
  createContext,
  useState,
  Dispatch,
  SetStateAction,
} from "react";

interface ConstantsType {
  sideWidth: number;
  setSideWidth: React.Dispatch<number>;
  minSideWidth: number;
  maxSideWidth: number;
}

type Props = {
  children: React.ReactNode;
};

export const ConstantsContext = createContext<ConstantsType | any>(undefined);

const ConstantsProvider = ({ children }: Props) => {
  const [sideWidth, setSideWidth] = useState<number>(310); // default value is 420px
  const minSideWidth: number = 250;
  const maxSideWidth: number = 500;

  return (
    <ConstantsContext.Provider
      value={{ sideWidth, setSideWidth, minSideWidth, maxSideWidth }}
    >
      {children}
    </ConstantsContext.Provider>
  );
};

export default ConstantsProvider;
