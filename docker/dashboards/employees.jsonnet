local grafana = import 'grafonnet/grafana.libsonnet';

local dashboard = grafana.dashboard;
local table = grafana.tablePanel;

dashboard.new(
  'Employees',
  schemaVersion=16,
  tags=['local']
)
.addPanel(
  table.new(
    'Employees',
    datasource='Employees',
  ).addTarget({
    delimiter: ',',
    header: true,
    ignoreUnknown: false,
    skipRows: 0,
    schema: [],
  },), gridPos={
    x: 0,
    y: 0,
    w: 12,
    h: 12,
  }
)
