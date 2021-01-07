local grafana = import 'grafonnet/grafana.libsonnet';

local dashboard = grafana.dashboard;
local graph = grafana.graphPanel;

dashboard.new(
  'Natural Gas Prices',
  schemaVersion=16,
  tags=['http']
)
.addPanel(
  graph.new(
    'Price',
    datasource='Natural Gas Prices (DataHub)',
  ).addTarget({
    delimiter: ',',
    header: true,
    ignoreUnknown: false,
    skipRows: 0,
    schema: [{
      name: 'Date',
      type: 'time',
    }, {
      name: 'Price',
      type: 'number',
    }],
  },), gridPos={
    x: 0,
    y: 0,
    w: 24,
    h: 12,
  }
)
