// Task 1
// WRITE YOUR CODE HERE - Create the logDairy function and use a for...of loop to log each item in the dairy array

// Task 2
// WRITE YOUR CODE HERE - Create the birdCan function and use a for...of loop to log bird object's own properties

// Task 3
// WRITE YOUR CODE HERE - Create the animalCan function and use a for...in loop to log all bird properties, including inherited ones
var dairy = ['cheese', 'sour cream', 'milk', 'yogurt', 'ice cream', 'milkshake'];


function logDairy() {
    for (items of dairy) {
       console.log(items)
   }
}

logDairy();


const animal = {
    canJump : true
};

const bird = Object.create(animal);
bird.canFly = true;
bird.hasFeathers = true;

function birdCan() {
    for (propertey of Object.getOwnPropertyNames(bird)) {
        console.log(`${propertey}: ${bird[propertey]}`);
    }
}

birdCan()

function animalCan() {
    for (property in bird) {
          console.log(`${property}: ${bird[property]}`);
    }
}

animalCan()
