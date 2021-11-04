## Вычислятор ОПН :)

Программа, которая умеет вводимое инфиксное выражение конвертировать в постфиксное (обратная польская нотация)
и вычислять его.

## Установка

Для запуска приложения, необходимо иметь установленный GoLang последней версии, скачать можно по
ссылке https://golang.org/dl/, далее скачиваем подходящий к вашей ОС.

## Запуск

В папке с программой в окне терминала напечатать

```shell
go mod tidy
go test ./...
go run .
```

Эти команды установят зависимости приложения, прогонят тесты и запустят приложение

## Решения

### REPL

Приложение реализовано как `REPL`, это значит что вы можете бесконечно вводить выражения в консольном интерфейсе и
программа будет их вычислять, после вычисления выводить результат на окно терминала и предлагать ввести новое выражение.

### Синтаксис

На вход могут подаваться строки в инфиксной форме. Например:

```text
(2 + 6) * 6
```

Разрешенные символы

| Символ   | Приоритет   | Унарный     | Обработка|
|:--------:|:-----------:|:-----------:|----------|
| +        | 2           | Да          | Складывает два операнда |
| -        | 2           | Да          | Вычитает два операнда |
| *        | 3           | Нет         | Умножает два операнда |
| /        | 3           | Нет         | Делит два операнда |
| (        | -           | -           | Открывает группу подвыражения |
| )        | -           | -           | Закрывает группу подвыражения |
| .        | -           | -           | Разделитель дробной и целой части числа |

#### Валидация

Все выражения поступающие в программу валидируются следующими валидаторами.

#### Пустое выражение

Пустые выражения и выражения состоящие только из пробелов не допускаются, пример:

```text
>>>
empty expressions not acceptable
```

#### Недопустимые символы

Все символы выражения проверяются на допустимость. Допустимы следующие символы

- Любые числа
- Пробелы
- Все символы из таблицы "Разрешенные символы"

Если интерпретатору на вход поступает выражение с недопустимым символом, в ответ возвращается ошибка. Пример:

```text
>>> 328 ; 12
328 ; 12
    ^
error: illegal character ';'
```

#### Неверное количество скобок

Если на вход поступает выражение в котором не совпадает количество открывающих и закрывающих скобок, такое выражение
отбрасывается и выводится ошибка. Пример

```text
>>> 2 * (3 / 7))
2 * (3 / 7))
^
error: incorrect number of open braces and closed braces [1:2]
```

В квадратных скобках указано количество открывающих и закрывающих скобок, в данном примере у нас есть 1 открывающая и 2
закрывающих скобки.

#### Некорректное начало выражения

Если выражение начинается с нескольких операторов, унарных или обычных, такое выражение отбрасывается и выводится
ошибка, пример:

```text
>>> --289 * 2
--289 * 2
 ^
error: after unary operator can be only open brace or number
```

При этом такое выражение допустимо:

```text
>>> -287 + 125
> -287 125 +
> -162
```

Другими словами. В начале выражения может быть только

- Число
- Открывающая скобка '('
- Унарный +
- Унарный -

Все остальные символы считаются запрещенными. Все пробелы в начале строки отбрасываются и на результат не влияют.

#### Последовательные операторы

В выражении не может быть подряд несколько операторов, пример:

```text
>>> 2 +* 3
2 +* 3
   ^
error: after '+' can follow only '-', '+', '(' or number
```

Такие выражения считаются некорректными и отбрасываются. Но есть случай когда такое выражение допустимо, например:

```text
>>> 2 +- 3
> 2 -3 +
> -1
```

В данном случае минус после плюса относится к числу 3 и делает его отрицательным. Но например когда у нас есть такое
выражение:

```text
>>> 2 - + - 3
2 - + - 3
    ^
error: three or more operators sequence
```

Мы получаем ошибку о большом количестве операторов между двумя операндами.

#### Некорректный последний символ

Когда выражение оканчивается оператором, оно отбрасывается и выводится ошибка, например:

```text
>>> 98 + 32 -
98 + 32 -
        ^
error: expression can't end with operator
```

Это защищает от использования постфиксных выражений. Либо некорректных инфиксных.

#### Перед открывающей скобкой должен быть оператор

За исключением случаев начала строки. Пример выражения:

```text
>>> 2 (3 * 2)
2 (3 * 2)
  ^
error: before open brace always must be operator except start of expression
```

Перед открывающей скобкой должен быть какой-нибудь оператор. Но такое выражение считается корректным

```text
>>> -(29+36) * 2
> 29 36 + 2 * -
> -130
```

#### Пустые скобки

Пустые скобки не допускаются, пример выражения которое выведет ошибку:

```text
>>> 9 * ()
9 * ()
    ^
error: empty braces not allowed
```

### Деление на ноль

Если в выражении в явном виде передан 0 в делителе или выражение подставляемое в делитель вычисляется в 0, в этом случае
интерпретатор выдает ошибку. Примеры:

```text
>>> 9 / 0
division by zero

>>> 135 / (10 - 8 - 2)
division by zero
```

### Дробные числа

Интерпретатор позволяет делать операции над дробными числами. Примеры:

```text
>>> 1.5 + 0.5
> 1.5 0.5 +
> 2

>>> 9.5 + 15
> 9.5 15 +
> 24.5
```

### Чувствительность к пробелам

Интерпретатор не чувствителен к пробелам, вы можете писать так:

```text
3 + 2
```

Так

```text
3+2
```

И даже так

```text
3     +2  - 8
```

В связи с этим есть эффект, можно разделять числа пробелами

```text
>>> 100 000 + 50 000
> 100000 50000 +
> 150000
```

Первое число преобразуется в `100000`, а второе в `50000`, я не стал перекрывать эту возможность намеренно.

### Не десятичные системы счисления

Интерпретатор может принять числа начинающиеся с `0`, но воспринимать их будет как десятичные, а не восьмеричные, все
ведущие нули отбросятся при вычислениях, но попадут в строковое представление ОПН. Пример:

```text
>>> 09 + 20
> 09 20 +
> 29
```

Я не стал решать эту проблему так как эффекта на вычисление она не оказывает, но накинул в голове несколько решений
данной проблемы.

- Можно на этапе валидации выражения проверять числа на наличие ведущих нулей
- В строковое представление ОПН складывать обработанные числа, тогда нулей не будет
- При токенизации выражения откидывать нули из числа вручную

### Примеры выражений и результатов

```text
>>> 9 * 2
> 9 2 *
> 18

>>> (1 + 2) * 4 + 3
> 1 2 + 4 * 3 +
> 15

>>> -(8 / 2) * 4 
> 8 2 / 4 * -
> -16

>>> 97 * 13 / (8 / 2) * (74 / 4)
> 97 13 * 8 2 / / 74 4 / *
> 5832.125
```

### Возможные улучшения

Наверное в следующей реализации я бы сделал валидацию и пост обработку токенов после токенизации, это дало бы
возможность писать более эффективные алгоритмы проверки токенов и прочее. При токенизации нужно было бы сохранить еще
позицию токенов в исходной строке, чтобы в дальнейшем при выводе ошибок ее можно было подсвечивать.

🐛 Конечно возможны баги 
