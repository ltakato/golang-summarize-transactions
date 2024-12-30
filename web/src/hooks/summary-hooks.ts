"use client";

import { useQuery } from "@tanstack/react-query";
import { Summary } from "@/app/app-types";
import { useParams } from "next/navigation";
import { TransactionsSummaryApiClient } from "@/lib/transactions-summary-api-client";

export function useSummary() {
  const { userId } = useParams();

  return useQuery<Summary>({
    queryKey: ["summary"],
    queryFn: async () => {
      const apiClient = new TransactionsSummaryApiClient({
        userId: userId as string,
      });
      return apiClient.getSummary();
    },
    enabled: !!userId,
  });
}
