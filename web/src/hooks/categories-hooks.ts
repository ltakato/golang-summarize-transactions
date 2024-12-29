"use client";

import { useQuery } from "@tanstack/react-query";
import { Category } from "@/app/app-types";

export function useCategories(currentDate?: string | null) {
  return useQuery<Category[]>({
    queryKey: ["categories"],
    queryFn: async () => {
      if (!currentDate) {
        return [];
      }

      const res = await fetch(
        `${process.env.NEXT_PUBLIC_API_URL}/api/categories?date=${currentDate}`,
      );
      return res.json();
    },
  });
}
