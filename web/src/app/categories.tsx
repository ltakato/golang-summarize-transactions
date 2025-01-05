"use client";

import CategoryCard from "@/components/category-card/category-card";
import { Category } from "@/app/app-types";
import { useCategories } from "@/hooks/categories-hooks";
import { useStore } from "@/store/store";

export default function Categories() {
  const { setCurrentCategory } = useStore();
  const { data } = useCategories();

  const onClick = (category: Category) => {
    setCurrentCategory(category);
  };

  return (
    <div className="grid gap-4 grid-cols-4">
      {(data as Category[]).map((category: Category, index: number) => (
        <CategoryCard key={index} onClick={onClick} category={category} />
      ))}
    </div>
  );
}
