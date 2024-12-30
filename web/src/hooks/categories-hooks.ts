"use client";

import { useQuery } from "@tanstack/react-query";
import { Category } from "@/app/app-types";
import { useStore } from "@/store/store";
import { useParams } from "next/navigation";
import { TransactionsSummaryApiClient } from "@/lib/transactions-summary-api-client";

export function useCategories() {
  const { userId } = useParams();
  const { currentDate } = useStore();

  return useQuery<Category[]>({
    queryKey: ["categories", currentDate],
    queryFn: async () => {
      const apiClient = new TransactionsSummaryApiClient({
        userId: userId as string,
      });
      return apiClient.getCategories(currentDate as string);
    },
    enabled: !!userId && !!currentDate,
  });
}
