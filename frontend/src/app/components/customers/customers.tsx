"use client";

import {
  useContext,
  useEffect,
  useState,
  ChangeEvent,
  FormEvent,
  JSX,
  Suspense,
} from "react";
import { useRouter, useSearchParams, usePathname } from "next/navigation";
import { toast } from "sonner";
import {
  TableProperties,
  List,
  Grid3x3,
  SortDesc,
  ListFilterPlus,
  Plus,
} from "lucide-react";
import { Button, Dialog, DropdownMenu, Flex } from "@radix-ui/themes";

import { ConstantsContext } from "@/app/providers/constantsProvider";
import { CustomersTable } from "./ui/table";

import { CreateProductType, PorductType } from "@/app/products/types";
import { CreateProductAction } from "@/app/products/createProduct";
import { NEXT_BASE_URL, soonerDefaultDuration } from "@/app/constant";
import { PaginationUI } from "../ui/pagination";
import { AlertCircle } from "lucide-react";

import { Alert, AlertDescription, AlertTitle } from "@/components/ui/alert";
import { CustomerType } from "@/app/customers/types";

type Props = {
  data: CustomerType[];
  page: number;
  totalCustomers: number;
};

export default function Customers({ data, page, totalCustomers }: Props) {
  const router = useRouter();
  const pathname = usePathname();
  const searchParams = useSearchParams();

  const constants = useContext(ConstantsContext);

  const [isLoading, setIsLoading] = useState<boolean>(false)
  const [currentTab, setCurrentTab] = useState("table");
  const [productTitle, setProductTitle] = useState("");
  const [productDescription, setProductDescription] = useState("");
  const [productPrice, setProductPrice] = useState<number | undefined>();
  const [productImage, setProductImage] = useState<File | undefined>();
  const [range, setRange] = useState<[number, number]>([0, 999999]);
  const [isError, setIsError] = useState<boolean>(false);
  const [errorMessage, setErrorMessage] = useState<any>("");
  const [isCreateModal, setIsCreateModal] = useState<boolean>(false);

  useEffect(() => {
    console.log(data);
  }, [data]);


  const handleSubmit = async (e: FormEvent) => {
    e.preventDefault();

    setIsLoading(true)

    if (!productTitle) {
      setIsError(true);
      setErrorMessage("missing product title");
      return;
    }
    if (!productPrice) {
      setIsError(true);
      setErrorMessage("missing product price");
      return;
    }
    if (!productDescription) {
      setIsError(true);
      setErrorMessage("missing product description");
      return;
    }
    if (!productImage) {
      setIsError(true);
      setErrorMessage("missing product image");
      return;
    }

    const payload: CreateProductType = {
      product_title: productTitle,
      product_description: productDescription,
      product_price: productPrice,
      product_cover: productImage,
    };

    const response = await CreateProductAction(payload);
    if (response?.status === 200) {
      setIsLoading(false)

      setIsCreateModal(false);

      // set push path with current page in query params
      const pageNum: any = searchParams.get("page");
      const params = new URLSearchParams(searchParams.toString());
      params.set("page", pageNum); // Set or update page param

      toast.success("Product Created.",  { duration: soonerDefaultDuration });
      router.push(`${pathname}?${params.toString()}`);
    }
    if (response?.status != 200) {
      setIsLoading(false)
      setIsError(true);
      setErrorMessage(response?.message);
    }
  };

  const CreateProductForm = () => (
    <div>
      <form onSubmit={handleSubmit}>
        <div className="flex flex-col gap-4">
          <Dialog.Title>Create Product</Dialog.Title>

          {isError ? (
            <Alert variant="destructive">
              <AlertCircle className="h-4 w-4" />
              <AlertTitle>Error</AlertTitle>
              <AlertDescription>{errorMessage}</AlertDescription>
            </Alert>
          ) : (
            <p>Make sure to fill all inputs.</p>
          )}

          <div className="flex flex-col gap-3">
            <div className="w-full flex flex-col gap-2">
              <label className="font-bold">Title</label>

              <input
                type="text"
                value={productTitle}
                className="w-full py-[6px] px-2 outline outline-neutral-400 rounded-sm focus:outline-blue-500"
                placeholder="Enter product title"
                onChange={(e: ChangeEvent<HTMLInputElement>) =>
                  setProductTitle(e.target.value)
                }
              />
            </div>

            <div className="w-full flex flex-col gap-2">
              <label className="font-bold">Price</label>
              <input
                type="number"
                value={productPrice ?? ""}
                className="w-full py-[6px] px-2 outline outline-neutral-400 rounded-sm focus:outline-blue-500"
                placeholder="Enter product price"
                onChange={(e: ChangeEvent<HTMLInputElement>) =>
                  setProductPrice(Number(e.target.value))
                }
              />
            </div>

            <div className="w-full flex flex-col gap-2">
              <label className="font-bold">Description</label>
              <input
                type="text"
                value={productDescription}
                className="w-full py-[6px] px-2 outline outline-neutral-400 rounded-sm focus:outline-blue-500"
                placeholder="Enter product description"
                onChange={(e: ChangeEvent<HTMLInputElement>) =>
                  setProductDescription(e.target.value)
                }
              />
            </div>

            <div className="w-full flex flex-col gap-2">
              <label className="font-bold">Image</label>
              <input
                type="file"
                accept="image/*"
                className="w-full py-[6px] px-2 outline outline-neutral-400 rounded-sm focus:outline-blue-500"
                onChange={(e: ChangeEvent<HTMLInputElement>) => {
                  const file = e.target.files?.[0];
                  if (file) setProductImage(file);
                }}
              />
            </div>

            <Flex gap="3" mt="4" justify="end">
              <Dialog.Close>
                <Button variant="soft" color="gray">
                  Cancel
                </Button>
              </Dialog.Close>
              <Button loading={isLoading} type="submit">Save</Button>
            </Flex>
          </div>
        </div>
      </form>
    </div>
  );

  return (
    <div
      style={{
        marginLeft:
          constants?.sideWidth > constants?.minSideWidth
            ? constants.sideWidth
            : constants.minSideWidth,
      }}
      className="bg-white w-full mt-16 h-auto flex flex-col justify-between p-9"
    >
      <div className="flex flex-col gap-4 h-full">
        <div className="flex items-center justify-between w-full">
          <div className="flex gap-9 items-center">
            <h1 className="text-2xl font-medium">Customers</h1>
            <div className="flex items-center gap-5">
              {["table", "list", "grid"].map((tab) => {
                const icons: Record<string, JSX.Element> = {
                  table: <TableProperties size={22} />,
                  list: <List size={22} />,
                  grid: <Grid3x3 size={22} />,
                };
                return (
                  <button
                    key={tab}
                    onClick={() => setCurrentTab(tab)}
                    className={`cursor-pointer rounded-md border border-neutral-300 flex flex-col items-center gap-2 ${
                      currentTab === tab ? "bg-neutral-200" : ""
                    }`}
                  >
                    <div className="flex items-center gap-2 px-5 py-2">
                      {icons[tab]}
                      <p className="capitalize">{tab}</p>
                    </div>
                  </button>
                );
              })}
            </div>
          </div>

          <div className="flex items-center gap-4">
            <DropdownMenu.Root>
              <DropdownMenu.Trigger>
                <button className="border border-neutral-300 py-2 px-3 rounded-sm flex items-center gap-2">
                  <p>Sort By</p>
                  <SortDesc size={16} />
                </button>
              </DropdownMenu.Trigger>
              <DropdownMenu.Content>
                <DropdownMenu.Item>Price Low to High</DropdownMenu.Item>
                <DropdownMenu.Item>Price High to Low</DropdownMenu.Item>
              </DropdownMenu.Content>
            </DropdownMenu.Root>

            <DropdownMenu.Root>
              <DropdownMenu.Trigger>
                <button className="border border-neutral-300 py-2 px-3 rounded-sm flex items-center gap-2">
                  <p>Filter</p>
                  <ListFilterPlus size={16} />
                </button>
              </DropdownMenu.Trigger>
              <DropdownMenu.Content>
                <div className="w-60 p-3 flex flex-col gap-2">
                  <h1 className="text-xl font-medium">By Price</h1>
                  <input
                    type="number"
                    className="w-full"
                    placeholder="From"
                    value={range[0]}
                    onChange={(e) =>
                      setRange([Number(e.target.value), range[1]])
                    }
                  />
                  <input
                    type="number"
                    className="w-full"
                    placeholder="To"
                    value={range[1]}
                    onChange={(e) =>
                      setRange([range[0], Number(e.target.value)])
                    }
                  />
                </div>
              </DropdownMenu.Content>
            </DropdownMenu.Root>

            <Dialog.Root open={isCreateModal} onOpenChange={setIsCreateModal}>
              <Dialog.Trigger>
                <button className="bg-neutral-200 border-1 border-transparent text-black  py-2 px-3 rounded-sm cursor-pointer flex items-center gap-2">
                  <p>Add Customer</p>
                  <Plus size={16} />
                </button>
              </Dialog.Trigger>
              <Dialog.Content>
                <Suspense fallback={<div>loading ...</div>}>
                  {CreateProductForm()}
                </Suspense>
              </Dialog.Content>
            </Dialog.Root>
          </div>
        </div>
        <CustomersTable data={data} />
        <PaginationUI
          totalPages={totalCustomers}
          currentPage={page}
          onPageChange={(pageNum) => {
            router.push(NEXT_BASE_URL + `/customers?page=${pageNum}`);
          }}
        />
      </div>
    </div>
  );
}
