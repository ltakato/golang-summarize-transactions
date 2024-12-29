"use client";

import { useEffect } from "react";
import CategoryCard from "@/components/category-card/category-card";
import { Category } from "@/app/app-types";
import { useStore } from "@/store/store";
import { useCategories } from "@/hooks/categories-hooks";

export default function Categories() {
  const { currentDate } = useStore();
  const { isPending, error, data, refetch } = useCategories(currentDate);

  useEffect(() => {
    refetch();
  }, [currentDate, refetch]);

  if (isPending) return "loading...";

  if (error) return "error!";

  return (
    <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-6">
      {data.map((category: Category, index: number) => (
        <CategoryCard key={index} category={category} />
      ))}
    </div>
  );
}
