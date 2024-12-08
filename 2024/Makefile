.PHONY: test

test:
	go test ./test/day$(d)
newDay:
	cp -r test/day02 test/day$(d)
	cp -r solution/day02 solution/day$(d)
	touch input/$(d).txt
