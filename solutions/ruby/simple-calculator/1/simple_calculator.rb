class SimpleCalculator
  ALLOWED_OPERATIONS = ['+', '/', '*'].freeze

  class UnsupportedOperation < RuntimeError; end

  def self.calculate(first_operand, second_operand, operation)
    if !ALLOWED_OPERATIONS.include? operation
      raise UnsupportedOperation.new("Unsupported operation.")
    end

    if !first_operand.is_a?(Numeric) || !second_operand.is_a?(Numeric)
      raise ArgumentError, "Invalid argument type, must be numeric"
    end

    if operation == '/' && second_operand == 0
      return "Division by zero is not allowed."
    end

    result = first_operand.public_send(operation, second_operand)

    "#{first_operand} #{operation} #{second_operand} = #{result}"
  end
end
