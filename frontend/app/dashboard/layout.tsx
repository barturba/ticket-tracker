"use client";

import { useState } from "react";
import Image from "next/image";
import logo from "@/public/static/images/logo.png";

import {
  Dialog,
  DialogBackdrop,
  DialogPanel,
  TransitionChild,
} from "@headlessui/react";
import { Bars3Icon, XMarkIcon } from "@heroicons/react/24/outline";
import NavLinks from "@/app/ui/dashboard/nav-links";

export default function Layout({ children }: { children: React.ReactNode }) {
  return <>{children}</>;
}
