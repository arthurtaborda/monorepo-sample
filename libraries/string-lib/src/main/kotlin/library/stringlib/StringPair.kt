package library.stringlib

data class StringPair(
    val first: String,
    val second: String
) {
    fun concat(): String = "$first $second"
}
