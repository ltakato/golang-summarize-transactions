"use client";

import { useQuery } from "@tanstack/react-query";
import CategoryCard from "@/app/category-card";

type Category = {
  name: string;
  totalAmount: number;
};

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
          <CategoryCard key={index} title={category.name}>
            <div className="text-2xl font-bold">$ {category.totalAmount}</div>
            <p className="text-xs text-muted-foreground">
              +20.1% from last month
            </p>
          </CategoryCard>
        ))}
      </div>
    </div>
  );
}
