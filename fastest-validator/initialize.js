// js/fastest-validator.js
const Validator = require('./index.min');

class GoJsonValidator {
   constructor(initOptions) {
      this.validator = new Validator(initOptions);
   }

   validate(schema, jsonData) {
      const checker = this.validator.compile(schema);
      return checker(jsonData);
   }
}

module.exports = GoJsonValidator;