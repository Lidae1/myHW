package Censor

/*func FuncCensor(str string) string {

	mySl := make([]rune, 0, len(str))
	strK := []rune{104, 116, 116, 112, 58, 47, 47} // контрольный слайс http://
	mySl = []rune(str)                             // переводим стрингу в слайс рун
	mySl = append(mySl, 32)                        // Добавили пробел
	for ind := 0; ind < len(mySl)-7; ind++ {       // цикл перебора реслайса
		if string(mySl[ind:ind+7]) == string(strK) { //контроль
			for { // цикл замены на *
				if mySl[ind+7] != rune(32) { // контроль на пробел
					mySl[ind+7] = rune(42)
					ind++
				} else {
					break
				}
			}

		}
	}

	return string(mySl)
}
*/
