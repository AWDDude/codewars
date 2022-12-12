def create_phone_number(arr)
  return arr.map{ |x| x.to_s }.insert(0,"(").insert(4,")").insert(5," ").insert(9,"-").join("")
end