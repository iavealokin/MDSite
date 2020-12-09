package model

//Bill struct
type Bill struct {
	id          int     // Уникальный id счёта
	name        string  // Имя счёта (неуникально)
	balanceRUB  int     // Текущий баланс счёта в рублях
	balanceUSD  int     // Текущий баланс счёта в долларах
	billType    string  // Тип счёта (кредитка, кредит, ипотека, дебет и т.д.)
	rate        float32 // Ставка по счёту (для кредитных продуктов)
	Overpayment int     // Переплата по счёту (для кредитных продуктов)
}

//GetUserBills ..
func GetUserBills() (bills []Bill, err error) {
	bills = []Bill{
		{1, "Тинькофф", 1050, 15, "дебет", 0, 0},
		{1, "Убрир", 1400, 20, "дебет", 0, 0},
	}
	return
}
