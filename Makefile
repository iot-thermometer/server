#gen:
#	mkdir -p gen
#	for f in $(find proto/proto -name "*.proto"); do echo $(basename "$f"); done
		protoc --go_out=gen --go_opt=paths=source_relative \
				--go-grpc_out=gen --go-grpc_opt=paths=source_relative \
				--proto_path proto/proto *.proto;

gen: proto/proto/*.proto
	mkdir -p gen
	for file in $^ ; do \
		protoc --go_out=gen --go_opt=paths=source_relative \
					--go-grpc_out=gen --go-grpc_opt=paths=source_relative \
					--proto_path proto/proto $${file} ; \
	done