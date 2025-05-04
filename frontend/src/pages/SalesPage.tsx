/*

GET sales data, likely grouped by month.

Filter by product(s) (query params or frontend filtering).

*/

import Navigation from "../components/Navigation";

import { BarChart } from "@mui/x-charts/BarChart";
import Button from "../components/Button";
import List from "../components/List";

import { getSales } from "../api/api";
import { useState, useEffect } from "react";
import * as t from "../types";

const ChartsOverviewDemo = ({ data }: { data: t.Sale[] }) => {
  return (
    <BarChart
      series={[
        { data: [35, 44, 24, 34] },
        { data: [51, 6, 49, 30] },
        { data: [15, 25, 30, 50] },
        { data: [60, 50, 15, 25] },
      ]}
      width={1000}
      xAxis={[{ data: ["Q1", "Q2", "Q3", "Q4"] }]}
    />
  );
};

const sortSalesByMonth = (sales: t.Sale[]): t.SaleGroupedByMonth[] => {};

const SalesPage = () => {
  const [sales, setSales] = useState<t.Sale[]>([]);
  const [salesByMonth, setSalesByMonth] = useState<t.SaleGroupedByMonth[]>([]);

  useEffect(() => {
    getSales().then((sales) => {
      setSales(sales);
      const sorted = sortSalesByMonth(sales);
      setSalesByMonth(sorted);
      console.log(salesByMonth);
    });
  }, []);

  return (
    <>
      <Navigation />
      <div className="flex flex-col items-center justify-center">
        <h1 className="text-3xl font-bold">Sales Page</h1>
      </div>
      <ChartsOverviewDemo data={sales} />
      <div className="flex flex-col items-center justify-center">
        <Button />
      </div>
      <List />
    </>
  );
};

export default SalesPage;
