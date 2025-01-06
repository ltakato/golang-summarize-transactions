import { create, StateCreator } from "zustand";
import { devtools, subscribeWithSelector } from "zustand/middleware";
import { Category } from "@/app/app-types";

type AppStoreState = {
  currentDate: string | null;
  currentCategory: Category | null;
};
type AppStoreActions = {
  setCurrentDate: (date: string | null) => void;
  setCurrentCategory: (category: Category | null) => void;
};
type AppStore = AppStoreState & AppStoreActions;

const initialState: AppStoreState = {
  currentDate: null,
  currentCategory: null,
};
const storeInitializer: StateCreator<AppStore> = (set) => ({
  ...initialState,
  setCurrentDate: (date: string | null) => set(() => ({ currentDate: date })),
  setCurrentCategory: (category: Category | null) =>
    set(() => ({ currentCategory: category })),
});
const withSubscribeWithSelector = subscribeWithSelector(storeInitializer);
const withDevtools = devtools(withSubscribeWithSelector);

export const useStore = create<AppStore>()(withDevtools);

useStore.subscribe(
  (state) => state.currentDate,
  () => useStore.getState().setCurrentCategory(null),
);
