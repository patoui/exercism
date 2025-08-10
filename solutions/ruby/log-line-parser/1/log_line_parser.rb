class LogLineParser
  TYPE_INFO = 'INFO'.freeze
  TYPE_WARNING = 'WARNING'.freeze
  TYPE_ERROR = 'ERROR'.freeze

  def initialize(line)
    @line = line
    @line_items = line.match(/\[(#{TYPE_ERROR}|#{TYPE_WARNING}|#{TYPE_INFO})\]:\s(.*)/)
    @msg = @line_items[2].strip
    @lvl = @line_items[1].downcase
  end

  def message
    @msg
  end

  def log_level
    @lvl
  end

  def reformat
    "#{message} (#{log_level})"
  end
end
