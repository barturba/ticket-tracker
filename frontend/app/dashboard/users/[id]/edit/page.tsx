import { Metadata } from "next";
import { notFound } from "next/navigation";
import { getUser } from "@/app/lib/actions/users";
import HeadingEdit from "@/app/application-components/heading-edit";
import HeadingSubEdit from "@/app/application-components/heading-sub-edit";
import EditUserForm from "@/app/ui/sections/users/edit-form";

export async function generateMetadata(props: {
  params: Promise<{ id: string }>;
}): Promise<Metadata> {
  const params = await props.params;
  let user = await getUser(params.id);

  return {
    title: user && `User #${user.id}`,
  };
}
export default async function User(props: { params: Promise<{ id: string }> }) {
  const params = await props.params;
  const id = params.id;
  const [user] = await Promise.all([getUser(id)]);

  if (!user) {
    notFound();
  }

  return (
    <>
      <HeadingEdit name="Users" backLink="/dashboard/users" />
      <HeadingSubEdit
        name={`User #${user.id}`}
        badgeState={user.state}
        badgeText={user.state}
      />
      <EditUserForm user={user} />
    </>
  );
}
