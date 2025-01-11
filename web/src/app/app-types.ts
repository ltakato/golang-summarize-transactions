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

export type Notification = {
  id: string;
  text: string;
  date: Date;
  read: boolean;
};
