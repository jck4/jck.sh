/*
Add 1 letter of the string to an element at a time, on a 10 ms delay. 
*/

function typeEffect(terminalLine, characterPosistion, stringLength, htmlElement) {
  setTimeout(function () {
    if (characterPosistion < stringLength) {
      htmlElement.innerHTML += terminalLine.charAt(characterPosistion++);
      typeEffect(terminalLine, characterPosistion, stringLength, htmlElement); //Recurse
    }
  }, 10);
}