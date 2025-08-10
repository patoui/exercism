=begin
Write your code for the 'Luhn' exercise in this file. Make the tests in
`luhn_test.rb` pass.

To get started with TDD, see the `README.md` file in your
`ruby/luhn` directory.
=end

module Luhn
  def self.valid?(card_number)
    card_number = card_number.delete(' ')
    if card_number.length < 2
      return false
    end

    i = card_number.length - 1
    is_even = false
    card_sum = 0

    until i < 0
      if !card_number[i].match?(/[[:digit:]]/)
        return false
      end

      current_number = card_number[i].to_i

      if is_even
        new_val = current_number * 2
        if new_val > 9
          new_val -= 9
        end
        card_sum += new_val
      else
        card_sum += current_number
      end
      is_even = !is_even
      i -= 1
    end

    card_sum == 80 || card_sum % 10 == 0
  end
end
