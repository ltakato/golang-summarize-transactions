"use client";

import { useQuery } from "@tanstack/react-query";
import { Summary } from "@/app/app-types";
import { useParams } from "next/navigation";

export function useSummary() {
  const { userId } = useParams();

  return useQuery<Summary>({
    queryKey: ["summary"],
    queryFn: async () => {
      const headers = new Headers({ "x-user-id": userId as string });
      const res = await fetch(
        `${process.env.NEXT_PUBLIC_API_URL}/api/summary`,
        { headers },
      );
      return res.json();
    },
    enabled: !!userId,
  });
}
