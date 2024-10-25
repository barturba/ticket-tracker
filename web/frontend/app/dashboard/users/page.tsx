import { getUsers } from "@/app/api/users/queries";
import AppHeading from "@/app/application-components/heading";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/app/components/table";
import { formatDateToLocal } from "@/app/lib/utils";
import PaginationApp from "@/app/ui/utils/pagination-app";
import { User, UsersResponse } from "@/types/users/base";
import type { Metadata } from "next";
import { Suspense } from "react";

export const metadata: Metadata = {
  title: "Users",
  description: "Manage system users",
};
interface SearchParamsProps {
  searchParams?: Promise<{
    query?: string;
    page?: number;
  }>;
}

interface UsersTableProps {
  users: User[];
}

function UsersTableSkeleton() {
  return (
    <div className="animate-pulse">
      <div className="h-8 bg-gray-200 rounded w-full mb-4" />
      {[...Array(5)].map((_, i) => (
        <div key={i} className="h-16 bg-gray-200 rounded w-full mb-2" />
      ))}
    </div>
  );
}

function UsersTable({ users }: UsersTableProps) {
  if (!users.length) {
    return (
      <div className="text-center py-8 text-gray-500">
        No users found. Try adjusting your search.
      </div>
    );
  }
  return (
    <Table className="mt-8 [--gutter:theme(spacing.6)] lg:[--gutter:theme(spacing.10)]">
      <TableHead>
        <TableRow>
          <TableHeader>ID </TableHeader>
          {/* <TableHeader>Updated date</TableHeader> */}
          <TableHeader>First Name</TableHeader>
          <TableHeader>Last Name</TableHeader>
        </TableRow>
      </TableHead>
      <TableBody>
        {users.map((user) => (
          <TableRow
            key={user.id}
            href={`/dashboard/users/${user.id}/edit`}
            title={`User #${user.id}`}
          >
            <TableCell>{user.id.slice(0, 8)}</TableCell>
            {/* <TableCell className="text-zinc-500">
                {formatDateToLocal(JSON.stringify(user.updated_at))}
              </TableCell> */}
            <TableCell>{user.first_name}</TableCell>
            <TableCell>{user.last_name}</TableCell>
          </TableRow>
        ))}
      </TableBody>
    </Table>
  );
}

async function UsersContent({ searchParams }: SearchParamsProps) {
  try {
    const resolvedParams = await searchParams;
    const query = resolvedParams?.query || "";
    const currentPage = Number(resolvedParams?.page) || 1;

    const usersResponse: UsersResponse = await getUsers(query, currentPage);
    return (
      <>
        <UsersTable users={usersResponse.users} />
        <div className="mt-6">
          <PaginationApp totalPages={usersResponse.metadata.last_page} />
        </div>
      </>
    );
  } catch (error) {
    console.error(`Error fetching users:`, error);
    return (
      <div className="rounded-md bg-red-50 p-4 mt-8">
        <div className="flex">
          <div className="ml-3">
            <h3 className="text-sm font-medium text-red-800">
              Error Loading Users
            </h3>
            <div className="mt-2 text-sm text-red-700">
              <p>Failed to load users. Please try again later.</p>
            </div>
          </div>
        </div>
      </div>
    );
  }
}

export default async function UsersPage({ searchParams }: SearchParamsProps) {
  return (
    <div className="space-y-6">
      <AppHeading
        name="Users"
        createLabel="Create User"
        createLink="/dashboard/users/create"
      />
      <Suspense fallback={<UsersTableSkeleton />}>
        <UsersContent searchParams={searchParams} />
      </Suspense>
    </div>
  );
}
