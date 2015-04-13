import math._

object IntPDF {

	def dnorm(x: Double) = {

		exp(-x * x / 2)/sqrt(2 * 3.14159)

	}

	def main(args: Array[String]) {

		var N = args(0).toInt
		var Res = 0

		for(i <- 0 to N) {

			var candx = random * 10 - 5
			var candy = random * .5
			if(candy < dnorm(candx)) {
				Res = Res + 1
			}

		}

		var out = 5.0 * Res / N
		println(out.toString)

	}
}
