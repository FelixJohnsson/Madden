import { Link } from "react-router-dom";

const Navigation = () => {
  return (
    <div className="w-full h-16 flex items-center justify-left pl-10 border-b-2 border-gray-200">
      <div className="flex items-center justify-center gap-10 text-xl">
        <Link to="/sales-page">Sales Page</Link>
        <Link to="/purchase-order-page">Purchase Order Page</Link>
      </div>
    </div>
  );
};

export default Navigation;
