import { BarChart } from "@mui/x-charts/BarChart";

const Chart = ({ data }: { data: number[] }) => {
  return (
    <BarChart
      series={[{ data: data }]}
      width={1000}
      xAxis={[
        {
          data: [
            "Jan",
            "Feb",
            "Mar",
            "Apr",
            "May",
            "Jun",
            "Jul",
            "Aug",
            "Sep",
            "Oct",
            "Nov",
            "Dec",
          ],
        },
      ]}
    />
  );
};

export default Chart;
