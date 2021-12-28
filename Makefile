TAG := ssl-handshake:1.5.2

build-image:
	docker build -t $(TAG) .

push-image: build-image
	docker tag $(TAG) ptuladhar/$(TAG)
	docker push ptuladhar/$(TAG)