plugins {
    application
    kotlin("jvm") version "1.9.23"
}

application {
    mainClass.set("util.Runner")
}

version = "1.0-SNAPSHOT"

repositories {
    mavenCentral()
}

dependencies {
    implementation(kotlin("stdlib-jdk8"))
    implementation("org.reflections", "reflections", "0.10.2")
    implementation("org.slf4j:slf4j-nop:2.0.9")

    testImplementation("org.junit.jupiter", "junit-jupiter-api", "5.10.1")
    testImplementation("org.junit.jupiter", "junit-jupiter-engine", "5.10.1")
    testImplementation("org.assertj:assertj-core:3.25.3")
    testImplementation("com.github.stefanbirkner:system-lambda:1.2.1")
}

tasks.test {
    useJUnitPlatform()
}

kotlin {
    jvmToolchain(17)
}
