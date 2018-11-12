
paths = {
  "/index.html",
  "/css/lanyon.css",
  "/img/2013-03-29-finding-your-way-home-with-clojure.png",
}

counter = 0

request = function()
  path = paths[counter]
  counter = counter + 1
  if counter > #paths then
    counter = 0
  end
  return wrk.format(nil, path)
end

