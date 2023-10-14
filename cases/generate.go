package cases

//go:generate protoc -I=. -I=.. --go_out=.  --go_opt=paths=source_relative bytes.proto  kitchen_sink.proto  messages.proto  oneofs.proto   repeated.proto   wkt_duration.proto   wkt_wrappers.proto bool.proto  enums.proto                 maps.proto          numbers.proto     strings.proto   wkt_any.proto  wkt_timestamp.proto other_package/embed.proto
