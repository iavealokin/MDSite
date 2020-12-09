package model

//Bill struct
type Bill struct {
	Id          int     // Уникальный id счёта
	Name        string  // Имя счёта (неуникально)
	BalanceRUB  int     // Текущий баланс счёта в рублях
	BalanceUSD  int     // Текущий баланс счёта в долларах
	BillType    string  // Тип счёта (кредитка, кредит, ипотека, дебет и т.д.)
	Rate        float32 // Ставка по счёту (для кредитных продуктов)
	Overpayment int     // Переплата по счёту (для кредитных продуктов)
	Bank        string  // Название кредитной органзации
}

//GetUserBills ..
func GetUserBills() (bills []Bill, err error) {
	bills = []Bill{
		{1, "Тинькофф", 1050, 15, "дебет", 0, 0, "Tinkoff"},
		{2, "Убрир", 1400, 20, "дебет", 0, 0, "Ubrir"},
	}
	return
}
