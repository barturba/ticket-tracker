"use client";

import { generatePagination } from "@/app/lib/utils";
import { usePathname, useSearchParams } from "next/navigation";
import {
  Pagination,
  PaginationGap,
  PaginationList,
  PaginationNext,
  PaginationPage,
  PaginationPrevious,
} from "@/app/components/pagination";

export default function PaginationApp({ totalPages }: { totalPages: number }) {
  const pathname = usePathname();
  const searchParams = useSearchParams();
  const currentPage = Number(searchParams.get("page")) || 1;

  const createPageURL = (pageNumber: number | string) => {
    const params = new URLSearchParams(searchParams);
    params.set("page", pageNumber.toString());
    return `${pathname}?${params.toString()}`;
  };

  const allPages = generatePagination(currentPage, totalPages);

  return (
    <div className="mt-2">
      <Pagination>
        <PaginationPrevious
          href={currentPage > 1 ? createPageURL(currentPage - 1) : null}
        />
        <PaginationList>
          {allPages.map((page, index) => {
            let position: "first" | "last" | "single" | "middle" | undefined;

            if (index === 0) position = "first";
            if (index === allPages.length - 1) position = "last";
            if (allPages.length === 1) position = "single";
            if (page === "...") position = "middle";

            if (position === "middle") {
              return <PaginationGap key={`gap-${index}`} />;
            } else {
              return (
                <PaginationPage
                  key={`page-${page}`}
                  current={currentPage === page}
                  href={createPageURL(page)}
                >
                  {page}
                </PaginationPage>
              );
            }
          })}
        </PaginationList>

        <PaginationNext
          href={
            currentPage < totalPages ? createPageURL(currentPage + 1) : null
          }
        />
      </Pagination>
    </div>
  );
}
