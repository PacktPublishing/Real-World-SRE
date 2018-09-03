require "sinatra"
require "statsd"
require "logger"
Statsd.logger = Logger.new(STDOUT)
$statsd = Statsd.new 'localhost', 9125

get "/" do
  $statsd.time "request.time" do
    $statsd.increment "request.count"

    "hello world"
  end
end

not_found do
  $statsd.time "request.time" do
    $statsd.increment "request.count"
    $statsd.increment "request.error"

    "This is not found."
  end
end

error do
  $statsd.time "request.time" do
    $statsd.increment "request.count"
    $statsd.increment "request.error"

    "An error occured!"
  end
end
