param (
	[string]$subcommand
)

function build
{
	$env:CGO_ENABLED=1
	go build -o .\out\ .
}

function run
{
	.\out\bongserver.exe
}

switch ($subcommand)
{
	"proto"
	{
		protoc --go_out=. --proto_path=..\ ..\issues.proto
		break
	}
	"build"
	{
		build
		break
	}
	"run"
	{
		build
		run
		break
	}
	default
	{
		Write-Host "Subcommand is not found"
	}
}
