GENERATED_DIR := protogen

proto: clean
	@mkdir ${GENERATED_DIR}
	@protoc --go_out=${GENERATED_DIR} --go_opt=paths=source_relative --go-grpc_out=${GENERATED_DIR} --go-grpc_opt=paths=source_relative controller.proto

clean:
	@rm -rf ${GENERATED_DIR}
