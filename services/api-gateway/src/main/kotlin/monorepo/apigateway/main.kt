package monorepo.apigateway

import library.numberlib.NumberPair
import library.stringlib.StringPair

fun main() {
    println("${StringPair("Api Gateway", "started").concat()} The answer for the universe is ${NumberPair(19, 23).sum()}")
}
