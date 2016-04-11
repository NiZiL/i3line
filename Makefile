install: default

%: _themes/%.go
	go build -o i3line _themes/$@.go
	mv i3line $(GOBIN)

