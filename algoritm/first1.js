/*program to insert at last

function insertElement() {
    let array = [1, 2, 3, 4, 5];


    let index = 5;

    
    let element = 73;
  
    array.splice(index, 0, element);
    console.log(array);
}

insertElement();*/
/*function checkDistinct(array) { 
    for (let i = 0; i < array.length; i++) { 
        if (array.indexOf(array[i]) !== i) { 
            return false; 
        } 
    } 
    return true; 
} 
  
// Input1 
const input1 = [7, 8, 1, 5, 9]; 
console.log(checkDistinct(input1)); 
  
//Input2 
const input2 = [7, 8, 1, 5, 5]; 
console.log(checkDistinct(input2));*/
function Elementfound(array, element) {
    
    for (let i = 0; i < array.length; i++) {
        if (array[i] === element) {
            return true; 
        }
    }
    
    return false;
}


const givenArray = [1, 2, 3, 4, 5];
const elementToFind = 3;

if (Elementfound(givenArray, elementToFind)) {
    console.log(true); 
} else {
    console.log(false); 
}
