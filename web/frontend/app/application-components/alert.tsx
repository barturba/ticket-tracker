import { deleteCookie, getCookie, getCookies, hasCookie } from "cookies-next";
import { useEffect } from "react";
import { toast, Bounce, ToastContainer } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";

export default function Alert() {
  const cookies = getCookies();
  useEffect(() => {
    if (hasCookie("success")) {
      toast.success(getCookie("success"), {
        position: "bottom-center",
        autoClose: false,
        closeOnClick: true,
        draggable: false,
        progress: undefined,
        theme: "dark",
        transition: Bounce,
      });
      deleteCookie("success");
    } else if (hasCookie("error")) {
      toast.error(getCookie("error"), {
        position: "bottom-center",
        autoClose: false,
        closeOnClick: true,
        draggable: false,
        progress: undefined,
        theme: "dark",
        transition: Bounce,
      });
      deleteCookie("error");
    } else if (hasCookie("info")) {
      toast.info(getCookie("info"), {
        position: "bottom-center",
        autoClose: false,
        closeOnClick: true,
        draggable: false,
        progress: undefined,
        theme: "dark",
        transition: Bounce,
      });
      deleteCookie("info");
    }
  }, [cookies]);
  return <ToastContainer />;
}
