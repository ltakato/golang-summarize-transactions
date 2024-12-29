"use client";

import { useQuery } from "@tanstack/react-query";
import CategoryCard from "@/components/category-card/category-card";
import { Category } from "@/app/app-types";

export default function Categories() {
  const { isPending, error, data } = useQuery<Category[]>({
    queryKey: ["repoData"],
    queryFn: () =>
      fetch("http://localhost:8080/api/categories").then((res) => res.json()),
  });

  if (isPending) return "loading...";

  if (error) return "error!";

  return (
    <div className="flex-1 space-y-4 p-8 pt-6">
      <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-4">
        {data.map((category: Category, index: number) => (
          <CategoryCard key={index} category={category}></CategoryCard>
        ))}
      </div>
    </div>
  );
}
