ACCOUNT := ptuladhar
REPO := ssl-handshake
VERSION := 1.6.1

build-image:
	docker build -t $(REPO):$(VERSION) .

push-image: build-image
	docker tag $(REPO):$(VERSION) $(ACCOUNT)/$(REPO):$(VERSION)
	docker push $(ACCOUNT)/$(REPO):$(VERSION)
	docker tag $(REPO):$(VERSION) $(ACCOUNT)/$(REPO):latest
	docker push $(ACCOUNT)/$(REPO):latest

