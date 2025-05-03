"use client";

import Pagination from "@mui/material/Pagination";
import Stack from "@mui/material/Stack";
import { useState } from "react";
import { useSearchParams } from 'next/navigation';
import { productsPerPage } from "@/app/constant";

interface Props {
  currentPage: number;
  totalPages: number;
  onPageChange: (page: number) => void;
}

export function PaginationUI({ currentPage, totalPages, onPageChange }: Props) {
  // useSearchParams to get the page number from the url 
  const searchParams = useSearchParams();
  const id: any = searchParams.get('page');
  const [pageNum, setPageNum] = useState(parseInt(id)) // parse it to number and the state will listen to the query param

  // make sure to get accurate pages count
  let Pages = Math.ceil(totalPages / productsPerPage)
  return (
    <div className="w-full flex items-center justify-end">
      <Pagination
        page={pageNum || 1}
        count={Pages}
        size="large"
        onChange={(_, page) => {
          setPageNum(page)
          onPageChange(page)
        }} // rounded shape MUI Component
        shape="rounded"
      />
    </div>
  );
}
