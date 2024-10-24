import { Metadata } from "next";
import { notFound } from "next/navigation";
import HeadingEdit from "@/app/application-components/heading-edit";
import HeadingSubEdit from "@/app/application-components/heading-sub-edit";
import EditUserForm from "@/app/ui/sections/users/edit-form";
import { getUser } from "@/app/api/users/queries";
import { Suspense } from "react";
import { User } from "@/types/users/base";
import { title } from "process";

interface UserPageProps {
  params: Promise<{ id: string }>;
}

class UserFetchError extends Error {
  constructor(message: string, public statusCode: number) {
    super(message);
    this.name = "UserFetchError";
  }
}

function UserFormSkeleton() {
  return (
    <div className="animate-pulse">
      <div className="h-8 w-1/4 bg-gray-200 rounded mb-4" />
      <div className="space-y-4">
        <div className="h-12 bg-gray-200 rounded" />
        <div className="h-12 bg-gray-200 rounded" />
        <div className="h-12 bg-gray-200 rounded" />
      </div>
    </div>
  );
}

export async function generateMetadata({
  params,
}: UserPageProps): Promise<Metadata> {
  try {
    const resolvedParams = await params;
    const user = await getUser(resolvedParams.id);

    if (!user) {
      return {
        title: "User Not Found",
      };
    }

    return {
      title: `Edit User #${user.id}`,
      description: "Edit user details for ${user.first_name} ${user.last_name}",
    };
  } catch (error) {
    console.error(`generate metadata error:`, error);
    return {
      title: "Edit User",
      description: "User edit page",
    };
  }
}

function UserFormWrapper({ user }: { user: User }) {
  return (
    <Suspense fallback={<UserFormSkeleton />}>
      <EditUserForm user={user} />
    </Suspense>
  );
}

export default async function UserEditPage({ params }: UserPageProps) {
  try {
    const resolvedParams = await params;
    const id = resolvedParams.id;

    if (!id || typeof id !== "string") {
      notFound();
    }

    const user = await getUser(id);

    if (!user) {
      notFound();
    }

    const userState = user.state || "unknown";

    return (
      <div className="space-y-6">
        <HeadingEdit name="Users" backLink="/dashboard/users" />

        <HeadingSubEdit
          name={`User #${user.id}`}
          badgeState={user.state}
          badgeText={user.state}
        />

        <UserFormWrapper user={user} />
      </div>
    );
  } catch (error) {
    console.error(`user edit page error:`, error);

    if (error instanceof UserFetchError) {
      if (error.statusCode === 404) {
        notFound();
      }
    }

    throw error;
  }
}

export function ErrorBoundary({ error }: { error: Error }) {
  return (
    <div className="rounded-md bg-red-50 p-4">
      <div className="flex">
        <div className="ml-3">
          <h3 className="text-sm font-medium text-red-800">
            Error Loading User
          </h3>
          <div className="mt-2 text-sm text-red-700">
            <p>{error.message}</p>
          </div>
        </div>
      </div>
    </div>
  );
}
