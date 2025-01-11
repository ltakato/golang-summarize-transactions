import React from "react";
import { Bell, BellDot } from "lucide-react";

const NotificationBell: React.FC<{ count: number }> = ({ count }) => {
  if (count === 0) return <Bell />;

  return (
    <div className="relative inline-block">
      <BellDot className="w-8 h-8 text-gray-600" />
      {count > 0 && (
        <span
          className="absolute top-0 right-0 inline-flex items-center justify-center
          w-4 h-4 text-xs font-bold text-white bg-red-500 rounded-full"
        >
          {count}
        </span>
      )}
    </div>
  );
};

export default NotificationBell;
