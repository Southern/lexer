/*
  
  Animal can be used as a base for different types of animals.

*/
function Animal(name) {
  this.name = name;
}

// Make our animal say hello!
Animal.prototype.sayHello = function() {
  return 'Hello from ' + this.name;
};

// Make our animal make some noise!
Animal.prototype.makeNoise = function() {
  return this.noise || '<chirp>';
};

/*

  Aw, look. It's a cute little dog.

*/
function Dog(name, breed) {
  this.name = name;
  this.breed = breed;
  this.noise = 'Woof!';
}

// Inherit Animal
Dog.prototype = new Animal();

// Our dog is smart. He can say hello AND his breed.
Dog.prototype.sayExtendedHello = function() {
  return this.sayHello() + ', ' + this.breed;
};

// Our dog can also bark. Not as impressive.
Dog.prototype.bark = function() {
  return this.noise;
};

// Expose our Animal and Dog to the outside world.
module.exports = {
  Animal: Animal,
  Dog: Dog,
};
