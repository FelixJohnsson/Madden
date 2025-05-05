import * as t from "../types";

const sortSalesByMonth = (sales: t.Sale[]): number[] => {
  const salesByMonth: number[] = sales.reduce((acc, sale) => {
    const month = new Date(sale.date).getMonth();
    acc[month] += sale.amount;
    return acc;
  }, Array(12).fill(0));

  return salesByMonth;
};

export default sortSalesByMonth;
