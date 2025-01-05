type UserInfo = {
  id: string;
  email: string;
};

export type Summary = {
  userInfo: UserInfo;
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
