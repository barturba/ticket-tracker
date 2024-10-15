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
        </Sidebar>
      }
    >
      {children}
    </SidebarLayout>
  );
}
