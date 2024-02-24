run:
	./gradlew run --args=$(d)
test:
	./gradlew test --info

gen:
	cp src/main/kotlin/days/Day1.kt src/main/kotlin/days/Day$(d).kt
	cp src/test/kotlin/days/Day1Test.kt src/test/kotlin/days/Day$(d)Test.kt
	touch src/main/resources/input_day_$(d).txt
	touch src/test/resources/input_day_$(d).txt
