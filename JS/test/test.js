let credit = 1000 // размер кредитов
let cash //баланс игрока
let bet // объект ставки
let rate //размер ставки
let win // размер выиграша
let color
let ruletka

alert('Игра \nРУЛЕТКА')
console.log('Правила игры:')
console.log('Размер кредитов: ' + credit)
console.log('Если выигрывает от 0 до 36 - Ваша ставка умножается на 35')
console.log('Если выигрывает красное или черное - Ваша ставка умножается на 2')
alert('Правила игры справа 👉')
console.log('')

for (cash = credit; cash >= 1;) {
    bet = prompt('Введите ставку \nот 0 до 36, красное или черное. \nquit для завершения игры. \nВаш баланс: ' + cash + '$');
    if (bet == 'quit' || bet == 'йгше') {
        alert('Вы завершили игру.');
        console.log('Вы завершили игру');
        break;
    } else if (bet >= 0 && bet <= 36 || bet == 'красное' || bet == 'черное') {
        rate = prompt('Введите размер ставки. ' + ' \nВаш баланс: ' + cash + ' $' + '\nquit для завершения игры.');
        if (rate == 'quit') {
            alert('Вы завершили игру.');
            console.log('Вы завершили игру');
            break;
        } else {
            if (rate > cash) {
                alert('Ваша ставка ' + rate + '$' + ' выше баланса ' + cash + '$')
                console.log('Ваша ставка ' + rate + '$' + ' выше баланса ' + cash + '$')
            } else {
                ruletka = Math.round(Math.random() * 36);
                if ((ruletka % 2) == 0) {
                    color = 'красное';
                } else {
                    color = 'черное';
                }
            }
        }
        //алгоритм выиграша
        if (bet == ruletka) {
            win = rate * 36;
        } else if (bet == color) {
            win = rate * 2;
        }
        if (bet == ruletka || bet == color) {
            console.log('ВЫИГРАЛИ!')
            console.log('Ваша ставка: ' + rate + '$' + ' на ' + bet);
            console.log('На рулетке: ' + ruletka + ', ' + color);
            cash = cash + win
            console.log('Ваш выиграш: ' + win + '$');
            console.log('Ваш баланс: ' + cash + '$');
            alert('ВЫИГРАЛИ!' + '\nВаша ставка: ' + rate + '$' + ' на ' + bet + '\nНа рулетке: ' + ruletka + ', ' + color + '\nВаш баланс: ' + cash + '$')
        } else {
            cash = cash - rate;
            console.log('ПРОИГРАЛИ!')
            console.log('Ваша ставка: ' + rate + '$' + ' на ' + bet);
            console.log('На рулетке: ' + ruletka + ', ' + color);
            console.log('Вы проиграли ставку: ' + rate + '$');
            console.log('Ваш баланс: ' + cash + '$');
            alert('ПРОИГРАЛИ!' + '\nВаша ставка: ' + rate + '$' + ' на ' + bet + '\nНа рулетке: ' + ruletka + ', ' + color + '\nВаш баланс: ' + cash + '$')
        }
    } else {
        alert('Неверный ввод: ' + bet + '\nВведите ставку \nот 0 до 36, красное или черное. \nquit для завершения игры.')
        console.log('Неверный ввод: ' + bet);
        console.log('Введите ставку от 0 до 36, красное или черное. quit для завершения игры.')
    }
}
console.log();
alert('Конец игры \nВаш баланс: ' + cash + '$')
console.log('Конец игры')
console.log('Ваш баланс: ' + cash + '$')
