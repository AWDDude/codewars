function solution(n) {
  const numerals = [
    { k: 3000, v: "MMM" },
    { k: 2000, v: "MM" },
    { k: 1000, v: "M" },
    { k: 900, v: "CM" },
    { k: 800, v: "DCCC" },
    { k: 700, v: "DCC" },
    { k: 600, v: "DC" },
    { k: 500, v: "D" },
    { k: 400, v: "CD" },
    { k: 300, v: "CCC" },
    { k: 200, v: "CC" },
    { k: 100, v: "C" },
    { k: 90, v: "XC" },
    { k: 80, v: "LXXX" },
    { k: 70, v: "LXX" },
    { k: 60, v: "LX" },
    { k: 50, v: "L" },
    { k: 40, v: "XL" },
    { k: 30, v: "XXX" },
    { k: 20, v: "XX" },
    { k: 10, v: "X" },
    { k: 9, v: "IX" },
    { k: 8, v: "VIII" },
    { k: 7, v: "VII" },
    { k: 6, v: "VI" },
    { k: 5, v: "V" },
    { k: 4, v: "IV" },
    { k: 3, v: "III" },
    { k: 2, v: "II" },
    { k: 1, v: "I" },
  ];

  let outStr = "";
  for (const numeral of numerals) {
    if (n < numeral.k) {
      continue;
    }
    n -= numeral.k;
    outStr += numeral.v;
    if (n == 0) {
      break;
    }
  }

  return outStr;
}

console.log(solution(1000));
console.log(solution(887));
