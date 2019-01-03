function sum(...nums) {
	let t = 0

	for (n of nums) {
		t += n
	}

	return t
}

const total = sum(1, 2, 3, 4, 5)
console.log(total)
