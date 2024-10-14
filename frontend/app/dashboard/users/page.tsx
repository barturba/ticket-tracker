import AppHeading from "@/app/application-components/heading";
import {
  Pagination,
  PaginationGap,
  PaginationList,
  PaginationNext,
  PaginationPage,
  PaginationPrevious,
} from "@/app/components/pagination";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/app/components/table";
import { getUsers } from "@/app/lib/actions/users";
import { UserData, UsersData } from "@/app/lib/definitions/users";
import { formatDateToLocal } from "@/app/lib/utils";
import type { Metadata } from "next";

export const metadata: Metadata = {
  title: "Users",
};

export default async function Users(props: {
  searchParams?: Promise<{
    query?: string;
    page?: string;
  }>;
}) {
  const searchParams = await props.searchParams;
  const query = searchParams?.query || "";
  const currentPage = Number(searchParams?.page) || 1;

  const usersData: UsersData = await getUsers(query, currentPage);

  console.log(`usersData: ${JSON.stringify(usersData, null, 2)}`);

  return (
    <>
      <AppHeading
        name="User"
        createLabel="Create User"
        createLink="/dashboard/users/create"
      />
      <Table className="mt-8 [--gutter:theme(spacing.6)] lg:[--gutter:theme(spacing.10)]">
        <TableHead>
          <TableRow>
            <TableHeader>ID </TableHeader>
            <TableHeader>Updated date</TableHeader>
            <TableHeader>First Name</TableHeader>
            <TableHeader>Last Name</TableHeader>
          </TableRow>
        </TableHead>
        <TableBody>
          {usersData.users.map((user) => (
            <TableRow
              key={user.id}
              href={`/dashboard/users/${user.id}/edit`}
              title={`User #${user.id}`}
            >
              <TableCell>{user.id}</TableCell>
              <TableCell className="text-zinc-500">
                {formatDateToLocal(user.updated_at)}
              </TableCell>
              <TableCell>{user.first_name}</TableCell>
              <TableCell>{user.last_name}</TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>

      <Pagination>
        <PaginationPrevious href="?page=2" />
        <PaginationList>
          <PaginationPage href="?page=1">1</PaginationPage>
          <PaginationPage href="?page=2">2</PaginationPage>
          <PaginationPage href="?page=3" current>
            3
          </PaginationPage>
          <PaginationPage href="?page=4">4</PaginationPage>
          <PaginationGap />
          <PaginationPage href="?page=65">65</PaginationPage>
          <PaginationPage href="?page=66">66</PaginationPage>
        </PaginationList>
        <PaginationNext href="?page=4" />
      </Pagination>
    </>
  );
}
