import { lusitana } from "@/app/ui/fonts";
import { CreateUser } from "@/app/ui/users/buttons";
import Pagination from "@/app/ui/utils/pagination";
import Table from "@/app/ui/users/table";
import Search from "@/app/ui/search";
import { Metadata } from "next";
import { Suspense } from "react";
import { fetchUsers } from "@/app/lib/actions/users";
import { UserData } from "@/app/lib/definitions/users";
import { UsersTableSkeleton } from "@/app/ui/skeletons/users";

export const metadata: Metadata = {
  title: "Users",
};

export default async function Page(props: {
  searchParams?: Promise<{
    query?: string;
    page?: string;
  }>;
}) {
  const searchParams = await props.searchParams;
  const query = searchParams?.query || "";
  const currentPage = Number(searchParams?.page) || 1;

  const userData: UserData = await fetchUsers(query, currentPage);

  const totalPages = userData.metadata.last_page;
  const users = userData.users;

  return (
    <div className="w-full">
      <div className="flex w-full items-center justify-between">
        <h1 className={`${lusitana.className} text-2xl`}>Users</h1>
      </div>

      <div className="mt-4 flex items-center justify-between gap-2 md:mt-8">
        <Search placeholder="Search users..." />
        <CreateUser />
      </div>
      <Suspense key={query + currentPage} fallback={<UsersTableSkeleton />}>
        <Table users={users} />
      </Suspense>
      <div className="mt-5 flex w-full justify-center">
        <Pagination totalPages={totalPages} />
      </div>
    </div>
  );
}
