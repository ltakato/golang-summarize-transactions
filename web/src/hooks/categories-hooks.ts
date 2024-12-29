"use client";

import { useQuery } from "@tanstack/react-query";
import { Category } from "@/app/app-types";
import { useStore } from "@/store/store";
import { useParams } from "next/navigation";

export function useCategories() {
  const { userId } = useParams();
  const { currentDate } = useStore();

  return useQuery<Category[]>({
    queryKey: ["categories", currentDate],
    queryFn: async () => {
      const headers = new Headers({ "x-user-id": userId as string });
      const res = await fetch(
        `${process.env.NEXT_PUBLIC_API_URL}/api/categories?date=${currentDate}`,
        { headers },
      );
      return res.json();
    },
    enabled: !!userId && !!currentDate,
  });
}
