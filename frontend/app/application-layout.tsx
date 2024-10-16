"use client";
import logo from "@/public/static/images/logo.png";
import {
  HomeIcon,
  Square2StackIcon,
  ArrowRightStartOnRectangleIcon,
  LightBulbIcon,
  ShieldCheckIcon,
  UserCircleIcon,
  CpuChipIcon,
  UserGroupIcon,
  BuildingOffice2Icon,
  ChevronUpIcon,
} from "@heroicons/react/24/outline";
import Image from "next/image";
import { usePathname } from "next/navigation";
import {
  Navbar,
  NavbarSpacer,
  NavbarSection,
  NavbarItem,
} from "@/app/components/navbar";
import { SidebarLayout } from "@/app/components/sidebar-layout";
import {
  Sidebar,
  SidebarBody,
  SidebarFooter,
  SidebarHeader,
  SidebarItem,
  SidebarLabel,
  SidebarSection,
} from "@/app/components/sidebar";
import { Avatar } from "@/app/components/avatar";
import {
  Dropdown,
  DropdownButton,
  DropdownDivider,
  DropdownItem,
  DropdownLabel,
  DropdownMenu,
} from "@/app/components/dropdown";
import { signIn, signOut, useSession } from "next-auth/react";
function AccountDropdownMenu({
  anchor,
}: {
  anchor: "top start" | "bottom end";
}) {
  return (
    <DropdownMenu className="min-w-64" anchor={anchor}>
      <DropdownItem href="#">
        <UserCircleIcon />
        <DropdownLabel>My account</DropdownLabel>
      </DropdownItem>
      <DropdownDivider />
      <DropdownItem href="#">
        <ShieldCheckIcon />
        <DropdownLabel>Privacy policy</DropdownLabel>
      </DropdownItem>
      <DropdownItem href="#">
        <LightBulbIcon />
        <DropdownLabel>Share feedback</DropdownLabel>
      </DropdownItem>
      <DropdownDivider />
      <DropdownItem href="#">
        <ArrowRightStartOnRectangleIcon />
        <DropdownLabel>Sign out</DropdownLabel>
      </DropdownItem>
    </DropdownMenu>
  );
}
export function ApplicationLayout({ children }: { children: React.ReactNode }) {
  const pathname = usePathname();
  const { data: session } = useSession();

  return (
    <SidebarLayout
      navbar={
        <Navbar>
          <NavbarSpacer />
          <NavbarSection>
            <Dropdown>
              <DropdownButton as={NavbarItem}>
                <Avatar src="/users/user.png" square />
              </DropdownButton>
              <AccountDropdownMenu anchor="bottom end" />
            </Dropdown>
          </NavbarSection>
        </Navbar>
      }
      sidebar={
        <Sidebar>
          <SidebarHeader>
            <Dropdown>
              <DropdownButton as={SidebarItem}>
                <Image
                  height={24}
                  width={24}
                  src={logo}
                  alt="Ticket Tracker Logo"
                />
                <SidebarLabel>Ticket Tracker</SidebarLabel>
              </DropdownButton>
            </Dropdown>
          </SidebarHeader>

          <SidebarBody>
            <SidebarSection>
              <SidebarItem
                href="/dashboard"
                current={pathname === "/dashboard"}
              >
                <HomeIcon />
                <SidebarLabel>Dashboard</SidebarLabel>
              </SidebarItem>

              <SidebarItem
                href="/dashboard/incidents"
                current={pathname.startsWith("/dashboard/incidents")}
              >
                <Square2StackIcon />
                <SidebarLabel>Incidents</SidebarLabel>
              </SidebarItem>

              <SidebarItem
                href="/dashboard/companies"
                current={pathname.startsWith("/dashboard/companies")}
              >
                <BuildingOffice2Icon />
                <SidebarLabel>Companies</SidebarLabel>
              </SidebarItem>

              <SidebarItem
                href="/dashboard/users"
                current={pathname.startsWith("/dashboard/users")}
              >
                <UserGroupIcon />
                <SidebarLabel>Users</SidebarLabel>
              </SidebarItem>

              <SidebarItem
                href="/dashboard/cis"
                current={pathname.startsWith("/dashboard/cis")}
              >
                <CpuChipIcon />
                <SidebarLabel>CIs</SidebarLabel>
              </SidebarItem>
            </SidebarSection>
          </SidebarBody>
          <SidebarFooter className="max-lg:hidden">
            <Dropdown>
              {session?.user ? (
                <>
                  <DropdownButton as={SidebarItem}>
                    <span className="flex min-w-0 items-center gap-3">
                      <Avatar
                        src={sessionStorage?.user?.image}
                        className="size-10"
                        square
                        alt=""
                      />
                      <span className="min-w-0">
                        <span className="block truncate text-sm/5 font-medium text-zinc-950 dark:text-white">
                          {sessionStorage?.user?.name}
                        </span>
                        <span className="block truncate text-xs/5 font-normal text-zinc-500 dark:text-zinc-400">
                          {sessionStorage?.user?.email}
                        </span>
                      </span>
                    </span>
                    <ChevronUpIcon />
                  </DropdownButton>
                  <DropdownButton as={SidebarItem}>
                    <UserCircleIcon />
                    <SidebarLabel onClick={() => signOut()}>
                      Sign out
                    </SidebarLabel>
                  </DropdownButton>
                </>
              ) : (
                <DropdownButton as={SidebarItem}>
                  <UserCircleIcon />
                  <SidebarLabel onClick={() => signIn()}>Sign in</SidebarLabel>
                </DropdownButton>
              )}
              <AccountDropdownMenu anchor="top start" />
            </Dropdown>
          </SidebarFooter>
        </Sidebar>
      }
    >
      {children}
    </SidebarLayout>
  );
}
