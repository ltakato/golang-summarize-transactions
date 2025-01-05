"use client";

import { useQuery } from "@tanstack/react-query";
import { Category, CategoryTransaction } from "@/app/app-types";
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

export function useCategoryTransactions() {
  const { userId } = useParams();
  const { currentDate, currentCategory } = useStore();

  return useQuery<CategoryTransaction[]>({
    queryKey: [
      "category-transactions",
      currentDate,
      currentCategory?.id as string,
    ],
    queryFn: async () => {
      const apiClient = new TransactionsSummaryApiClient({
        userId: userId as string,
      });
      return apiClient.getCategoryTransactions(
        currentDate as string,
        currentCategory?.id as string,
      );
    },
    enabled: !!userId && !!currentDate && !!currentCategory,
  });
}
