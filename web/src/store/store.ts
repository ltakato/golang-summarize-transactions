import { create } from "zustand";

interface BearState {
  currentDate: string | null;
  setCurrentDate: (date: string | null) => void;
}

export const useStore = create<BearState>((set) => ({
  currentDate: null,
  setCurrentDate: (date: string | null) => set(() => ({ currentDate: date })),
}));
