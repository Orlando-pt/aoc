package days

class Day4 : Day(4) {

    private val eyeColors = listOf("amb", "blu", "brn", "gry", "grn", "hzl", "oth")
    override fun partOne(): Any {

        return parseInput().count {
            it.isValid()
        }
    }

    override fun partTwo(): Any {
        return parseInput().count {
            it.isValidPart2()
        }
    }

    private data class Passport(
        var byr: String? = null,
        var iyr: String? = null,
        var eyr: String? = null,
        var hgt: String? = null,
        var hcl: String? = null,
        var ecl: String? = null,
        var pid: String? = null,
        var cid: String? = null
    )

    private fun Passport.attrList() = listOf(
        byr, iyr, eyr, hgt, hcl, ecl, pid
    )

    private fun Passport.isValid() = this.attrList().all { it != null }

    private fun Passport.isValidPart2(): Boolean {
        if (!this.isValid()) return false

        val birth = (byr?.toInt() ?: 0) in 1920..2002
        val issue = (iyr?.toInt() ?: 0) in 2010..2020
        val expiration = (eyr?.toInt() ?: 0) in 2020..2030

        val height = if (hgt!!.endsWith("cm"))
            hgt!!.substring(0, hgt!!.length - 2).toInt() in 150..193
        else if (hgt!!.endsWith("in")) hgt!!.substring(0, hgt!!.length - 2).toInt() in 59..76
        else false

        val hair = hcl!![0] == '#' && hcl!!.substring(1).matches("(\\d|[a-f]){6}".toRegex())

        val eye = eyeColors.any {
            it == ecl!!
        }

        val passport = pid!!.length == 9

        return birth && issue && expiration && height && hair && eye && passport
    }

    private fun parseInput(): List<Passport> {
        val passports = mutableListOf<Passport>()
        var passportLines = 0

        for (i in inputList.indices) {
            if (inputList[i].isBlank()) {
                passports.add(getPassport(inputList.subList(i - passportLines, i)))
                passportLines = 0
            } else {
                passportLines++
            }
        }
        passports.add(getPassport(inputList.subList(inputList.size - passportLines, inputList.size)))

        return passports
    }

    private fun getPassport(info: List<String>): Passport {
        val passport = Passport()
        for (line in info) {
            val lineSpaceParts = line.split(" ")

            for (part in lineSpaceParts)
                passport.parseStringProperty(part.split(":"))
        }

        return passport
    }

    private fun Passport.parseStringProperty(property: List<String>) = when (property[0].trim()) {
        "byr" -> this.byr = property[1].trim()
        "iyr" -> this.iyr = property[1].trim()
        "eyr" -> this.eyr = property[1].trim()
        "hgt" -> this.hgt = property[1].trim()
        "hcl" -> this.hcl = property[1].trim()
        "ecl" -> this.ecl = property[1].trim()
        "pid" -> this.pid = property[1].trim()
        "cid" -> this.cid = property[1].trim()
        else -> error("ups")
    }

}
