"use client";

import { QueryClientProvider, QueryClient } from "@tanstack/react-query";
import MainPage from "@/app/main-page";

const queryClient = new QueryClient();

export default function Home() {
  return (
    <QueryClientProvider client={queryClient}>
      <MainPage />
    </QueryClientProvider>
  );
}
