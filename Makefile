run:
	go build -o tmp-coco-application-gateway main.go;
	./tmp-coco-application-gateway run

build:
	go build -o coco-application-gateway main.go;