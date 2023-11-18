gen: proto/proto/*.proto
	mkdir -p gen
	for file in $^ ; do \
		protoc --go_out=gen --go_opt=paths=source_relative \
					--go-grpc_out=gen --go-grpc_opt=paths=source_relative \
					--proto_path proto/proto $${file} ; \
	done
