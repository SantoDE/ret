RET_IMAGE := $(if $(REPONAME),$(REPONAME),"santode/ret")

default: binary

binary:
	./script/make.sh binary

crossbinary:
	./script/make.sh crossbinary

image:
	docker build -t $(RET_IMAGE) .