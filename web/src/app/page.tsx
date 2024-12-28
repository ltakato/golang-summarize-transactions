"use client";

import { QueryClientProvider, QueryClient } from "@tanstack/react-query";
import Categories from "@/app/categories";

const queryClient = new QueryClient();

export default function Home() {
  return (
    <QueryClientProvider client={queryClient}>
      <Categories />
    </QueryClientProvider>
  );
}
