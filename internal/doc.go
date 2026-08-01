package internal

func CPFCheckDigitsValid(cpf string) bool {
	if len(cpf) != 11 {
		return false
	}

	sum := 0
	for i := 0; i < 9; i++ {
		sum += int(cpf[i]-'0') * (10 - i)
	}

	d1 := (sum * 10) % 11
	if d1 == 10 {
		d1 = 0
	}

	if d1 != int(cpf[9]-'0') {
		return false
	}

	sum = 0
	for i := 0; i < 10; i++ {
		sum += int(cpf[i]-'0') * (11 - i)
	}

	d2 := (sum * 10) % 11
	if d2 == 10 {
		d2 = 0
	}

	return d2 == int(cpf[10]-'0')
}

func CNPJCheckDigitsValid(d string) bool {
	w1 := []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	w2 := []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}

	sum := 0
	for i := 0; i < 12; i++ {
		sum += int(d[i]-'0') * w1[i]
	}
	r := sum % 11
	dv1 := 0
	if r >= 2 {
		dv1 = 11 - r
	}
	if int(d[12]-'0') != dv1 {
		return false
	}

	sum = 0
	for i := 0; i < 13; i++ {
		sum += int(d[i]-'0') * w2[i]
	}
	r = sum % 11
	dv2 := 0
	if r >= 2 {
		dv2 = 11 - r
	}
	return int(d[13]-'0') == dv2
}
