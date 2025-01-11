"use client";

import { useQuery } from "@tanstack/react-query";
import { useParams } from "next/navigation";
import { NotificationsApiClient } from "@/lib/notifications-api-client";

export function useNotifications() {
  const { userId } = useParams();

  return useQuery({
    queryKey: ["notifications"],
    refetchInterval: 15000,
    queryFn: async () => {
      const apiClient = new NotificationsApiClient({
        userId: userId as string,
      });
      return apiClient.getNotifications();
    },
    enabled: !!userId,
  });
}
