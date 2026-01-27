let a = 10;
let b = 5;
let op = "+";

switch (op) {
  case "+":
    console.log(a + b);
    break;

  case "-":
    console.log(a - b);
    break;

  case "*":
    console.log(a * b);
    break;

  case "/":
    console.log(a / b);
    break;

  default:
    console.log("error");
}


//continue
for (let i = 1; i <= 5; i++) {
  if (i === 3) {
    continue;  
  }
  console.log(i);
}
