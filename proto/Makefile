GENERATED_DIR := protogen

proto: clean
	mkdir -p ${GENERATED_DIR}
	protoc controller.proto --go_out=${GENERATED_DIR}

clean:
	@rm -rf ${GENERATED_DIR}

