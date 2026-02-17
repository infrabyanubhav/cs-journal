class Person{
    constructor(name, age, energy) {
        this.name = name;
        this.age = age;
        this.energy=energy
    }

    sleep() {
        return this.energy+10
    }
    doSomethingFun() {
        return this.energy-10
    }

}


class worker extends Person{
    constructor(name,age,energy,xp = 0, hourlyWage = 10) {
        super(name,age,energy)
        this.xp = xp
        this.hourlyWage=hourlyWage
        
    }

    goToWork() {
        return this.xp+10
    }
}

// Task 2: Code a Worker class
// WRITE YOUR CODE HERE - Define the Worker class that extends Person
// WRITE YOUR CODE HERE - Add a constructor with additional parameters
// WRITE YOUR CODE HERE - Add the goToWork() method

// Task 3: Code an intern object, run methods
function intern() {
    const intern1 = new worker(
        "Bob",
        21,
        110,0,10
    )
    intern1.goToWork()
    return intern1
    
}

// Task 4: Code a manager object, methods
function manager() {
     const manager1 = new worker(
        "Alice",
         30,
         120,
         100,
         30
    )
    manager1.doSomethingFun()
    return manager1

    
}
