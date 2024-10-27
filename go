#!/usr/bin/env ruby

require "timeout"

CORRECT_CODE = "1234"

def countdown(seconds)
  seconds.downto(1) do |i|
    print "\rВремя до следующей попытки: #{i} сек."
    sleep(1)
  end
  puts "\n"
end

def ask_question
  puts "Вопрос1"
  answer = gets.chomp
  if answer == "100"
    puts "Правильно! Переходим к следующему этапу."
    true
  else
    puts "Ошибка. Попробуйте еще раз."
    false
  end
end

def code_entry
  attempts = 0
  wait_time = 5

  loop do
    puts "Введите код:"
    input = gets.chomp

    if input == CORRECT_CODE
      puts "Поздравляем! Вы ввели правильный код."
      return true
    else
      attempts += 1
      if attempts >= 5
        puts "Превышено количество попыток. Заморозка на 3 часа."
        countdown(3 * 60 * 60)  # 3 часа
        attempts = 0
        wait_time = 10
      else
        puts "Неверный код. Попробуйте снова через #{wait_time} секунд."
        countdown(wait_time)
        wait_time += 1
      end
    end
  end
end

# Основной игровой цикл
loop do
  break if ask_question
end

code_entry

puts "Финальное сообщение: Вы успешно прошли квест!"

