Set-Location server
.\make.ps1 proto

Set-Location ../techalert-report/
.\make.ps1 proto

Set-Location ../techalert-dashboard/
bun run build:proto
