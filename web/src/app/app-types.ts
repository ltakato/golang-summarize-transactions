export type Summary = {
  availableDates: string[];
};

export type Category = {
  id: string;
  name: string;
  totalAmount: number;
};

export type CategoryTransaction = {
  title: string;
  date: Date;
  amount: number;
};
