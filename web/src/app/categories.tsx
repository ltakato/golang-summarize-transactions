"use client";

import { useQuery } from "@tanstack/react-query";
import CategoryCard from "@/components/category-card/category-card";
import { Category } from "@/app/app-types";
import { useStore } from "@/store/store";
import { useEffect } from "react";

export default function Categories() {
  const { currentDate } = useStore();
  const { isPending, error, data, refetch } = useQuery<Category[]>({
    queryKey: ["categories"],
    queryFn: async () => {
      const res = await fetch(
        `http://localhost:8080/api/categories?date=${currentDate}`,
      );
      return res.json();
    },
  });

  useEffect(() => {
    refetch();
  }, [currentDate, refetch]);

  if (isPending) return "loading...";

  if (error) return "error!";

  return (
    <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-4">
      {data.map((category: Category, index: number) => (
        <CategoryCard key={index} category={category}></CategoryCard>
      ))}
    </div>
  );
}
