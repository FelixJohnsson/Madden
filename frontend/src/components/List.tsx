import { Button, Stack } from "@mui/material";
import { DataGrid, GridColDef } from "@mui/x-data-grid";
import * as t from "../types";
import DeleteIcon from "@mui/icons-material/Delete";

const SalesList = ({ data }: { data: t.Sale[] }) => {
  const columns: GridColDef[] = [
    { field: "id", headerName: "ID", width: 70 },
    { field: "itemId", headerName: "Item ID", width: 100 },
    { field: "itemName", headerName: "Item Name", width: 150 },
    { field: "amount", headerName: "Amount", width: 100, type: "number" },
    { field: "currency", headerName: "Currency", width: 100 },
    { field: "date", headerName: "Date", width: 150 },
  ];

  return (
    <>
      <Stack direction="row" spacing={1} sx={{ mb: 1 }}></Stack>
      <div style={{ display: "flex", flexDirection: "column" }}>
        <DataGrid
          rows={data}
          columns={columns}
          loading={false}
          pageSizeOptions={[5, 10, 25]}
          initialState={{
            pagination: { paginationModel: { pageSize: 10 } },
          }}
        />
      </div>
    </>
  );
};

const PurchaseOrderList = ({
  data,
  handleDelete,
}: {
  data: t.PurchaseOrder[];
  handleDelete: (id: number) => void;
}) => {
  const columns: GridColDef[] = [
    { field: "id", headerName: "ID", width: 70 },
    { field: "itemId", headerName: "Item ID", width: 100 },
    { field: "itemName", headerName: "Item Name", width: 150 },
    { field: "amount", headerName: "Amount", width: 100, type: "number" },
    { field: "currency", headerName: "Currency", width: 100 },
    { field: "createdAt", headerName: "Date", width: 150 },
    { field: "status", headerName: "Status", width: 100 },
    {
      field: "actions",
      headerName: "Actions",
      width: 150,
      sortable: false,
      filterable: false,
      renderCell: (params) => (
        <Stack direction="row" spacing={1}>
          <Button
            size="small"
            color="error"
            startIcon={<DeleteIcon />}
            onClick={() => handleDelete(Number(params.id))}
          >
            Delete
          </Button>
        </Stack>
      ),
    },
  ];

  return (
    <>
      <Stack direction="row" spacing={1} sx={{ mb: 1 }}></Stack>
      <div style={{ display: "flex", flexDirection: "column" }}>
        <DataGrid
          rows={data}
          columns={columns}
          loading={false}
          pageSizeOptions={[5, 10, 25]}
          initialState={{
            pagination: { paginationModel: { pageSize: 10 } },
          }}
        />
      </div>
    </>
  );
};

export { SalesList, PurchaseOrderList };
