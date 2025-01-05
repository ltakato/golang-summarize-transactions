import { create } from "zustand";
import { devtools, subscribeWithSelector } from "zustand/middleware";
import { Category } from "@/app/app-types";

interface AppState {
  currentDate: string | null;
  setCurrentDate: (date: string | null) => void;
  currentCategory: Category | null;
  setCurrentCategory: (category: Category | null) => void;
}

export const useStore = create<AppState>()(
  devtools(
    subscribeWithSelector((set) => ({
      currentDate: null,
      setCurrentDate: (date: string | null) =>
        set(() => ({ currentDate: date })),
      currentCategory: null,
      setCurrentCategory: (category: Category | null) =>
        set(() => ({ currentCategory: category })),
    })),
  ),
);

useStore.subscribe(
  (state) => state.currentDate,
  () => useStore.getState().setCurrentCategory(null),
);
