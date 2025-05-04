/*

Display a list of purchase orders

Add new purchase orders

Delete existing purchase orders

*/

import Navigation from "../components/Navigation";

const PurchaseOrderPage = () => {
  return (
    <>
      <Navigation />
      <div className="flex flex-col items-center justify-center h-screen">
        <h1 className="text-3xl font-bold">Purchase Order Page</h1>
      </div>
    </>
  );
};

export default PurchaseOrderPage;
