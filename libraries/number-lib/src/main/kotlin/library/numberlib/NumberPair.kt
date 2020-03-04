package library.numberlib

data class NumberPair(
    val first: Int,
    val second: Int
) {
    fun sum(): Int = first + second
}
