class AssemblyLine
  def initialize(speed)
    @speed = speed
  end

  def production_rate_per_hour
    success_rate = 1
    if @speed == 10
      success_rate = 0.77
    elsif @speed == 9
      success_rate = 0.8
    elsif @speed > 4
      success_rate = 0.9
    end
    return @speed * 221 * success_rate
  end

  def working_items_per_minute
    (production_rate_per_hour / 60).to_i
  end
end
