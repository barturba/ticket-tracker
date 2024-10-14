import AppHeading from "@/app/application-components/heading";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/app/components/table";
import { getUsers } from "@/app/api/users/users";
import { UsersData } from "@/app/lib/definitions/users";
import { formatDateToLocal } from "@/app/lib/utils";
import PaginationApp from "@/app/ui/utils/pagination-app";
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

  return (
    <>
      <AppHeading
        name="Users"
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
      <PaginationApp totalPages={usersData.metadata.last_page} />
    </>
  );
}
