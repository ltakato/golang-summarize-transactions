"use client";

import CategoryCard from "@/components/category-card/category-card";
import { Category } from "@/app/app-types";
import { useCategories } from "@/hooks/categories-hooks";

export default function Categories() {
  const { data } = useCategories();

  return (
    <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-6">
      {(data as Category[]).map((category: Category, index: number) => (
        <CategoryCard key={index} category={category} />
      ))}
    </div>
  );
}
