import { create } from "zustand";

interface BearState {
  currentDate: string | null;
  availableDates: string[];
  setAvailableDates: (availableDates: string[]) => void;
  setCurrentDate: (date: string | null) => void;
}

export const useStore = create<BearState>((set) => ({
  currentDate: null,
  availableDates: [],
  setAvailableDates: (dates: string[]) =>
    set(() => ({ availableDates: dates })),
  setCurrentDate: (date: string | null) => set(() => ({ currentDate: date })),
}));
