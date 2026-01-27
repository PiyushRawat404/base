//for loop
let arr =[1,2,3,4,5]

for(let i=0;i<=arr.length;i++){
    if(i%2==0){
        console.log("divided by 2")
    }
    else{
        console.log("not divisible by 2")
    }

}

//while
let a =2
let sum =0

while(a<5){
    sum =sum +a;
    a++;
}
console.log(sum)

//do-while

let b=[1,2,3]
let add=0
let j=0

do{
    add+=b[j];
    j++;
}while(j<b.length)
console.log(add)

//for... in

let student = { name: "Piyush", age: 21, city: "Delhi" };

for (let key in student) {
  console.log(key, ":", student[key]);
}

//for ...of
for (let ch of "Hello") {
  console.log(ch);
}



