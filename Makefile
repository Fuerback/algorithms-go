img := go-quicksort

image:
	docker build -t $(img) .

run:
	docker run $(img) $(NUMS)

example:
	docker run $(img) 10 30 41 53 78 12 19 22