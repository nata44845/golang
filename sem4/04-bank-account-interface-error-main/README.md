# 04-bank-account-interface-error

Нам необходимо продолжить развивать наш банковский аккаунт таким образом, что при вызове функции WithDraw отдавалась 
кастомная ошибка withdrawError, интерфейс которой описан. Ошибка должна содержать текущий баланс, сколько пытались 
снять и метод Error() error должен возвращать текст ошибки. В main() мы должен обработать ошибку через errors.As(). 
Вывести текущий баланс и сколько пытались снять.

Deposit должен научится возвращать ошибку ErrNegativeValue и в функции main() нужно обработать эту ошибку через 
функции errors.Is()