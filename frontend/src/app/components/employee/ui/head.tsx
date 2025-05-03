import React from "react";
import {
  BriefcaseBusiness,
  CalendarDays,
  Check,
  ChevronsUpDown,
  Mail,
  Mails,
  Phone,
  Signature,
} from "lucide-react";
import Checkbox from "../../ui/checkbox";
import { PorductType, SelectedProduct } from "@/app/products/types";

type Props = {
  cellPadding: number;
  setIsSelectedAll: React.Dispatch<boolean>;
  isSelectedAll: boolean;
  data: PorductType[]
  setSelectedProducts: React.Dispatch<SelectedProduct[]>
};

function Head({ cellPadding, isSelectedAll, setIsSelectedAll , data, setSelectedProducts}: Props) {
  return (
    <tr>
      <th
        style={{ padding: cellPadding }}
        className="border border-neutral-200 text-left"
      >
        <div className="flex items-center gap-6">
          <div className="flex items-center justify-center">
            <button
              className="flex items-center justify-center rounded-lg border-1 border-neutral-300 h-6 w-6  cursor-pointer"
              onClick={() => {
                if (isSelectedAll) {
                  setSelectedProducts([])
                  setIsSelectedAll(false)
                  return;
                }
                setSelectedProducts(data)
                setIsSelectedAll(!isSelectedAll)
              }}
            >
              {isSelectedAll && <Check size={17} />}
            </button>
          </div>
          <div className="flex items-center gap-3 ">
            <span className="text-sm text-black">Product ID</span>
            <ChevronsUpDown className="text-neutral-500" size={16} />
          </div>
        </div>
      </th>
      <th
        style={{ padding: cellPadding }}
        className="border border-neutral-200 text-left"
      >
        <div className="flex items-center gap-2">
          <span className="text-sm">Title</span>
          <ChevronsUpDown className="text-neutral-500" size={16} />
        </div>
      </th>
      <th
        style={{ padding: cellPadding }}
        className="border border-neutral-200 text-left"
      >
        <div className="flex items-center gap-2">
          <span className="text-sm">Description</span>
          <ChevronsUpDown className="text-neutral-500" size={16} />
        </div>
      </th>
      <th
        style={{ padding: cellPadding }}
        className="border border-neutral-200 text-left"
      >
        <div className="flex items-center gap-2">
          <span className="text-sm">Price</span>
          <ChevronsUpDown className="text-neutral-500" size={16} />
        </div>
      </th>
      <th
        style={{ padding: cellPadding }}
        className="border border-neutral-200 text-left"
      >
        <div className="flex items-center gap-2">
          <span className="text-sm">Image</span>
          <ChevronsUpDown className="text-neutral-500" size={16} />
        </div>
      </th>
      <th
        style={{ padding: cellPadding }}
        className="border border-neutral-200 text-left"
      >
        <div className="flex items-center gap-2">
          <span className="text-sm">Created At</span>
          <ChevronsUpDown className="text-neutral-500" size={16} />
        </div>
      </th>
      <th
        style={{ padding: cellPadding }}
        className="border border-neutral-200 text-left"
      >
        <span className="text-sm">Actions</span>
      </th>
    </tr>
  );
}

export default Head;
