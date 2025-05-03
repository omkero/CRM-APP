import React, { useState, useEffect, FormEvent } from "react";
import Checkbox from "../../ui/checkbox";
import { Check, EllipsisVertical, Trash, X } from "lucide-react";
import {
  Button,
  Dialog,
  DropdownMenu,
  Flex,
  Spinner,
  Text,
  TextField,
} from "@radix-ui/themes";
import { PorductType, SelectedProduct } from "@/app/products/types";
import { BASE_URL, NEXT_BASE_URL, productsPerPage, soonerDefaultDuration } from "@/app/constant";
import { toast } from "sonner";
import { DeleteProductAction } from "@/app/products/deleteProduct";
import { useRouter, useSearchParams, usePathname } from "next/navigation";
import { ParseDate } from "@/lib/parseData";
import Image from "next/image";
import { CustomerType } from "@/app/customers/types";

type Props = {
  cellPadding: number;
  data: CustomerType;
  allData: CustomerType[];
  SelectedProducts: CustomerType[];
  setSelectedProducts: React.Dispatch<React.SetStateAction<CustomerType[]>>;
};

function Row({
  cellPadding,
  data,
  allData,
  SelectedProducts,
  setSelectedProducts,
}: Props) {
  const router = useRouter();
  const pathname = usePathname();
  const searchParams = useSearchParams();
  const pageNum: any = searchParams.get("page");

  const selectedData: CustomerType = data;

  const [isCreateModal, setIsCreateModal] = useState<boolean>(false);
  const [isViewImage, setIsViewImage] = useState<boolean>(false);
  const [isTrigger, setIsTrigger] = useState<boolean>(false);
  const [pageNuber, setPageNumber] = useState<number>(pageNum);
  const [isLoading, setIsLoading] = useState<boolean>(false);

  const isChecked = SelectedProducts.some(
    (product) => product?.customer_id === data?.customer_id
  );

  const handleToggle = () => {
    setSelectedProducts((prev) => {
      const alreadySelected = prev.some(
        (product) => product?.customer_id === selectedData?.customer_id
      );
      if (alreadySelected) {
        return prev.filter(
          (product) => product?.customer_id !== selectedData?.customer_id
        );
      } else {
        return [...prev, selectedData];
      }
    });
  };

  async function HandleDeleteProducts(e: FormEvent) {
    e.preventDefault();

    setIsLoading(true);
    setIsTrigger(false);

    const response = await DeleteProductAction(data?.customer_id);
    const targetedPage = Math.ceil(allData.length / productsPerPage);

    if (response?.status === 200) {
      setIsCreateModal(false);
      setIsLoading(false);

      // set push path with current page in query params
      const params = new URLSearchParams(searchParams.toString());
      params.set("page", pageNum); // Set or update page param

      toast.success("Product Deleted.",  { duration: soonerDefaultDuration });
      router.push(`${pathname}?${params.toString()}`);
    }
    if (response?.status != 200) {
      setIsLoading(false);
      toast.error(response?.message,  { duration: soonerDefaultDuration });
    }
  }

  const DeleteProduct = () => (
    <Dialog.Root open={isCreateModal} onOpenChange={setIsCreateModal}>
      <Dialog.Trigger>
        <Dialog.Close>
          <button className="bg-neutral-200 text-black py-1 px-3 rounded-sm cursor-pointer flex items-center gap-2">
            <p>Delete</p>
            <Trash size={16} />
          </button>
        </Dialog.Close>
      </Dialog.Trigger>
      <Dialog.Content maxWidth="450px">
        <form onSubmit={HandleDeleteProducts}>
          <Dialog.Title className=" font-medium">
            Are you absoloutley Sure ?
          </Dialog.Title>
          <div>
            <p>
              Lorem ipsum dolor sit amet consectetur adipisicing elit. Animi
              quod quis quibusdam! Ea, dolorum. Laborum ipsam
            </p>
          </div>
          <Flex gap="3" mt="4" justify="end">
            <Dialog.Close>
              <Button
                onClick={() => {
                  setIsTrigger(false);
                }}
                variant="soft"
                color="gray"
              >
                Cancel
              </Button>
            </Dialog.Close>

            <Button type="submit">
              {isLoading ? <Spinner size="2" /> : "Delete"}
            </Button>
          </Flex>
        </form>
      </Dialog.Content>
    </Dialog.Root>
  );

  const ViewImage = () => {
    return (
      <Dialog.Root open={isViewImage} onOpenChange={setIsViewImage}>
        <Dialog.Trigger>
          <Dialog.Close>
            <button className="bg-neutral-200 text-black py-1 px-3 rounded-sm cursor-pointer flex items-center gap-2">
              <p>View</p>
            </button>
          </Dialog.Close>
        </Dialog.Trigger>
        <Dialog.Content>
          <div className="relsative">
            <button
              className="bg-neutral-200 opacity-90 text-black py-3 px-3 absolute top-6 right-6 cursor-pointer rounded-full flex items-center gap-2"
              onClick={() => {
                setIsViewImage(false);
              }}
            >
              <X size={19} color="black" />
            </button>
            <Image
              src={"http://example.com/gg.jpg"}
              alt="img"
              height={500}
              width={600}
            />
          </div>
        </Dialog.Content>
      </Dialog.Root>
    );
  };

  return (
    <tr>
      <td
        style={{ padding: cellPadding }}
        className="border border-neutral-200"
      >
        <div className="flex items-center gap-2 relative">
          <button
            className="flex items-center justify-center rounded-lg border-1 border-neutral-300 h-6 w-6 cursor-pointer"
            onClick={handleToggle}
          >
            {isChecked && <Check size={17} />}
          </button>
          <span className="text-base ml-6">{data?.customer_id}</span>
        </div>
      </td>
      <td
        style={{ padding: cellPadding }}
        className="border border-neutral-200"
      >
        <span className="text-xs">{data?.customer_full_name}</span>
      </td>
      <td
        style={{ padding: cellPadding }}
        className="border border-neutral-200"
      >
        <span className="text-xs">{data?.customer_position}</span>
      </td>
      <td
        style={{ padding: cellPadding }}
        className="border   border-neutral-200"
      >
        <span className="text-xs bg-violet-100 text-violet-600 rounded-lg p-2 text-center w-full">
          {data?.customer_phone_number}$
        </span>
      </td>
      <td
        style={{ padding: cellPadding }}
        className="border border-neutral-200"
      >
        {ViewImage()}
      </td>
      <td
        style={{ padding: cellPadding }}
        className="border border-neutral-200"
      >
        <span className="text-xs">{ParseDate(data?.created_at)}</span>
      </td>
      <td
        style={{ padding: cellPadding }}
        className="border border-neutral-200 relative"
      >
        <div className="  flex items-center justify-center relative">
          <DropdownMenu.Root open={isTrigger} onOpenChange={setIsTrigger}>
            <DropdownMenu.Trigger>
              <button className="hover:bg-neutral-300 p-2 cursor-pointer duration-200 rounded-full">
                <EllipsisVertical />
              </button>
            </DropdownMenu.Trigger>
            <DropdownMenu.Content>
              <DropdownMenu.Item shortcut="⌘ E">Edit</DropdownMenu.Item>
              <DropdownMenu.Separator />
              <DeleteProduct />
            </DropdownMenu.Content>
          </DropdownMenu.Root>
        </div>
      </td>
    </tr>
  );
}

export default Row;
