import Navigation from "../components/Navigation";

import Button from "../components/Button";
import Chart from "../components/Chart";
import { SalesList } from "../components/List";

import { getSales } from "../api/api";
import { useState, useEffect } from "react";
import * as t from "../types";
import sortSalesByMonth from "../utils/sortSalesByMonth";

const SalesPage = () => {
  const [sales, setSales] = useState<t.Sale[]>([]);
  const [salesByMonth, setSalesByMonth] = useState<number[]>([]);
  const [selectedItem, setSelectedItem] = useState<string>("");

  useEffect(() => {
    getSales().then((sales) => {
      setSalesByMonth(sortSalesByMonth(sales));
      setSales(sales);
    });
  }, []);

  useEffect(() => {
    if (selectedItem) {
      const filteredSales = sales.filter(
        (sale) => sale.itemName === selectedItem
      );

      setSalesByMonth(sortSalesByMonth(filteredSales));
    } else {
      setSalesByMonth(sortSalesByMonth(sales));
    }
  }, [selectedItem]);

  return (
    <>
      <Navigation />
      <div className="flex flex-col items-center justify-center">
        <h1 className="text-3xl font-bold">Sales Page</h1>
      </div>

      <Chart data={salesByMonth} />
      <div className="flex flex-wrap items-center justify-center gap-2 mx-auto w-2/3">
        {Array.from(new Set(sales.map((sale) => sale.itemName))).map(
          (itemName) => (
            <Button
              key={itemName}
              onClick={() => setSelectedItem(itemName)}
              selected={selectedItem === itemName}
            >
              {itemName}
            </Button>
          )
        )}
        <Button
          onClick={() => setSelectedItem("")}
          selected={selectedItem === ""}
        >
          All
        </Button>
      </div>
      <SalesList data={sales} />
    </>
  );
};

export default SalesPage;
