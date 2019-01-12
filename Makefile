clean:
	cd cmd/web/ && rm web

build-serve:
	cd cmd/web && go build && ./web