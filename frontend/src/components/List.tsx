import { Stack } from "@mui/material";
import { DataGrid } from "@mui/x-data-grid";

const data = {
  rows: [
    { id: 1, col1: "Hello", col2: "World" },
    { id: 2, col1: "DataGrid", col2: "Demo" },
  ],
  columns: [
    { field: "col1", headerName: "Column 1", width: 150 },
    { field: "col2", headerName: "Column 2", width: 150 },
  ],
};
const List = () => {
  return (
    <>
      <Stack direction="row" spacing={1} sx={{ mb: 1 }}></Stack>
      <div style={{ display: "flex", flexDirection: "column" }}>
        <DataGrid
          {...data}
          rows={data.rows}
          loading={false}
          columns={data.columns}
        />
      </div>
    </>
  );
};

export default List;
