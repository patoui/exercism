=begin
Write your code for the 'Isogram' exercise in this file. Make the tests in
`isogram_test.rb` pass.

To get started with TDD, see the `README.md` file in your
`ruby/isogram` directory.
=end

module Isogram
  def self.isogram?(input)
    letter_count = []
    input.downcase.each_char { |l|
      if l.match(/^[A-Za-z]/)
        if !letter_count.include?(l)
          letter_count << l
        else
          return false
        end
      end
    }
    return true
  end
end
