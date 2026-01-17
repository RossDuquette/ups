.PHONY: build

build:
	@./docker_run make -f target.mk build

%:
	@./docker_run make -f target.mk $@
