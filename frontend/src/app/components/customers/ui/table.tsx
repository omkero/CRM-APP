import React, { useEffect, useState } from "react";
import Row from "./row";
import Head from "./head";
import {
  BriefcaseBusiness,
  CalendarDays,
  Check,
  Mail,
  Mails,
  Phone,
  Signature,
  Trash,
} from "lucide-react";
import Checkbox from "../../ui/checkbox";
import "../../../globals.css";
import { Button, Flex, Dialog } from "@radix-ui/themes";
import { PorductType, SelectedProduct } from "@/app/products/types";
import { CustomerType } from "@/app/customers/types";

type Props = {
  data: CustomerType[];
};

export function CustomersTable({ data }: Props) {
  const [isSelectedAll, setIsSelectedAll] = useState<boolean>(false);
  const [selectedProducts, setSelectedProducts] = useState<CustomerType[]>(
    []
  );
  let SelectedProducts: SelectedProduct[]  = [];
  const cellPadding: number = 13;

  const DeleteAllProducts = () => (
    <Dialog.Root>
      <Dialog.Trigger>
        <button className="bg-neutral-200 text-black py-2 px-3 rounded-sm cursor-pointer flex items-center gap-2">
          <p>Delete</p>
          <Trash size={16} />
        </button>
      </Dialog.Trigger>

      <Dialog.Content maxWidth="450px">
        <Dialog.Title className=" font-medium">
          Are you absoloutley Sure ?
        </Dialog.Title>

        <div>
          <p>
            Lorem ipsum dolor sit amet consectetur adipisicing elit. Animi quod
            quis quibusdam! Ea, dolorum. Laborum ipsam
          </p>
        </div>
        <Flex gap="3" mt="4" justify="end">
          <Dialog.Close>
            <Button variant="soft" color="gray">
              Cancel
            </Button>
          </Dialog.Close>
          <Dialog.Close>
            <Button
              style={{ backgroundColor: "red" }}
              className=" bg-red-200 text-red-700"
            >
              Delete
            </Button>
          </Dialog.Close>
        </Flex>
      </Dialog.Content>
    </Dialog.Root>
  );

  useEffect(() => {
    console.log(selectedProducts);
    
  }, [SelectedProducts, isSelectedAll])
  return (
    <div className="w-full flex h-full overflow-x-auto flex-col gap-3">
      {selectedProducts.length >= 1 && (
        <div className="flex justify-between items-center w-full">
          <div className="flex items-center gap-2">
            <p>Selected All</p>
            <Check size={19} />
          </div>
          <DeleteAllProducts />
        </div>
      )}
      <div className="overflow-x-auto ">
        <table className="w-full table-auto  border-collapse rounded-tl-2xl rounded-tr-2xl">
          <thead className="rounded-tl-2xl rounded-tr-2xl">
            <Head
              cellPadding={cellPadding}
              isSelectedAll={isSelectedAll}
              setIsSelectedAll={setIsSelectedAll}
              data={data}
              setSelectedProducts={setSelectedProducts}
            />
          </thead>
          <tbody>
            {data?.map((item, i) => (
              <Row
                key={i}
                cellPadding={cellPadding}
                data={item}
                allData={data}
                SelectedProducts={selectedProducts}
                setSelectedProducts={setSelectedProducts}
              />
            ))}
          </tbody>
        </table>
      </div>
    </div>
  );
}
