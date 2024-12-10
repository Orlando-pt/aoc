.PHONY: test

test:
	go test ./test/day$(d)
newDay:
	cp -r test/day0 test/day$(d)
	cp -r solution/day0 solution/day$(d)
	touch input/$(d).txt
