package Present

type RealPresent struct {
}

func (r *RealPresent) Present(data string) string {
	mySl := make([]rune, 0, len(data))
	strK := []rune{'h', 't', 't', 'p', ':', '/', '/'} // контрольный слайс http://
	mySl = []rune(data)                               // переводим стринг в слайс рун
	mySl = append(mySl, ' ')                          // Добавили пробел
	for ind := 0; ind < len(mySl)-7; ind++ {          // цикл перебора реслайса
		if string(mySl[ind:ind+7]) == string(strK) { //контроль
			for { // цикл замены на *
				if mySl[ind+7] != ' ' { // контроль на пробел
					mySl[ind+7] = '*'
					ind++
				} else {
					break
				}
			}

		}
	}

	return string(mySl)
}
