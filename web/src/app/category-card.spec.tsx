import { render, screen } from "@testing-library/react";
import { Category } from "@/app/app-types";
import CategoryCard from "../app/category-card";

describe("CategoryCard", () => {
  const category: Category = {
    name: "Food",
    totalAmount: 150.75,
  };

  it("renders the category name and totalAmount", () => {
    render(<CategoryCard category={category} />);

    expect(screen.getByText(category.name)).toBeInTheDocument();
    expect(screen.getByText(`R$ ${category.totalAmount}`)).toBeInTheDocument();
  });
});
