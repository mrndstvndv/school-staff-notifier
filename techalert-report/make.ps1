param (
	[string]$subcommand
)

switch ($subcommand) {
    "proto" {
		protoc --plugin=protoc-gen-ts_proto=.\node_modules\.bin\protoc-gen-ts_proto.exe --ts_proto_out=.\src -I ..\ ..\issues.proto
    }
    default {
        Write-Output "Invalid input. Please use 'proto'."
    }
}
