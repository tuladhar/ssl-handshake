TAG := ssl-handshake:1.6.0

build-image:
	docker build -t $(TAG) .

push-image: build-image
	docker tag $(TAG) ptuladhar/$(TAG)
	docker push ptuladhar/$(TAG)